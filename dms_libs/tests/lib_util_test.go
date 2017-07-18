package dmslibs_test

import (
	"go_server/dms_libs"
	"testing"
)

func TestPrintFuncName(t *testing.T) {
	dmslibs.LogDebug(dmslibs.GetFunctionName())
}

func TestGetPackageDir(t *testing.T) {
	dmslibs.GetPackageDir()
}

func TestIsFile(t *testing.T) {
	testFile := "lib_util_test.go"

	if !dmslibs.IsFile(testFile) {
		t.Error(testFile + " file not found, but should have been")
	}
}

func TestRunCommand(t *testing.T) {
	testCommand := "ls"

	res, err := dmslibs.RunCommand(testCommand)

	if err != nil {
		t.Error("command " + testCommand + " failed")
	}

	if len(res) == 0 {
		t.Error("output from command " + testCommand + " failed")
	}
}

func TestIsRunning(t *testing.T) {
	// ACTION: set to known active process
	testApplication := "gocode"

	if !dmslibs.IsRunning(testApplication) {
		t.Error(testApplication + " command not running")
	}
}

func TestStripRet(t *testing.T) {
	testArray := []byte{50, 40, 30, 20, 10}

	res := dmslibs.StripRet(testArray)

	if len(res) != len(testArray)-1 {
		t.Error("command failed")
	}
}

func TestGetPIDCount(t *testing.T) {
	// ACTION: set to known active process
	testApplication := "gocode"

	if dmslibs.GetPIDCount(testApplication) < 1 {
		t.Error("command failed")
	}
}

func TestGetPID(t *testing.T) {
	// ACTION: set to known active process
	testApplication := "gocode"

	if dmslibs.GetPID(testApplication) == 0 {
		t.Error("command failed")
	}
}
