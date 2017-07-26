package main

import (
	"go-distributed-motion-s3/dms3libs"
	"go-distributed-motion-s3/dms3server"
)

func main() {

	dms3libs.CreateLogger(dms3server.LogLevel, dms3server.LogDevice, dms3server.LogLocation, dms3server.LogFilename)
	dms3server.StartServer(dms3server.ServerPort)

}
