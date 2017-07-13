package client

import (
	"go_server/dms_libs"
)

// Logging enable (1) or disable (0) application logging
// NOTE: passing in 2 sets logging to STDOUT (useful when debugging or running as daemon)
const Logging = 2

// LogFilename is the logging filename
// Ignored if LOGGING == 0
const LogFilename = "dms_client.log"

// LogLocation is the location of logfile (full path)
// By default, this is in the local folder (e.g., /etc/distributed_motion_surveillance/dms_client)
// Ignored if LOGGING == 0
var LogLocation = dmslibs.GetPackageDir()

// CheckInterval is the interval (in seconds) for checking the dms server
const CheckInterval = 5

// ServerIP is the address on which the dms server is running
const ServerIP = "localhost"

// ServerPort is the port on which the dms server is running
const ServerPort = 1965
