package main

import (
	"github.com/roboime/roboime-atlas/src/connect"
)

func main() {
	go connect.ListenToVision()

	select {}
}
