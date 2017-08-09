package main

import (
	"go-distributed-motion-s3/dms3libs"
	"go-distributed-motion-s3/dms3server"
)

func main() {

	dms3libs.LoadLibConfig("/etc/distributed-motion-s3/dms3libs/dms3libs.toml")
	dms3server.LoadServerConfig("/etc/distributed-motion-s3/dms3server/dms3server.toml")

	cfg := dms3server.ServerConfig.Logging
	dms3libs.CreateLogger(cfg.LogLevel, cfg.LogDevice, cfg.LogLocation, cfg.LogFilename)

	dms3server.StartServer(dms3server.ServerConfig.ServerPort)

}
