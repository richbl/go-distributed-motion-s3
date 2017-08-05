package dms3libs

// PlayAudio uses the shell aplay command (system default) to play audioFile
func PlayAudio(audioFile string) {

	if _, err := RunCommand(LibConfig.SysCommands["APLAY"] + " -q " + audioFile); err != nil {
		LogInfo(err.Error())
	}

}
