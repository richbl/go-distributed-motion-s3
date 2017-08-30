package dms3libs_test

import (
	"go-distributed-motion-s3/dms3libs"
	"testing"
)

func init() {
	dms3libs.LoadLibConfig("../../config/dms3libs.toml")
}

func TestRunCommand(t *testing.T) {

	testCommand := "ls"

	if res, err := dms3libs.RunCommand(testCommand); err != nil {
		t.Error("command " + testCommand + " failed")
	} else if len(res) == 0 {
		t.Error("output from command " + testCommand + " failed")
	}

}

func TestIsRunning(t *testing.T) {

	// ACTION: set to known active process
	testApplication := "gocode"

	if !dms3libs.IsRunning(testApplication) {
		t.Error(testApplication + " command not running")
	}

}

func TestGetPIDCount(t *testing.T) {

	// ACTION: set to known active process
	testApplication := "gocode"

	if dms3libs.GetPIDCount(testApplication) < 1 {
		t.Error("command failed")
	}

}

func TestGetPID(t *testing.T) {

	// ACTION: set to known active process
	testApplication := "gocode"

	if dms3libs.GetPID(testApplication) == 0 {
		t.Error("command failed")
	}

}

func TestStartStopApplication(t *testing.T) {

	// ACTION: set to known installed application configured to run as service
	// NOTE: should be run with root permissions (if running as daemon)
	//
	testApplication := "motion"

	if !dms3libs.StartStopApplication(dms3libs.Start, testApplication) {
		t.Error("start failed")
	}

	if !dms3libs.StartStopApplication(dms3libs.Stop, testApplication) {
		t.Error("stop failed")
	}

}
