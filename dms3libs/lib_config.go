package dms3libs

import (
	"log"

	"github.com/BurntSushi/toml"
)

// LibConfig contains dms3Libs configuration settings read from TOML file
var LibConfig *structConfig

type structConfig struct {
	SysCommands mapSysCommands
}

// mapSysCommands provides a location mapping of required system commands
type mapSysCommands map[string]string

// StructLogging is used for dms3 logging
type StructLogging struct {
	LogLevel    int
	LogDevice   int
	LogFilename string
	LogLocation string
}

// LoadLibConfig loads a TOML configuration file and parses entries into parameter values
func LoadLibConfig(configFile string) {

	if !IsFile(configFile) {
		log.Fatalln(configFile + " file not found")
	}

	if _, err := toml.DecodeFile(configFile, &LibConfig); err != nil {
		log.Fatalln(err.Error())
	}

}
