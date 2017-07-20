package dmslibs

// PlayAudio uses the shell aplay command (system default) to play audioFile
func PlayAudio(audioFile string) {

	if _, err := RunCommand(SysCommand["APLAY"] + " -q " + audioFile); err != nil {
		Info.Println(err.Error())
	}

}
