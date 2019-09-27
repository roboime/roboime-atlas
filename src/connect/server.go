package connect

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"github.com/golang/protobuf/proto"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"

	"github.com/roboime/roboime-atlas/src/protos/ssl"
)

const (
	crt = "../localhost.crt"
	key = "../localhost.key"
)

// Flags
var (
	VisionAddress  *string
	VisionPort1    *int
	VisionPort2    *int
	RefboxAddress  *string
	RefboxPort1    *int
	RefboxPort2    *int
	ServiceAddress *string
	ServicePort    *int
)

func init() {
	VisionAddress = flag.String("visionaddr", "224.5.23.2", "")
	VisionPort1 = flag.Int("visionport1", 10020, "")
	VisionPort2 = flag.Int("visionport2", 10025, "")
	RefboxAddress = flag.String("refboxaddr", "224.5.23.2", "")
	RefboxPort1 = flag.Int("refboxport1", 10021, "")
	RefboxPort2 = flag.Int("refboxport2", 10026, "")
	ServiceAddress = flag.String("serviceaddr", "127.0.0.1", "")
	ServicePort = flag.Int("serviceport", 9090, "")
}

type matchConn struct {
	vision *net.UDPConn
	refbox *net.UDPConn
}

func buildAddresses() ([]net.UDPAddr, []net.UDPAddr) {
	vision1 := net.UDPAddr{
		Port: *VisionPort1,
		IP:   net.ParseIP(*VisionAddress),
	}
	vision2 := net.UDPAddr{
		Port: *VisionPort2,
		IP:   net.ParseIP(*VisionAddress),
	}

	refbox1 := net.UDPAddr{
		Port: *RefboxPort1,
		IP:   net.ParseIP(*RefboxAddress),
	}
	refbox2 := net.UDPAddr{
		Port: *RefboxPort2,
		IP:   net.ParseIP(*RefboxAddress),
	}

	return []net.UDPAddr{vision1, vision2}, []net.UDPAddr{refbox1, refbox2}
}

func buildConnections() ([]*matchConn, error) {
	visionAddrs, refboxAddrs := buildAddresses()

	conns := []*matchConn{}

	for i, visionAddr := range visionAddrs {
		vision, err := net.ListenUDP("udp", &visionAddr)
		if err != nil {
			return nil, err
		}

		refbox, err := net.ListenUDP("udp", &refboxAddrs[i])
		if err != nil {
			return nil, err
		}

		conns = append(conns, &matchConn{
			vision: vision,
			refbox: refbox,
		})
	}

	return conns, nil
}

// NewRoboIMEAtlas creates a new instance of RoboIMEAtlas
func NewRoboIMEAtlas() (*RoboIMEAtlas, error) {
	flag.Parse()
	conns, err := buildConnections()
	if err != nil {
		return nil, err
	}

	return &RoboIMEAtlas{
		lastGeometry: map[int]*ssl.SSL_WrapperPacket{},
		cameraFrame:  map[int][]*ssl.SSL_WrapperPacket{},
		refbox:       map[int]*ssl.SSL_Referee{},
		conns:        conns,
		mux:          &sync.Mutex{},
	}, nil
}

// ListenToVision starts listening to vision and start the GRPC server
func (atlas *RoboIMEAtlas) ListenToVision(updateInterval time.Duration) {

	initialFrames := map[int][]*ssl.SSL_WrapperPacket{}
	for i := range atlas.conns {
		initialFrames[i] = make([]*ssl.SSL_WrapperPacket, 8)
	}

	atlas.cameraFrame = initialFrames

	var buf [2048]byte
	log.Println("vision server started!")

	pkt := &ssl.SSL_WrapperPacket{}
	for {

		for i, conn := range atlas.conns {
			size, _, err := conn.vision.ReadFromUDP(buf[:])
			if err != nil {
				panic(err)
			}

			if err := proto.Unmarshal(buf[:size], pkt); err != nil {
				log.Println(err)
				continue
			}

			detection := pkt.GetDetection()
			if detection != nil {
				atlas.cameraFrame[i][*detection.CameraId] = &ssl.SSL_WrapperPacket{
					Detection: detection,
				}
				continue
			}

			geometry := pkt.GetGeometry()
			if geometry != nil {
				atlas.lastGeometry[i] = &ssl.SSL_WrapperPacket{
					Geometry: geometry,
				}
			}

			time.Sleep(updateInterval * time.Millisecond)
		}
	}
}

// ListenToRefbox starts listening to the refbox info
func (atlas *RoboIMEAtlas) ListenToRefbox(updateInterval time.Duration) {

	var buf [2048]byte
	log.Println("vision server started!")

	for {

		for i, conn := range atlas.conns {
			pkt := &ssl.SSL_Referee{}
			size, _, err := conn.refbox.ReadFromUDP(buf[:])
			if err != nil {
				panic(err)
			}

			if err := proto.Unmarshal(buf[:size], pkt); err != nil {
				log.Println(err)
				continue
			}
			atlas.mux.Lock()
			atlas.refbox[i] = pkt
			atlas.mux.Unlock()
		}

		time.Sleep(updateInterval * time.Millisecond)
	}
}

