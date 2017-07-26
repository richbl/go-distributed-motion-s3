package dms3libs_test

import (
	"go-distributed-motion-s3/dms3libs"
	"testing"
)

func TestPrintFuncName(t *testing.T) {

	dms3libs.LogDebug(dms3libs.GetFunctionName())

}

func TestGetPackageDir(t *testing.T) {

	dms3libs.GetPackageDir()

}

func TestIsFile(t *testing.T) {

	testFile := "lib_util_test.go"

	if !dms3libs.IsFile(testFile) {
		t.Error(testFile + " file not found, but should have been")
	}

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

func TestStripRet(t *testing.T) {

	testArray := []byte{50, 40, 30, 20, 10}

	res := dms3libs.StripRet(testArray)

	if len(res) != len(testArray)-1 {
		t.Error("command failed")
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
