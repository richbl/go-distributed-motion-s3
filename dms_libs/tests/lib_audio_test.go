package dmslibs_test

import (
	"go_server/dms_libs"
	"go_server/dms_server"
	"testing"
)

func TestPlayAudio(t *testing.T) {

	testFile := "lib_audio_test.wav"

	if dmslibs.IsFile(testFile) {
		dmslibs.PlayAudio(testFile)
	} else {
		t.Error("Test file", testFile, "not found")
	}

}

func TestAudioConfig(t *testing.T) {

	if dmslibs.IsFile(server.AudioMotionStart) {
		dmslibs.PlayAudio(server.AudioMotionStart)
	} else {
		t.Error("Audio file", server.AudioMotionStart, "not found")
	}

	if dmslibs.IsFile(server.AudioMotionStop) {
		dmslibs.PlayAudio(server.AudioMotionStop)
	} else {
		t.Error("Audio file", server.AudioMotionStop, "not found")
	}

}
