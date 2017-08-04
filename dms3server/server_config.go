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
	CheckInterval            int
	ServerPort               int
	PlayAudio                int
	AudioMotionDetectorStart string
	AudioMotionDetectorStop  string
	ScanForTime              bool
	AlwaysOnRange            []string
	IPBase                   string
	IPRange                  []int
	MacsToFind               []string
	Logging                  *dms3libs.StructLogging
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

	if config.AudioMotionDetectorStart == "" || !dms3libs.IsFile(config.AudioMotionDetectorStart) {
		config.AudioMotionDetectorStart = filepath.Join(dms3libs.GetPackageDir(), "/media/motion_start.wav")
	}

	if config.AudioMotionDetectorStop == "" || !dms3libs.IsFile(config.AudioMotionDetectorStop) {
		config.AudioMotionDetectorStop = filepath.Join(dms3libs.GetPackageDir(), "/media/motion_stop.wav")
	}
}
