package main

import (
	"github.com/roboime/roboime-atlas/src/connect"
)

func main() {
	atlas := connect.NewRoboIMEAtlas()
	go atlas.ListenToVision()

	select {}
}
