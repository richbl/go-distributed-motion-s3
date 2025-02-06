package dms3libs_test

import (
	"os"
	"testing"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

const (
	logFile = "tmpFile"
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
		t.Error(logFile, logger.LogFilename, "not created")
	} else {
		t.Log(logFile, logger.LogFilename, "created")
		os.Remove(logger.LogFilename)
		t.Log(logFile, logger.LogFilename, "removed")
	}

}
