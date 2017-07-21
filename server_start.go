package main

import (
	"go_server/dms_libs"
	"go_server/dms_server"
)

func main() {

	dmslibs.CreateLogger(server.LogLevel, server.LogDevice, server.LogLocation, server.LogFilename)
	server.StartServer(server.ServerPort)

}
