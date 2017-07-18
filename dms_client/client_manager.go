package client

import (
	"go_server/dms_libs"
)

// ProcessMotionDetectorState processes the application state received from the server
func ProcessMotionDetectorState(state dmslibs.MotionDetectorState) {
	dmslibs.LogDebug(dmslibs.GetFunctionName())

	switch state {
	case dmslibs.Start, dmslibs.Stop:
		startStopMotionDetector(state)
	default:
		dmslibs.LogInfo("Unanticipated response from server: politely ignored")
	}

}

// startStopMotionDetector starts/stops the motion detector application
func startStopMotionDetector(value dmslibs.MotionDetectorState) {
	cmdStr := " started"

	if value == dmslibs.Stop {
		cmdStr = " stopped"
	}

	if dmslibs.StartStopApplication(value, dmslibs.MotionDetector.Command) {
		dmslibs.LogInfo(dmslibs.MotionDetector.Command + cmdStr)
	}

}
