// Package dms3libs configuration structures and variables
package dms3libs

import (
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

// Configuration file names
const (
	DMS3Mail      = "dms3mail"
	DMS3Client    = "dms3client"
	DMS3Server    = "dms3server"
	DMS3Libs      = "dms3libs"
	DMS3Release   = "dms3_release"
	DMS3Dashboard = "dms3dashboard"
	DMS3Config    = "config"
	DMS3TOML      = "dms3client.toml"
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

		if _, err := toml.DecodeFile(configFile, &LibConfig); err != nil {
			LogFatal(err.Error())
		}

	} else {
		LogFatal(configFile + " file not found")
	}

}

// LoadComponentConfig loads a TOML configuration file of client/server configs into parameter values
func LoadComponentConfig(structConfig interface{}, configFile string) {

	// check if config file exists
	if _, err := os.Stat(configFile); err != nil {
		LogFatal(err.Error())
	}

	if _, err := toml.DecodeFile(configFile, structConfig); err != nil {
		LogFatal(err.Error())
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
