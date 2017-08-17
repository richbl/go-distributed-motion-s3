package dms3libs_test

import (
	"go-distributed-motion-s3/dms3libs"
	"go-distributed-motion-s3/dms3server"
	"testing"
)

func init() {
	dms3libs.LoadLibConfig("../../dms3libs/lib_config.toml")
}

func TestPlayAudio(t *testing.T) {

	testFile := "lib_audio_test.wav"

	if dms3libs.IsFile(testFile) {
		dms3libs.PlayAudio(testFile)
		t.Log("Test file", testFile, "played successfully")
	} else {
		t.Error("Test file", testFile, "not found")
	}

}

func TestAudioConfig(t *testing.T) {

	dms3server.LoadServerConfig("../../dms3server/server_config.toml")

	if dms3libs.IsFile(dms3server.ServerConfig.PlayMotionStart) {
		dms3libs.PlayAudio(dms3server.ServerConfig.PlayMotionStart)
		t.Error("Audio file", dms3server.ServerConfig.PlayMotionStart, "played successfully")
	} else {
		t.Error("Audio file", dms3server.ServerConfig.PlayMotionStart, "not found")
	}

	if dms3libs.IsFile(dms3server.ServerConfig.PlayMotionStop) {
		dms3libs.PlayAudio(dms3server.ServerConfig.PlayMotionStop)
		t.Error("Audio file", dms3server.ServerConfig.PlayMotionStop, "played successfully")
	} else {
		t.Error("Audio file", dms3server.ServerConfig.PlayMotionStop, "not found")
	}

}
