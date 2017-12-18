package dms3libs_test

import (
	"go-distributed-motion-s3/dms3libs"
	"path/filepath"
	"testing"
)

type structSettings struct {
	Server *structServer
}

// server details
type structServer struct {
	Port          int
	CheckInterval int
	Logging       *dms3libs.StructLogging
}

func init() {
	dms3libs.LoadLibConfig("../../config/dms3libs.toml")
}

func TestLoadComponentConfig(t *testing.T) {

	testSettings := new(structSettings)
	configPath := dms3libs.GetPackageDir()

	dms3libs.LoadComponentConfig(&testSettings, filepath.Join(configPath, "../../config/dms3server.toml"))
	t.Log("component configuration loaded succesfully")

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

func TestSetLogFileLocation(t *testing.T) {

	testSettings := new(dms3libs.StructLogging)
	testSettings.LogLocation = ""
	dms3libs.SetLogFileLocation(testSettings)
	t.Log("log location set to", testSettings.LogLocation, "succesfully")

}
