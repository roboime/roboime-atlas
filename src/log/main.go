package main

import (
	"io"
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

func serveLogPackages(reader *persistence.Reader, addr *net.UDPAddr, refbox *net.UDPAddr) {
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}

	refConn, err := net.DialUDP("udp", nil, refbox)
	if err != nil {
		panic(err)
	}

	log.Printf("starting vision at %v:%v", addr.IP, addr.Port)
	log.Printf("starting refbox at %v:%v", refbox.IP, refbox.Port)
	for {
		msg, err := reader.ReadMessage()
		if err != nil {
			if err == io.EOF {
				log.Println("End of log!")
			} else {
				panic(err)
			}
		}

		if msg.MessageType.Id == persistence.MessageSslVision2014 {
			conn.Write(msg.Message)
		}

		if msg.MessageType.Id == persistence.MessageSslRefbox2013 {
			refConn.Write(msg.Message)
		}

		time.Sleep(time.Millisecond * interval)
	}
}

type server struct {
	reader *persistence.Reader
	addr   *net.UDPAddr
	refbox *net.UDPAddr
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
				refbox: &net.UDPAddr{
					IP:   net.ParseIP("127.0.0.1"),
					Port: initialPort + 1,
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
		go serveLogPackages(server.reader, server.addr, server.refbox)
	}

	select {}
}
