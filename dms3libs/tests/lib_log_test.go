package dms3libs_test

import (
	"go-distributed-motion-s3/dms3libs"
	"os"
	"testing"
)

func TestCreateLogger(t *testing.T) {

	logger := dms3libs.StructLogging{
		LogLevel:    2,
		LogDevice:   1,
		LogFilename: "lib_log_test.log",
		LogLocation: dms3libs.GetPackageDir(),
	}

	dms3libs.CreateLogger(&logger)

	if !dms3libs.IsFile(logger.LogFilename) {
		t.Error("Log file", logger.LogFilename, "not created")
	} else {
		t.Log("Log file", logger.LogFilename, "created")
		os.Remove(logger.LogFilename)
		t.Log("Log file", logger.LogFilename, "removed")
	}

}
