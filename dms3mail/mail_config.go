package dms3mail

import (
	"go-distributed-motion-s3/dms3libs"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// MailConfig contains dms3mail configuration settings read from TOML file
var MailConfig *structSettings

type structEmail struct {
	From string
	To   string
	Body string
}

type structSMTP struct {
	Address            string
	Port               int
	Domain             string
	Username           string
	Password           string
	Authentication     string
	EnableStartTLSAuto bool
}

// motion mail configuration parameters
type structSettings struct {
	Email   *structEmail
	SMTP    *structSMTP
	Logging *dms3libs.StructLogging
}

// LoadMailConfig loads a TOML configuration file and parses entries into parameter values
func LoadMailConfig(configFile string) {

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Fatalln(configFile + " structConfig file not found")
	} else if err != nil {
		log.Fatalln(err.Error())
	}

	if _, err := toml.DecodeFile(configFile, &MailConfig); err != nil {
		log.Fatalln(err.Error())
	}

	setLogLocation(MailConfig)

}

func setLogLocation(config *structSettings) {

	if config.Logging.LogLocation == "" || !dms3libs.IsFile(config.Logging.LogLocation) {
		config.Logging.LogLocation = dms3libs.GetPackageDir()
	}

}
