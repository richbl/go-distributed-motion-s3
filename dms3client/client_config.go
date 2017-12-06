package dms3client

import (
	"go-distributed-motion-s3/dms3libs"
<<<<<<< Updated upstream
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// ClientConfig contains dms3Client configuration settings read from TOML file
var ClientConfig *structSettings
=======
	"time"
)

var startTime time.Time

// clientConfig contains dms3Client configuration settings read from TOML file
var clientConfig *structSettings
>>>>>>> Stashed changes

// client-side configuration parameters
type structSettings struct {
	CheckInterval int
	ServerIP      string
	ServerPort    int
	Logging       *dms3libs.StructLogging
}

// LoadClientConfig loads a TOML configuration file and parses entries into parameter values
func LoadClientConfig(configFile string) {

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Fatalln(configFile + " structConfig file not found")
	} else if err != nil {
		log.Fatalln(err.Error())
	}

	if _, err := toml.DecodeFile(configFile, &ClientConfig); err != nil {
		log.Fatalln(err.Error())
	}

	setLogLocation(ClientConfig)

}

func setLogLocation(config *structSettings) {

	if config.Logging.LogLocation == "" || !dms3libs.IsFile(config.Logging.LogLocation) {
		config.Logging.LogLocation = dms3libs.GetPackageDir()
	}

}
