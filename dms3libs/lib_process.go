// Package dms3libs process provides process-related services for dms3 device components
package dms3libs

import (
	"errors"
	"os/exec"
)

var (
	ErrInvalidConfiguration = errors.New("invalid configuration")
	ErrNoCommand            = errors.New("empty command provided")
)

// isRunning checks if application is currently running (has PID > 0)
func isRunning(application string) bool {

	LogInfo("Check if already running: " + application)
	cmd := LibConfig.SysCommands["PGREP"] + " -if " + "'" + application + "'"

	var err error

	if _, err = RunCommand(cmd); err == nil {
		LogInfo("Already running: " + application)
		return true
	}

	handleCommandErrors(err, cmd)

	return false
}

// handleCommandErrors handles errors encountered while running commands
func handleCommandErrors(err error, cmd string) {

	switch err.(type) {
	case *exec.ExitError:
		LogInfo("Process not found when running " + cmd)
	default:
		LogFatal("Failed to run " + cmd + ": " + err.Error())
	}

}

// startApplication starts the specified application if it is not already running
func startApplication(application string) bool {

	// check already running
	if isRunning(application) {
		return false
	}

	LogInfo("Attempting to start: " + application)

	var err error

	if _, err = RunCommand(application); err == nil {
		LogInfo("Successfully started: " + application)
		return true
	}

	handleCommandErrors(err, application)

	return false
}

// stopApplication stops the specified application if it is currently running
func stopApplication(application string) bool {

	// check already stopped
	if !isRunning(application) {
		return false
	}

	LogInfo("Attempting to stop: " + application)
	cmd := LibConfig.SysCommands["PKILL"] + " -if " + "'" + application + "'"

	var err error

	if _, err = RunCommand(cmd); err == nil {
		LogInfo("Successfully stopped: " + application)
		return true
	}

	handleCommandErrors(err, cmd)

	return false
}

// RunCommand is a simple wrapper for the exec.Command() call
//
// NOTE: this call is blocking (non-threaded)
func RunCommand(cmd string) (res []byte, err error) {

	LogInfo("Command to be run: " + LibConfig.SysCommands["BASH"] + " -c " + cmd)

	// nolint (gosec): this is a closed application (not a package/library)
	return exec.Command(LibConfig.SysCommands["BASH"], "-c", cmd).Output()
}

// StartStopApplication enables/disables the application passed in based on state
func StartStopApplication(state MotionDetectorState, application string) bool {

	switch state {
	case Start:
		return startApplication(application)
	case Stop:
		return stopApplication(application)
	default:
		LogInfo("Unanticipated application state: ignored")
		return false
	}

}
