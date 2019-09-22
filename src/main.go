package main

import (
	"github.com/roboime/roboime-atlas/src/connect"
)

func main() {
	atlas, err := connect.NewRoboIMEAtlas()
	if err != nil {
		panic(err)
	}
	go atlas.StartRoboIMEAtlasServer()
	go atlas.ListenToVision()
	go atlas.ListenToRefbox()

	select {}
}
