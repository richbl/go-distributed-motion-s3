package dmslibs

import (
	"fmt"
	"os/exec"
)

// ConfirmMotionInstall confirm that motion is installed
func ConfirmMotionInstall() bool {
	exec.Command("/bin/bash", "-q", "type motion > /dev/null 2>&1")
	// TODO exec.Command success/fail?
	return true
}

// RunningMotion determines whether motion is running, returning PID
func RunningMotion() int {
	// TODO
	return 0
}

// MotionDaemon enable/disables the motion daemon
func MotionDaemon(command string) bool {

	switch command {
	case "start":
		{
			if RunningMotion() > 0 {
				return false
			}
			// TODO we need first argument? (MOTION is absolute path)
			exec.Command("/bin/bash", "-q", SysCommands["MOTION"]+" > /dev/null 2>&1")
		}
	case "stop":
		{
			motionPID := RunningMotion()
			if !(RunningMotion() > 0) {
				return false
			} else {
				// TODO
				fmt.Println(motionPID)
				// Process.kill('KILL', motion_pid)
			}

		}
	}
	return true
}
