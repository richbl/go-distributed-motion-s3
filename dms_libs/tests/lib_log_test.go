package dmslibs_test

import (
	"go_server/dms_libs"
	"os"
	"testing"
)

func TestCreateLogger(t *testing.T) {

	testLog := "lib_log_test.log"

	dmslibs.CreateLogger(1, dmslibs.GetPackageDir(), testLog)
	if !dmslibs.IsFile(testLog) {
		t.Error("Log file", testLog, "not created")
	} else {
		os.Remove(testLog)
	}
}
