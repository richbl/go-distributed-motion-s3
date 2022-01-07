// Package dms3server configuration structures and variables
//
package dms3server

import (
	"time"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

var startTime time.Time

// ServerConfig contains dms3Server configuration settings read from TOML file
var ServerConfig *structSettings

// server-side configuration parameters
type structSettings struct {
	Server    *structServer
	Audio     *structAudio
	AlwaysOn  *structAlwaysOn
	UserProxy *structUserProxy
	Logging   *dms3libs.StructLogging
}

// server details
type structServer struct {
	Port            int
	CheckInterval   int
	EnableDashboard bool
}

// audio parameters used when the motion detector application starts/stops
type structAudio struct {
	Enable          bool
	PlayMotionStart string
	PlayMotionStop  string
}

// Always On feature parameters (enable the motion detector application based on time of day)
type structAlwaysOn struct {
	Enable    bool
	TimeRange []string
}

// User proxy parameters (representing the user's existence on the LAN)
type structUserProxy struct {
	IPBase     string
	IPRange    []int
	MacsToFind []string
}
