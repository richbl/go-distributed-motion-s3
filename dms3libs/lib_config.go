// Package dms3libs configuration structures and variables
package dms3libs

import (
	"os"
	"path"
	"path/filepath"

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

	DMS3clientTOML    = "dms3client.toml"
	DMS3serverTOML    = "dms3server.toml"
	DMS3libsTOML      = "dms3libs.toml"
	DMS3buildTOML     = "dms3build.toml"
	DMS3dashboardTOML = "dms3dashboard.toml"
	DMS3mailTOML      = "dms3mail.toml"
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

// Setup performs common installer tasks and takes a list of binary and config files to copy
func DeviceInstaller(binFiles, configDirs []string) {

	// Load libs config file from dms3_release folder on remote device
	configPath := filepath.Join(DMS3Release, DMS3Config, DMS3Libs, DMS3clientTOML)
	LoadLibConfig(configPath)

	binaryInstallDir := filepath.Join(string(filepath.Separator), "usr", "local", "bin")
	configInstallDir := filepath.Join(string(filepath.Separator), "etc", "distributed-motion-s3")
	logDir := filepath.Join(string(filepath.Separator), "var", "log", "dms3")

	// Create log folder
	MkDir(logDir)

	// Copy binary files into binaryInstallDir
	for _, binFile := range binFiles {
		src := filepath.Join(DMS3Release, "cmd", binFile)
		dst := filepath.Join(binaryInstallDir, binFile)
		CopyFile(src, dst)
	}

	// Create config folder and copy configuration files
	MkDir(configInstallDir)
	for _, configDir := range configDirs {
		src := filepath.Join(DMS3Release, DMS3Config, configDir)
		CopyDir(src, configInstallDir)
	}

	// Remove the release directory
	RmDir(DMS3Release)
}

// InitComponent initializes the common configuration and logging for a dms3 component.
func InitComponent(configPath, componentName, componentTOML string, config interface{}, logging *StructLogging) {

	LogDebug(filepath.Base(GetFunctionName()))

	LoadLibConfig(filepath.Join(configPath, DMS3Libs, DMS3libsTOML))
	LoadComponentConfig(config, filepath.Join(configPath, componentName, componentTOML))

	SetLogFileLocation(logging)
	CreateLogger(logging)

	LogInfo(componentName + " " + GetProjectVersion() + " started")
}
