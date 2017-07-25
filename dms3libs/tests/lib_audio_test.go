package dms3libs_test

import (
	"go_server/dms3libs"
	"go_server/dms3server"
	"testing"
)

func TestPlayAudio(t *testing.T) {

	testFile := "lib_audio_test.wav"

	if dms3libs.IsFile(testFile) {
		dms3libs.PlayAudio(testFile)
	} else {
		t.Error("Test file", testFile, "not found")
	}

}

func TestAudioConfig(t *testing.T) {

	if dms3libs.IsFile(dms3server.AudioMotionDetectorStart) {
		dms3libs.PlayAudio(dms3server.AudioMotionDetectorStart)
	} else {
		t.Error("Audio file", dms3server.AudioMotionDetectorStart, "not found")
	}

	if dms3libs.IsFile(dms3server.AudioMotionDetectorStop) {
		dms3libs.PlayAudio(dms3server.AudioMotionDetectorStop)
	} else {
		t.Error("Audio file", dms3server.AudioMotionDetectorStop, "not found")
	}

}
