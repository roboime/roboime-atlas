package main

import (
	"io/ioutil"
	"log"
	"net"
	"regexp"
	"time"

	"github.com/RoboCup-SSL/ssl-go-tools/pkg/persistence"
)

const (
	interval = 5
)

func serveLogPackages(reader *persistence.Reader, addr *net.UDPAddr) {
	conn, err := net.DialUDP("udp", nil, addr)

	if err != nil {
		panic(err)
	}

	log.Printf("starting broadcast at %v:%v", addr.IP, addr.Port)
	for {
		msg, err := reader.ReadMessage()
		if err != nil {
			panic(err)
		}
		if msg.MessageType.Id == persistence.MessageSslVision2014 {
			_, err := msg.ParseVisionWrapper()
			if err != nil {
				panic(err)
			}

			_, err = conn.Write(msg.Message)

			time.Sleep(time.Millisecond * interval)
		}
	}
}

type server struct {
	reader *persistence.Reader
	addr   *net.UDPAddr
}

func buildLocalServers() ([]*server, error) {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		return nil, err
	}

	servers := []*server{}

	initialPort := 10020

	reg := regexp.MustCompile(`(.)+\.log\.gz`)
	for _, file := range files {
		if reg.MatchString(file.Name()) {

			reader, err := persistence.NewReader(file.Name())
			if err != nil {
				return nil, err
			}

			s := &server{
				addr: &net.UDPAddr{
					IP:   net.ParseIP("127.0.0.1"),
					Port: initialPort,
				},
				reader: reader,
			}

			initialPort += 5

			servers = append(servers, s)
		}
	}

	return servers, nil
}

func main() {
	servers, err := buildLocalServers()
	if err != nil {
		panic(err)
	}

	for _, server := range servers {
		go serveLogPackages(server.reader, server.addr)
	}

	select {}
}
