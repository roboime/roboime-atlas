package main

import (
	"os"
	"time"

	"github.com/roboime/roboime-atlas/src/connect"
)

const (
	nOfClients      = 1
	requestInterval = 100
	durationSeconds = 20
)

func main() {
	after := time.After(durationSeconds * time.Second)

	tick := time.Tick((requestInterval / nOfClients) * time.Millisecond)
	for i := 0; i < nOfClients; i++ {
		client := connect.NewClient()
		go connect.RunClient(client, requestInterval)
		<-tick
	}

	<-after

	os.Exit(0)
}
