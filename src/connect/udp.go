package connect

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"

	"github.com/roboime/roboime-atlas/src/protos/ssl"
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
	return stream.Send(r.currentFrame)
}

func StartRoboIMEAtlasServer(atlas *roboIMEAtlas) {
	lis, err := net.Listen("tcp", "127.0.0.1:10002")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	ssl.RegisterRoboIMEAtlasServer(grpcServer, atlas)

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func ClientTest() {
	conn, err := grpc.Dial("127.0.0.1:10002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := ssl.NewRoboIMEAtlasClient(conn)

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
