package main

import (
	"go-distributed-motion-s3/dms3libs"
	"go-distributed-motion-s3/dms3mail"
)

func main() {

	dms3libs.LoadLibConfig("dms3libs/lib_config.toml")
	dms3mail.LoadMailConfig("dms3mail/mail_config.toml")

	cfg := dms3mail.MailConfig.Logging
	dms3libs.CreateLogger(cfg.LogLevel, cfg.LogDevice, cfg.LogLocation, cfg.LogFilename)

	dms3mail.GenerateEventEmail()

}
