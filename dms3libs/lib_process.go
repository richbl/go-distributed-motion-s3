// Package dms3libs process provides process-related services for dms3 device components
package dms3libs

import (
	"os/exec"
)

// RunCommand is a simple wrapper for the exec.Command() call
//
// NOTE: this call is blocking (non-threaded), and will return only after the command
// completes
func RunCommand(cmd string) (res []byte, err error) {

	LogInfo("Command to be run: " + LibConfig.SysCommands["BASH"] + " -c " + cmd)
	return exec.Command(LibConfig.SysCommands["BASH"], "-c", cmd).Output()
}

// IsRunning checks if application is currently running (has PID > 0)
func IsRunning(application string) bool {

	cmd := LibConfig.SysCommands["PGREP"] + " -if " + application

	if _, err := RunCommand(cmd); err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			LogInfo("Process not found when running '" + cmd + "'")
			return false
		}
		LogFatal("Failed to run '" + cmd + "': " + err.Error())
		return false
	}
	return true

}

// StartStopApplication enable/disables the application passed in
func StartStopApplication(state MotionDetectorState, application string) bool {

	switch state {
	case Start:
		{

			if IsRunning(application) {
				LogInfo("Already running: " + application)
				return false
			}

			if _, err := RunCommand(application); err != nil {
				LogInfo("Failed to start " + application + ": " + err.Error())
				return false
			} else {
				LogInfo("Successfully started " + application)
				return true
			}

		}
	case Stop:
		{

			if !IsRunning(application) {
				return false // already stopped
			}

			cmd := LibConfig.SysCommands["PKILL"] + " -if " + application

			if _, err := RunCommand(cmd); err == nil {
				return true
			} else {

				switch err.(type) {
				case *exec.ExitError: // no process matched
					LogInfo("Process not found when running '" + cmd + "'")
				default: // fatal command error
					LogFatal("Failed to run '" + cmd + "': " + err.Error())
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
