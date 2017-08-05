package dms3libs_test

import (
	"go-distributed-motion-s3/dms3libs"
	"os"
	"testing"
)

func TestCreateLogger(t *testing.T) {

	testLog := "lib_log_test.log"
	dms3libs.CreateLogger(4, 1, dms3libs.GetPackageDir(), testLog)

	if !dms3libs.IsFile(testLog) {
		t.Error("Log file", testLog, "not created")
	} else {
		t.Log("Log file", testLog, "created")
		os.Remove(testLog)
		t.Log("Log file", testLog, "removed")
	}

}
