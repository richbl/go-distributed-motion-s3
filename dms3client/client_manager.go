package dms3client

import (
	"go_server/dms3libs"
)

// ProcessMotionDetectorState processes the application state received from the server
func ProcessMotionDetectorState(state dms3libs.MotionDetectorState) {

	dms3libs.LogDebug(dms3libs.GetFunctionName())

	switch state {
	case dms3libs.Start, dms3libs.Stop:
		startStopMotionDetector(state)
	default:
		dms3libs.LogInfo("Unanticipated motion detector state: ignored")
	}

}

// startStopMotionDetector starts/stops the motion detector application
func startStopMotionDetector(state dms3libs.MotionDetectorState) {

	dms3libs.LogDebug(dms3libs.GetFunctionName())
	var cmdStr string

	switch state {
	case dms3libs.Start:
		cmdStr = "started"
	case dms3libs.Stop:
		cmdStr = "stopped"
	default:
		dms3libs.LogInfo("Unanticipated motion detector state: ignored")
	}

	motionCommand := dms3libs.MotionDetector.Command()

	if dms3libs.StartStopApplication(state, motionCommand) {
		dms3libs.LogInfo(motionCommand + " " + cmdStr)
	}

}
