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
	"time"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"github.com/golang/protobuf/proto"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"

	"github.com/roboime/roboime-atlas/src/protos/ssl"
)

const (
	crt = "../local.crt"
	key = "../local.key"
)

// Flags
var (
	VisionAddress  *string
	VisionPort     *int
	ServiceAddress *string
	ServicePort    *int
)

func init() {
	VisionAddress = flag.String("visionaddr", "224.5.23.2", "")
	VisionPort = flag.Int("visionport", 10020, "")
	ServiceAddress = flag.String("serviceaddr", "127.0.0.1", "")
	ServicePort = flag.Int("serviceport", 9090, "")
}

// ListenToVision starts listening to vision and start the GRPC server
func ListenToVision() {
	flag.Parse()

	addr := net.UDPAddr{
		Port: *VisionPort,
		IP:   net.ParseIP(*VisionAddress),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}

	log.Printf("Connected vision on %v:%v", *VisionAddress, *VisionPort)

	var buf [2048]byte
	pkt := &ssl.SSL_WrapperPacket{}
	atlas := &roboIMEAtlas{
		cameraFrame: make([]*ssl.SSL_WrapperPacket, 8),
	}
	log.Println("Server started!")
	go StartRoboIMEAtlasServer(atlas)

	for {
		size, _, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			panic(err)
		}

		if err := proto.Unmarshal(buf[:size], pkt); err != nil {
			log.Println(err)
			continue
		}

		detection := pkt.GetDetection()
		if detection != nil {
			log.Println("read package ", *detection)
			atlas.cameraFrame[*detection.CameraId] = pkt
			continue
		}

		geometry := pkt.GetGeometry()
		if geometry != nil {
			atlas.lastGeometry = pkt
		}
	}
}

type roboIMEAtlas struct {
	cameraFrame  []*ssl.SSL_WrapperPacket
	lastGeometry *ssl.SSL_WrapperPacket
}

func (r *roboIMEAtlas) GetFrame(timestamp *ssl.Timestamp, stream ssl.RoboIMEAtlas_GetFrameServer) error {
	for _, frame := range r.cameraFrame {
		if frame != nil {
			err := stream.Send(frame)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func StartRoboIMEAtlasServer(atlas *roboIMEAtlas) {
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
				stream, err := client.GetFrame(context.Background(), &ssl.Timestamp{})
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
