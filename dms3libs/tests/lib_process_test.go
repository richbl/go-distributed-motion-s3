package dms3libs_test

import (
	"path/filepath"
	"testing"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

func init() {
	dms3libs.LoadLibConfig(filepath.Join("..", "..", "config", "dms3libs.toml"))
}

func TestRunCommand(t *testing.T) {

	testCommand := dms3libs.LibConfig.SysCommands["ARP"]

	if res, err := dms3libs.RunCommand(testCommand); err != nil {
		t.Error("Command " + testCommand + " failed")
	} else if len(res) == 0 {
		t.Error("Output from command " + testCommand + " failed")
	}

}

func TestIsRunning(t *testing.T) {

	// ACTION: set to known active process
	testApplication := "gopls"

	if !dms3libs.IsRunning(testApplication) {
		t.Error(testApplication + " command not running")
	}

}

func TestStartStopApplication(t *testing.T) {

	// ACTION: set to known installed application configured to run as service
	// NOTE: assumes motion program (https://motion-project.github.io/) is installed
	//
	testApplication := "motion"

	if !dms3libs.StartStopApplication(dms3libs.Start, testApplication) {
		t.Error("start failed")
	}

	if !dms3libs.StartStopApplication(dms3libs.Stop, testApplication) {
		t.Error("stop failed")
	}

}
