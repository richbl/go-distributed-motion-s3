package dmslibs

import (
	"os/exec"
)

// PlayAudio uses the shell aplay command (system default) to play AudioFile
func PlayAudio(AudioFile string) {
	_, err := exec.Command(SysCommands["APLAY"], "-q", AudioFile).Output()

	if err != nil {
		Info.Println("audio file not found")
	}
}
