package main

import (
	"go-distributed-motion-s3/dms3libs"
	"go-distributed-motion-s3/dms3server"
)

func main() {

	dms3libs.LoadLibConfig("dms3libs/lib_config.toml")
	dms3server.LoadServerConfig("dms3server/server_config.toml")

	cfg := dms3server.ServerConfig.Logging
	dms3libs.CreateLogger(cfg.LogLevel, cfg.LogDevice, cfg.LogLocation, cfg.LogFilename)

	dms3server.StartServer(dms3server.ServerConfig.ServerPort)

}
