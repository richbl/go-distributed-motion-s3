package main

import (
	"go_server/dms3client"
	"go_server/dms3libs"
)

func main() {
	dms3libs.CreateLogger(dms3client.LogLevel, dms3client.LogDevice, dms3client.LogLocation, dms3client.LogFilename)
	dms3client.StartClient(dms3client.ServerIP, dms3client.ServerPort)
}
