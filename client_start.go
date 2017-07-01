package main

import (
	"go_server/dms_server"
)

func main() {
	server.Initialize(server.ServerPort, server.DetermineMotionState)
}
