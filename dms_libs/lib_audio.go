package dmslibs

// PlayAudio uses the shell aplay command (system default) to play AudioFile
func PlayAudio(AudioFile string) {

	_, err := RunCommand(SysCommands["APLAY"] + " -q " + AudioFile)

	if err != nil {
		Info.Println("audio file not found")
	}
}
