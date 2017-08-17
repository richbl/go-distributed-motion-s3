package dms3server

import (
	"go-distributed-motion-s3/dms3libs"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

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

// LoadServerConfig loads a TOML configuration file and parses entries into parameter values
func LoadServerConfig(configFile string) {

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Fatalln(configFile + " structConfig file not found")
	} else if err != nil {
		log.Fatalln(err.Error())
	}

	if _, err := toml.DecodeFile(configFile, &ServerConfig); err != nil {
		log.Fatalln(err.Error())
	}

	setLogLocation(ServerConfig)
	setMediaLocation(ServerConfig)

}

func setLogLocation(config *structSettings) {

	if config.Logging.LogLocation == "" || !dms3libs.IsFile(config.Logging.LogLocation) {
		config.Logging.LogLocation = dms3libs.GetPackageDir()
	}

}

func setMediaLocation(config *structSettings) {

	if config.Audio.PlayMotionStart == "" || !dms3libs.IsFile(config.Audio.PlayMotionStart) {
		config.Audio.PlayMotionStart = filepath.Join(dms3libs.GetPackageDir(), "/media/motion_start.wav")
	}

	if config.Audio.PlayMotionStop == "" || !dms3libs.IsFile(config.Audio.PlayMotionStop) {
		config.Audio.PlayMotionStop = filepath.Join(dms3libs.GetPackageDir(), "/media/motion_stop.wav")
	}
}
