package dms3mail

import (
	"go-distributed-motion-s3/dms3libs"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// MailConfig contains dms3mail configuration settings read from TOML file
var MailConfig *structSettings

// motion mail configuration parameters
type structSettings struct {
	EmailFrom              string
	EmailTo                string
	EmailBody              string
	SMTPAddress            string
	SMTPPort               int
	SMTPDomain             string
	SMTPUsername           string
	SMTPPassword           string
	SMTPAuthentication     string
	SMTPEnableStartTLSAuto bool
	Logging                *dms3libs.StructLogging
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
