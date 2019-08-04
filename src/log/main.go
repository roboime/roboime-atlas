package main

import (
	"fmt"
	"net"
	"time"

	"github.com/RoboCup-SSL/ssl-go-tools/pkg/persistence"
)

const (
	interval = 5
)

func main() {
	reader, err := persistence.NewReader("logfile.log.gz")
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	addr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 10020,
	}

	for {
		msg, err := reader.ReadMessage()
		if err != nil {
			panic(err)
		}
		if msg.MessageType.Id == persistence.MessageSslVision2014 {
			pkg, err := msg.ParseVisionWrapper()
			if err != nil {
				panic(err)
			}

			fmt.Println(pkg)
			conn, err := net.DialUDP("udp", nil, addr)
			if err != nil {
				panic(err)
			}

			_, err = conn.Write(msg.Message)
			if err != nil {
				panic(err)
			}

			time.Sleep(time.Millisecond * interval)
		}
	}
}
