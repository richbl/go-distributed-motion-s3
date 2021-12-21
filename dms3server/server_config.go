// Package dms3server configuration structures and variables
//
package dms3server

import (
	"log"
	"path"
	"path/filepath"
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
// application start/stop events
//
func setMediaLocation(configPath string, config *structSettings) {

	type mediaPath struct {
		configLocation *string
		mediaLocation  string
	}

	media := []mediaPath{
		{
			&config.Audio.PlayMotionStart,
			"/dms3server/media/motion_start.wav",
		},
		{
			&config.Audio.PlayMotionStop,
			"/dms3server/media/motion_stop.wav",
		},
	}

	fail := false

	for i := range media {

		relPath := filepath.Join(configPath, media[i].mediaLocation)
		devPath := filepath.Join(path.Dir(dms3libs.GetPackageDir()), media[i].mediaLocation)

		if !dms3libs.IsFile(*media[i].configLocation) {

			// if no location set, set to release folder, else set to development folder
			if *media[i].configLocation == "" {

				if dms3libs.IsFile(relPath) {
					*media[i].configLocation = relPath
				} else if dms3libs.IsFile(devPath) {
					*media[i].configLocation = devPath
				} else {
					fail = true
				}

			} else {
				fail = true
			}

			if fail {
				log.Fatalln("unable to set media location... check TOML configuration file")
			}

		}

	}

}
