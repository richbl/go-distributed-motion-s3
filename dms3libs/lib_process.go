// Package dms3libs process provides process-related services for dms3 device components
//
package dms3libs

import (
	"os/exec"
	"strconv"
)

// RunCommand is a simple wrapper for the exec.Command() call
// NOTE: this call is blocking (non-threaded), and will return only after the command
// completes
//
func RunCommand(cmd string) (res []byte, err error) {
	return exec.Command(LibConfig.SysCommands["BASH"], "-c", cmd).Output()
}

// IsRunning checks if application is currently running (has PID > 0)
func IsRunning(application string) bool {

	if res, err := RunCommand(LibConfig.SysCommands["PGREP"] + " -c " + application); err != nil {
		LogInfo("Failed to run '" + LibConfig.SysCommands["PGREP"] + " -c " + application + "': " + err.Error())
		return false
	} else {
		count, _ := strconv.Atoi(string(StripRet(res)))
		return count > 0
	}

}

// StartStopApplication enable/disables the application passed in
func StartStopApplication(state MotionDetectorState, application string) bool {

	switch state {
	case Start:
		{

			if IsRunning(application) {
				return false // already running
			}

			if _, err := RunCommand(application); err != nil {
				LogInfo("Failed to start " + application + ": " + err.Error())
				return false
			} else {
				return true
			}

		}
	case Stop:
		{

			if !IsRunning(application) {
				return false // already stopped
			}

			if _, err := RunCommand(LibConfig.SysCommands["PKILL"] + " " + application); err == nil {
				return true
			} else {
				LogInfo("Failed to stop running process: " + application)
				return false
			}
		}
	default:
		{
			LogInfo("Unanticipated application state: ignored")
			return false
		}
	}

}
