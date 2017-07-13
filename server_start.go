package main

import (
	"go_server/dms_libs"
	"go_server/dms_server"
)

func main() {
	dmslibs.CreateLogger(server.Logging, server.LogLocation, server.LogFilename)
	server.Initialize(server.ServerPort, server.DetermineMotionState)
}
