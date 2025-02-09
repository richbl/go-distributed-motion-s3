// Package dms3client configuration structures and variables
package dms3client

import (
	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

// clientConfig contains dms3Client configuration settings read from TOML file
var clientConfig *structSettings

// client-side configuration parameters
type structSettings struct {
	Server  *structServer
	Logging *dms3libs.StructLogging
}

// server connection details
type structServer struct {
	IP            string
	Port          int
	CheckInterval uint16
}
