package dms3libs

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// LibConfig contains dms3Libs configuration settings read from TOML file
var LibConfig *structConfig

type structConfig struct {
	SysCommands mapSysCommands
}

// mapSysCommands provides a location mapping of required system commands
type mapSysCommands map[string]string

// StructLogging comment
type StructLogging struct {
	LogLevel    int
	LogDevice   int
	LogFilename string
	LogLocation string
}

// LoadLibConfig loads a TOML configuration file and parses entries into parameter values
func LoadLibConfig(configFile string) {

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Fatalln(configFile + " structConfig file not found")
	} else if err != nil {
		log.Fatalln(err.Error())
	}

	if _, err := toml.DecodeFile(configFile, &LibConfig); err != nil {
		log.Fatalln(err.Error())
	}

}
