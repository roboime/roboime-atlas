package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/roboime/roboime-atlas/src/protos/ssl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	nOfClients      = 10
	requestInterval = 100
	durationSeconds = 60
	certFile        = "../../localhost.crt"
)

func main() {

	args := os.Args[1:]
	if len(args) > 0 {
		nOfClients, _ = strconv.Atoi(args[0])
	}

	creds, _ := credentials.NewClientTLSFromFile(certFile, "")
	conn, _ := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(creds))

	client := ssl.NewRoboIMEAtlasClient(conn)

	fmt.Printf("nOfClients: %v, requestInterval: %v, duration: %v\n", nOfClients, requestInterval, durationSeconds)
	ClientTest(client, nOfClients, requestInterval, durationSeconds)

	os.Exit(0)
}

// ClientTest simulates clients
func ClientTest(client ssl.RoboIMEAtlasClient, nOfClients, interval, duration int) {
	tick := time.Tick(time.Duration(interval+(interval/nOfClients)) * time.Millisecond)
	quit := make(chan bool)
	for i := 0; i < nOfClients; i++ {
		go func(i int) {
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
