package client

import (
	"go_server/dms_libs"
)

// ProcessAppState processes the application state received from server
func ProcessAppState(appState string) {

	dmslibs.PrintFuncName()

	switch appState {
	case "enable":
		StartAppDaemon()
	case "disable":
		StopAppDaemon()
	default:
		dmslibs.Info.Println("Unanticipated server response: " + appState)
	}

}

// StartAppDaemon comment
func StartAppDaemon() {
	if dmslibs.AppDaemon("start", dmslibs.SysCommands["SURVEILLANCE_CMD"]) {
		dmslibs.Info.Println(dmslibs.SysCommands["SURVEILLANCE_CMD"] + " started")
	}
}

// StopAppDaemon comment
func StopAppDaemon() {
	if dmslibs.AppDaemon("stop", dmslibs.SysCommands["SURVEILLANCE_CMD"]) {
		dmslibs.Info.Println(dmslibs.SysCommands["SURVEILLANCE_CMD"] + " stopped")
	}
}
