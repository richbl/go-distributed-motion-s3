package client

import (
	"fmt"
	"go_server/libs"
)

// confirmMotionInstall comments
func confirmMotionInstall() {

}

// ProcessMotionState comments
func ProcessMotionState(motionState string) {

	libconfig.PrintFunctionName()

	switch motionState {
	case "enable":
		{
			fmt.Println("ENABLED")
		}
	case "disable":
		{
			fmt.Println("DISABLED")
		}
	default:
		{
			fmt.Println("!Unanticipated response")
		}
	}

}
