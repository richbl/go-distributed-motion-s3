package main

import (
	"go_server/dms3libs"
	"go_server/dms3server"
)

func main() {

	dms3libs.CreateLogger(dms3server.LogLevel, dms3server.LogDevice, dms3server.LogLocation, dms3server.LogFilename)
	dms3server.StartServer(dms3server.ServerPort)

}
