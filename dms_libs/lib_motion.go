package dmslibs

import (
	"os"
	"strconv"
)

// ConfirmMotionInstall confirm that motion is installed
func ConfirmMotionInstall() bool {
	res, err := RunCommand("command -v motion")

	if err != nil {
		Info.Println("command failed")
	}

	return (len(string(res)) > 0)
}

// RunningMotion determines whether motion is running, returning PID (0 if no process)
func RunningMotion() int {
	res, err := RunCommand("pgrep -x code | grep -m1 ''")

	if err != nil {
		return 0
	}

	i, _ := strconv.Atoi(string(res[:len(res)-1]))
	return i
}

// MotionDaemon enable/disables the motion daemon
func MotionDaemon(command string) bool {

	switch command {
	case "start":
		{
			if RunningMotion() > 0 {
				return false
			}
			RunCommand("motion")
		}
	case "stop":
		{
			motionPID := RunningMotion()

			if !(RunningMotion() > 0) {
				return false
			}

			// find motion process and kill it
			proc, err := os.FindProcess(motionPID)

			if err != nil {
				Info.Println("command failed")
			} else {
				proc.Kill()
			}
		}
	}
	return true
}
