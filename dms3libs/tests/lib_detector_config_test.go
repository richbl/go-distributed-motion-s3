package dms3libs_test

import (
	"testing"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

const (
	errFunctionFailed = "function failed"
)

func TestCommand(t *testing.T) {

	if dms3libs.LibConfig.SysCommands["MOTION"] == "" {
		t.Error(errFunctionFailed)
	}

}

func TestState(t *testing.T) {

	dms3libs.MotionDetector.SetState(dms3libs.Start)

	if dms3libs.MotionDetector.State() != dms3libs.Start {
		t.Error(errFunctionFailed)
	}

}

func TestSetState(t *testing.T) {

	if !dms3libs.MotionDetector.SetState(dms3libs.Start) {
		t.Error(errFunctionFailed)
	}

}
