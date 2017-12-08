// Package dms3client manager processes dms3client device component messages passed to it by
// the dms3server device component
//
package dms3client

import (
	"go-distributed-motion-s3/dms3libs"
)

// ProcessMotionDetectorState starts/stops the motion detector application
func ProcessMotionDetectorState() {

	dms3libs.LogDebug(dms3libs.GetFunctionName())

	cmdStr := ""
	state := dms3libs.MotionDetector.State()

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
