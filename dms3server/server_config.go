package dms3server

import (
	"go-distributed-motion-s3/dms3libs"
	"path/filepath"
	"time"
)

var startTime time.Time

// serverConfig contains dms3Server configuration settings read from TOML file
var serverConfig *structSettings

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
	Port          int
	CheckInterval int
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

// setMediaLocation sets the location where audio files are located for motion detection
// application start/stop
func setMediaLocation(config *structSettings) {

	if config.Audio.PlayMotionStart == "" || !dms3libs.IsFile(config.Audio.PlayMotionStart) {
		config.Audio.PlayMotionStart = filepath.Join(dms3libs.GetPackageDir(), "/media/motion_start.wav")
	}

	if config.Audio.PlayMotionStop == "" || !dms3libs.IsFile(config.Audio.PlayMotionStop) {
		config.Audio.PlayMotionStop = filepath.Join(dms3libs.GetPackageDir(), "/media/motion_stop.wav")
	}
}
