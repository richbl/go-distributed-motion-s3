package main

import (
	"fmt"
	"go_server/dms_libs"
	"go_server/dms_server"
)

func main() {
	dmslibs.CreateLogger(server.Logging, server.LogLocation, server.LogFilename)

	fmt.Println(dmslibs.GetPID("chrome"))

	// server.Initialize(server.ServerPort, server.DetermineMotionState)
}
