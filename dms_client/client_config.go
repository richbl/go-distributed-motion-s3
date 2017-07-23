package client

import (
	"go_server/dms_libs"
)

// LogLevel sets the log levels for application logging using the following table:
//
//		0 - OFF, no logging
//		1 - FATAL, report fatal events
//		2 - INFO, report informational events
//    4 - DEBUG, report debugging events
//
const LogLevel = 2

// LogDevice determines to what device logging should be set using the following table:
//
//		0 - STDOUT (terminal)
//		1 - log file
//
// Ignored if LogLevel == 0
//
const LogDevice = 0

// LogFilename is the logging filename
// Ignored if LogLevel == 0 or LogDevice == 0
//
const LogFilename = "dms_client.log"

// LogLocation is the location of logfile (full path)
// By default, this is in the local folder (e.g., /etc/distributed_motion_surveillance/dms_client)
// Ignored if LogLevel == 0 or LogDevice == 0
//
var LogLocation = dmslibs.GetPackageDir()

// CheckInterval is the interval (in seconds) for checking the dms server
const CheckInterval = 5

// ServerIP is the address on which the dms server is running
const ServerIP = "localhost"

// ServerPort is the port on which the dms server is running
const ServerPort = 1965
