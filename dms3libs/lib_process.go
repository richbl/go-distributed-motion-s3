// Package dms3libs process provides process-related services for dms3 device components
//
package dms3libs

import (
	"os/exec"
)

// RunCommand is a simple wrapper for the exec.Command() call
//
// NOTE: this call is blocking (non-threaded), and will return only after the command
// completes
//
func RunCommand(cmd string) (res []byte, err error) {
	return exec.Command(LibConfig.SysCommands["BASH"], "-c", cmd).Output()
}

// IsRunning checks if application is currently running (has PID > 0)
//
func IsRunning(application string) bool {

	if _, err := RunCommand(LibConfig.SysCommands["PGREP"] + " -i " + application); err != nil {

		switch err.(type) {
		case *exec.ExitError: // no process found
			LogInfo("Process not found when running '" + LibConfig.SysCommands["PGREP"] + " -i " + application)
		default: // fatal command error
			LogFatal("Failed to run '" + LibConfig.SysCommands["PGREP"] + " -i " + application + "': " + err.Error())
		}
		return false
	} else {
		return true
	}

}

// StartStopApplication enable/disables the application passed in
//
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

			if _, err := RunCommand(LibConfig.SysCommands["PKILL"] + " -i " + application); err == nil {
				return true
			} else {

				switch err.(type) {
				case *exec.ExitError: // no process matched
					LogInfo("Process not found when running '" + LibConfig.SysCommands["PKILL"] + " -i " + application)
				default: // fatal command error
					LogFatal("Failed to run '" + LibConfig.SysCommands["PKILL"] + " -i " + application + "': " + err.Error())
				}
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
