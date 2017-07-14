package main

import (
	"fmt"
	"go_server/dms_libs"
)

func main() {
	fmt.Println(dmslibs.RunningMotion())
	// dmslibs.CreateLogger(server.Logging, server.LogLocation, server.LogFilename)
	// server.Initialize(server.ServerPort, server.DetermineMotionState)
}
