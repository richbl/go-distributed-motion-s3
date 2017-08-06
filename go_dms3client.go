package main

import (
	"go-distributed-motion-s3/dms3client"
	"go-distributed-motion-s3/dms3libs"
)

func main() {

	dms3libs.LoadLibConfig("/etc/distributed-motion-s3/dms3libs.toml")
	dms3client.LoadClientConfig("/etc/distributed-motion-s3/dms3client.toml")

	cfg := dms3client.ClientConfig.Logging
	dms3libs.CreateLogger(cfg.LogLevel, cfg.LogDevice, cfg.LogLocation, cfg.LogFilename)

	dms3client.StartClient(dms3client.ClientConfig.ServerIP, dms3client.ClientConfig.ServerPort)

}
