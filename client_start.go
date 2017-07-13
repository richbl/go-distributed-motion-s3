package main

import (
	"go_server/dms_client"
	"go_server/dms_libs"
)

func main() {
	dmslibs.CreateLogger(client.Logging, client.LogLocation, client.LogFilename)
	client.Initialize(client.ServerIP, client.ServerPort, client.ProcessMotionState)
}
