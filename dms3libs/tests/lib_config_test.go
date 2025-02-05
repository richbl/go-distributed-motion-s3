package dms3libs_test

import (
	"path/filepath"
	"testing"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

func TestLoadLibConfig(t *testing.T) {

	libsLocation := filepath.Join("..", "..", dms3libs.DMS3Config, "dms3libs.toml")

	dms3libs.LoadLibConfig(libsLocation)
	t.Log("libs configuration file loaded from", libsLocation, "successfully")

}

func TestConfiguration(t *testing.T) {

	for k, v := range dms3libs.LibConfig.SysCommands {

		if dms3libs.IsFile(v) {
			t.Log(k, "confirmed at", v)
		} else {
			t.Error(k, "not found at", v)
		}

	}

}

func TestLoadComponentConfig(t *testing.T) {

	type structServer struct {
		Port          int
		CheckInterval int
		Logging       *dms3libs.StructLogging
	}

	type structSettings struct {
		Server *structServer
	}

	testSettings := new(structSettings)
	configPath := dms3libs.GetPackageDir()
	configLocation := filepath.Join("..", "..", dms3libs.DMS3Config, "dms3server.toml")

	dms3libs.LoadComponentConfig(&testSettings, filepath.Join(configPath, configLocation))
	t.Log("component configuration file loaded from", configLocation, "successfully")

}

func TestSetLogFileLocation(t *testing.T) {

	testSettings := new(dms3libs.StructLogging)
	testSettings.LogLocation = ""

	dms3libs.SetLogFileLocation(testSettings)
	t.Log("log location set to", testSettings.LogLocation, "successfully")

}
