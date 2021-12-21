package dms3libs_test

import (
	"path/filepath"
	"testing"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
	"github.com/richbl/go-distributed-motion-s3/dms3server"
)

func init() {
	dms3libs.LoadLibConfig("../../config/dms3libs.toml")
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

	configPath := dms3libs.GetPackageDir()

	dms3libs.LoadComponentConfig(&dms3server.ServerConfig, filepath.Join(configPath, "../../config/dms3server.toml"))

	mediaFileStart := dms3server.ServerConfig.Audio.PlayMotionStart
	mediaFileStop := dms3server.ServerConfig.Audio.PlayMotionStop

	if mediaFileStart == "" {
		mediaFileStart = "../../dms3server/media/motion_start.wav"
	}

	if dms3libs.IsFile(mediaFileStart) {
		dms3libs.PlayAudio(mediaFileStart)
		t.Log("Audio file", mediaFileStart, "played successfully")
	} else {
		t.Error("Audio file", mediaFileStart, "not found")
	}

	if mediaFileStop == "" {
		mediaFileStop = "../../dms3server/media/motion_stop.wav"
	}

	if dms3libs.IsFile(mediaFileStop) {
		dms3libs.PlayAudio(mediaFileStop)
		t.Log("Audio file", mediaFileStop, "played successfully")
	} else {
		t.Error("Audio file", mediaFileStop, "not found")
	}

}
