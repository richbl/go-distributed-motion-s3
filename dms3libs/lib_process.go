package dms3libs

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// RunCommand is a simple wrapper for the exec.Command() call
// NOTE: this call is blocking (non-threaded), and will return only after the command
// completes
//
func RunCommand(cmd string) (res []byte, err error) {
	return exec.Command("bash", "-c", cmd).Output()
}

// IsRunning checks if application is currently running (has PID > 0)
func IsRunning(application string) bool {
	return (GetPID(application) > 0)
}

// GetPIDCount returns the count of application PIDs
func GetPIDCount(application string) int {

	res, _ := RunCommand(LibConfig.SysCommands["PGREP"] + " -x -c " + application)
	count, _ := strconv.Atoi(string(StripRet(res)))
	return count

}

// GetPIDList returns application PIDs (0 if no process)
func GetPIDList(application string) (int, []int) {

	pidCount := GetPIDCount(application)

	switch pidCount {
	case 0: // no process running
		return 0, []int{0}
	default: // one or more processes running
		{
			res, _ := RunCommand(LibConfig.SysCommands["PGREP"] + " -x " + application)
			strPIDs := strings.Split(string(StripRet(res)), "\n")

			PIDs := []int{}
			for _, i := range strPIDs {
				pid, _ := strconv.Atoi(i)
				PIDs = append(PIDs, pid)
			}
			return pidCount, PIDs
		}
	}

}

// GetPID returns the application PID (0 if no process)
func GetPID(application string) int {

	pidCount, pidList := GetPIDList(application)

	switch pidCount {
	case 0, 1:
		return pidList[0]
	default: // >1 processes running
		{
			Fatal.Fatalln("multiple instances of " + application + " process running")
			return 0
		}
	}

}

// StartStopApplication enable/disables the application passed in
func StartStopApplication(state MotionDetectorState, application string) bool {

	switch state {
	case Start:
		{
			if IsRunning(application) {
				return false // already running
			}

			_, err := RunCommand(application)

			if err != nil {
				LogInfo("Failed to run " + application + ": " + err.Error())
				return false
			}

			return true
		}
	case Stop:
		{
			if !IsRunning(application) {
				return false // already stopped
			}

			// find application process and kill it
			appPID := GetPID(application)
			proc, err := os.FindProcess(appPID)

			if err != nil {
				LogInfo("unable to find PID")
				return false
			}

			proc.Kill()
			return true
		}
	default:
		{
			LogInfo("Unanticipated motion detector state: ignored")
			return false
		}
	}

}

// GetArgs parses the command line, returning an array of args
func GetArgs() []string {

	if len(os.Args) == 1 {
		LogFatal("no arguments passed")
	}

	return os.Args

}
