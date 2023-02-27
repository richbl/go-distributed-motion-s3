// Package dms3libs configuration structures and variables
package dms3libs

import (
	"errors"
	"io/fs"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

// LibConfig contains dms3Libs configuration settings read from TOML file
var LibConfig *structConfig

type structConfig struct {
	SysCommands mapSysCommands
}

// mapSysCommands provides a location mapping of required system commands
type mapSysCommands map[string]string

// LoadLibConfig loads a TOML configuration file of system commands into parameter values
func LoadLibConfig(configFile string) {

	if IsFile(configFile) {
		if _, error := toml.DecodeFile(configFile, &LibConfig); error != nil {
			LogFatal(error.Error())
		}
	} else {
		LogFatal(configFile + " file not found")
	}

}

// LoadComponentConfig loads a TOML configuration file of client/server configs into parameter values
func LoadComponentConfig(structConfig interface{}, configFile string) {

	if _, error := os.Stat(configFile); error == nil {

		if _, error := toml.DecodeFile(configFile, structConfig); error != nil {
			LogFatal(error.Error())
		}

	} else {

		if errors.Is(error, fs.ErrNotExist) {
			LogFatal(configFile + " file not found")
		} else {
			LogFatal(error.Error())
		}

	}

}

// SetLogFileLocation sets the location of the log file based on TOML configuration
func SetLogFileLocation(config *StructLogging) {

	projectDir := path.Dir(GetPackageDir())

	if !IsFile(config.LogLocation) {
		if config.LogLocation == "" && IsFile(projectDir) { // if no config location set, set to development project folder
			config.LogLocation = projectDir
		} else {
			LogFatal("unable to set log location... check TOML configuration file")
		}
	}

}
