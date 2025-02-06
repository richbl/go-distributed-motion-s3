package dms3libs_test

import (
	"path/filepath"
	"testing"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
	"github.com/richbl/go-distributed-motion-s3/dms3server"
)

const (
	successfulPlayback = "successful playback"
	errNotFound        = "not found"
	audioFile          = "audio file"
	testFile           = "Test file"
)

func init() {
	dms3libs.LoadLibConfig(filepath.Join("..", "..", dms3libs.DMS3Config, dms3libs.DMS3libsTOML))
}

func TestPlayAudio(t *testing.T) {

	testFile := "lib_audio_test.wav"

	if dms3libs.IsFile(testFile) {
		dms3libs.PlayAudio(testFile)
		t.Log(testFile, testFile, successfulPlayback)
	} else {
		t.Error(testFile, testFile, errNotFound)
	}

}

func TestAudioConfig(t *testing.T) {

	configPath := dms3libs.GetPackageDir()

	dms3libs.LoadComponentConfig(&dms3server.ServerConfig, filepath.Join(configPath, "..", "..", dms3libs.DMS3Config, dms3libs.DMS3serverTOML))

	mediaFileStart := dms3server.ServerConfig.Audio.PlayMotionStart
	mediaFileStop := dms3server.ServerConfig.Audio.PlayMotionStop

	if mediaFileStart == "" {
		mediaFileStart = filepath.Join("..", "..", dms3libs.DMS3Server, "media", "motion_start.wav")
	}

	if dms3libs.IsFile(mediaFileStart) {
		dms3libs.PlayAudio(mediaFileStart)
		t.Log(audioFile, mediaFileStart, successfulPlayback)
	} else {
		t.Error(audioFile, mediaFileStart, errNotFound)
	}

	if mediaFileStop == "" {
		mediaFileStop = filepath.Join("..", "..", dms3libs.DMS3Server, "media", "motion_stop.wav")
	}

	if dms3libs.IsFile(mediaFileStop) {
		dms3libs.PlayAudio(mediaFileStop)
		t.Log(audioFile, mediaFileStop, successfulPlayback)
	} else {
		t.Error(audioFile, mediaFileStop, errNotFound)
	}

}
