package client

import (
	"go_server/dms_libs"
)

func init() {
	confirmMotionInstall()
}

// confirmMotionInstall confirms that the motion program is installed
func confirmMotionInstall() {

}

// ProcessMotionState processes state received from server
func ProcessMotionState(motionState string) {

	dmslibs.PrintFuncName()

	switch motionState {
	case "enable":
		StartMotionDaemon()
	case "disable":
		StopMotionDaemon()
	default:
		dmslibs.Info.Println("Unanticipated server response: " + motionState)
	}

}

// StartMotionDaemon comment
func StartMotionDaemon() {

	// TODO change string to token/type?
	if MotionDaemon("start") {
		dmslibs.Info.Println("Motion started")
	}

}

// StopMotionDaemon comment
func StopMotionDaemon() {
	if MotionDaemon("stop") {
		dmslibs.Info.Println("Motion stopped")
	}
}
