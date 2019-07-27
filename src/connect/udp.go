package connect

import (
	"context"
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
	crt = "../localhost.crt"
	key = "../localhost.key"
)

func ListenToVision() {
	addr := net.UDPAddr{
		Port: 10020,
		IP:   net.ParseIP("224.5.23.2"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}

	var buf [1024]byte
	pkt := &ssl.SSL_WrapperPacket{}
	atlas := &roboIMEAtlas{}
	log.Println("Server started!")
	go StartRoboIMEAtlasServer(atlas)
	for {
		size, _, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			panic(err)
		}

		if err := proto.Unmarshal(buf[:size], pkt); err != nil {
			continue
		}

		atlas.currentFrame = pkt
	}
}

type roboIMEAtlas struct {
	currentFrame *ssl.SSL_WrapperPacket // TODO: make this a buffer
}

func (r *roboIMEAtlas) GetFrame(timestamp *ssl.Timestamp, stream ssl.RoboIMEAtlas_GetFrameServer) error {
	if r.currentFrame == nil {
		return nil
	}
	log.Println("geometry", r.currentFrame.GetGeometry())
	return stream.Send(r.currentFrame)
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
		Addr:    fmt.Sprintf(":9090"),
		Handler: h,
	}

	logger.Infof("Starting server on port 9091 with tls")

	if err := httpServer.ListenAndServeTLS(crt, key); err != nil {
		logger.Fatalf("failed starting http2 server: %v", err)
	}
	// if err := httpServer.ListenAndServe(); err != nil {
	// 	logger.Fatalf("failed starting http2 server: %v", err)
	// }
}

func ClientTest() {
	creds, err := credentials.NewClientTLSFromFile(crt, "")
	if err != nil {
		panic(err)
	}
	conn, err := grpc.Dial("https://127.0.0.1:9090", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := ssl.NewRoboIMEAtlasClient(conn)

	log.Println("Client test started")
	for {
		time.Sleep(17 * time.Millisecond)
		stream, err := client.GetFrame(context.Background(), &ssl.Timestamp{})
		if err != nil {
			continue
		}
		for {
			pkt, err := stream.Recv()
			if err != nil {
				break
			}

			log.Println(pkt)
		}
	}
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
