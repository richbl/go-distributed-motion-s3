package dms3libs_test

import (
	"path/filepath"
	"testing"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

func init() {
	dms3libs.LoadLibConfig(filepath.Join("..", "..", dms3libs.DMS3Config, dms3libs.DMS3libsTOML))
}

func TestRunCommand(t *testing.T) {

	testCommand := dms3libs.LibConfig.SysCommands["ENV"]

	if res, err := dms3libs.RunCommand(testCommand); err != nil {
		t.Error("Command " + testCommand + " failed")
	} else if len(res) == 0 {
		t.Error("Output from command " + testCommand + " failed")
	}

}

func TestStartStopApplication(t *testing.T) {

	// NOTE: assumes Motion/MotionPlus application (https://motion-project.github.io/) is installed
	// and properly configured
	testApplication := dms3libs.LibConfig.SysCommands["MOTION"]

	if !dms3libs.StartStopApplication(dms3libs.Start, testApplication) {
		t.Error("Start failed")
	}

	if !dms3libs.StartStopApplication(dms3libs.Stop, testApplication) {
		t.Error("Stop failed")
	}

}