// RoboIMEAtlas defines the RoboIMEAtlas struct
type RoboIMEAtlas struct {
	cameraFrame  map[int][]*ssl.SSL_WrapperPacket
	lastGeometry map[int]*ssl.SSL_WrapperPacket

	refbox map[int]*ssl.SSL_Referee

	conns []*matchConn

	mux *sync.Mutex
}

// GetFrame retrieves a stream of detection frame given a match
func (atlas *RoboIMEAtlas) GetFrame(req *ssl.FrameRequest, stream ssl.RoboIMEAtlas_GetFrameServer) error {
	if req == nil {
		log.Fatalln("FATAL: empty request for GetFrame")
		return nil
	}

	for _, frame := range atlas.cameraFrame[int(req.MatchId)] {
		if frame != nil {
			err := stream.Send(frame)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// GetMatchInfo retrieves referee data from a specific match
func (atlas *RoboIMEAtlas) GetMatchInfo(ctx context.Context, req *ssl.MatchInfoRequest) (*ssl.SSL_Referee, error) {
	matchID := req.MatchId
	atlas.mux.Lock()
	ref, ok := atlas.refbox[int(matchID)]
	atlas.mux.Unlock()

	if !ok {
		return nil, fmt.Errorf("invalid match id")
	}

	return ref, nil
}

// GetGeometry returns the match's geometry
func (atlas *RoboIMEAtlas) GetGeometry(ctx context.Context, req *ssl.FrameRequest) (*ssl.SSL_GeometryData, error) {
	matchID := req.MatchId

	pkt, ok := atlas.lastGeometry[int(matchID)]
	if !ok {
		return nil, fmt.Errorf("invalid match id")
	}

	//	fmt.Println("GetGeometry request for ", req.MatchId)

	if pkt == nil {
		return nil, nil
	}

	return pkt.GetGeometry(), nil
}

// GetActiveMatches returns info about the active matches being gathered
func (atlas *RoboIMEAtlas) GetActiveMatches(ctx context.Context, req *ssl.ActiveMatchesRequest) (*ssl.MatchesPacket, error) {
	matches := []*ssl.MatchData{}

	atlas.mux.Lock()
	refboxMap := atlas.refbox
	atlas.mux.Unlock()
	for i, matchRef := range refboxMap {
		data := &ssl.MatchData{
			MatchId: int32(i),
		}

		if matchRef.Yellow != nil && matchRef.Blue != nil {
			data.MatchName = fmt.Sprintf("%v x %v", *matchRef.Yellow.Name, *matchRef.Blue.Name)
		}

		matches = append(matches, data)
	}

	return &ssl.MatchesPacket{
		Match: matches,
	}, nil
}

// StartRoboIMEAtlasServer starts the grpc RoboIMEAtlas server
func (atlas *RoboIMEAtlas) StartRoboIMEAtlasServer() {
	grpcServer := grpc.NewServer()
	ssl.RegisterRoboIMEAtlasServer(grpcServer, atlas)
	logger := grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)

	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}
	h := allowCORS(http.HandlerFunc(handler))
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", *ServicePort),
		Handler: h,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	logger.Infof("Starting server on port %d with tls", *ServicePort)

	if err := httpServer.ListenAndServeTLS(crt, key); err != nil {
		logger.Fatalf("failed starting http2 server: %v", err)
	}
}

// ClientTest simulates clients
func ClientTest(nOfClients, interval, duration int) {
	tick := time.Tick(time.Duration(interval+(interval/nOfClients)) * time.Millisecond)
	quit := make(chan bool)
	for i := 0; i < nOfClients; i++ {
		go func(i int) {
			creds, err := credentials.NewClientTLSFromFile(crt, "")
			if err != nil {
				panic(err)
			}

			addr := fmt.Sprintf("%s:%d", *ServiceAddress, *ServicePort)
			conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds), grpc.WithBlock(), grpc.WithTimeout(time.Minute))
			if err != nil {
				panic(err)
			}
			defer conn.Close()

			client := ssl.NewRoboIMEAtlasClient(conn)

			log.Println("Client test started")
			for {
				time.Sleep(time.Duration(interval) * time.Millisecond)
				stream, err := client.GetFrame(context.Background(), &ssl.FrameRequest{
					MatchId: 0,
				})
				if err != nil {
					// log.Println(err)
					continue
				}
				for {
					_, err := stream.Recv()
					if err != nil {
						// log.Println(err)
						break
					}

					// log.Println("goroutine", i)
				}

				select {
				case <-quit:
					return
				default:
				}
			}
		}(i)

		<-tick
	}

	after := time.After(time.Duration(duration) * time.Second)
	<-after

	for i := 0; i < nOfClients; i++ {
		quit <- true
	}
	log.Println("test ended")

	os.Exit(0)
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization", "x-grpc-web"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}
