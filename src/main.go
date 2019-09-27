package main

import (
	"github.com/roboime/roboime-atlas/src/connect"
)

const (
	visionUpdateInterval = 3  // approx. 240 fps
	refboxUpdateInterval = 17 // approx. 60 fps
)

func main() {
	atlas, err := connect.NewRoboIMEAtlas()
	if err != nil {
		panic(err)
	}
	go atlas.StartRoboIMEAtlasServer()
	go atlas.ListenToVision(visionUpdateInterval)
	go atlas.ListenToRefbox(refboxUpdateInterval)

	select {}
}
