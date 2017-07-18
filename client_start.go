package main

import (
	"go_server/dms_client"
	"go_server/dms_libs"
)

func main() {
	dmslibs.CreateLogger(client.LogLevel, client.LogDevice, client.LogLocation, client.LogFilename)
	client.StartClient(client.ServerIP, client.ServerPort)
}
