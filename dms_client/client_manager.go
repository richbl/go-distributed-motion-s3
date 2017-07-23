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
		dmslibs.LogInfo("Unanticipated motion detector state: ignored")
	}

}

// startStopMotionDetector starts/stops the motion detector application
func startStopMotionDetector(state dmslibs.MotionDetectorState) {

	dmslibs.LogDebug(dmslibs.GetFunctionName())
	var cmdStr string

	switch state {
	case dmslibs.Start:
		cmdStr = "started"
	case dmslibs.Stop:
		cmdStr = "stopped"
	default:
		dmslibs.LogInfo("Unanticipated motion detector state: ignored")
	}

	motionCommand := dmslibs.MotionDetector.Command()

	if dmslibs.StartStopApplication(state, motionCommand) {
		dmslibs.LogInfo(motionCommand + " " + cmdStr)
	}

}
