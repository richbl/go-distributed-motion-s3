package dms3libs_test

import (
	"go-distributed-motion-s3/dms3libs"
	"testing"
)

func init() {
	dms3libs.LoadLibConfig("../../config/dms3libs.toml")
}

func TestCommand(t *testing.T) {

	if dms3libs.MotionDetector.Command() != "motion" {
		t.Error("command failed")
	}

}

func TestSetState(t *testing.T) {

	if !dms3libs.MotionDetector.SetState(dms3libs.Start) {
		t.Error("command failed")
	}

}

func TestState(t *testing.T) {

	dms3libs.MotionDetector.SetState(dms3libs.Start)

	if dms3libs.MotionDetector.State() != dms3libs.Start {
		t.Error("command failed")
	}

}
