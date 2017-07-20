package dmslibs

import (
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// GetFunctionName uses reflection (runtime) to return current function name
func GetFunctionName() string {
	pc := make([]uintptr, 10)

	// get program counter index (call stack)
	runtime.Callers(2, pc)
	fn := runtime.FuncForPC(pc[0])
	return fn.Name()
}

// GetPackageDir returns the absolute path of the calling package
func GetPackageDir() string {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		log.Fatal()
	}

	return path.Dir(filename)
}

// IsFile returns true/false on existence of file passed in
func IsFile(filename string) bool {

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

// RunCommand is a simple wrapper for the exec.Command() call
// Note that this call is blocking, and will return only after the command completes
// (non-threaded)
func RunCommand(cmd string) (res []byte, err error) {
	return exec.Command("bash", "-c", cmd).Output()
}

// IsRunning checks if application is currently running (has PID > 0)
func IsRunning(application string) bool {
	return (GetPID(application) > 0)
}

// StripRet strips the rightmost byte from the byte array
func StripRet(value []byte) []byte {

	if len(value) <= 1 {
		return value
	}

	return value[:len(value)-1]
}

// GetPIDCount returns the count of application PIDs
func GetPIDCount(application string) int {
	res, _ := RunCommand(SysCommand["PGREP"] + " -x -c " + application)
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
			res, _ := RunCommand(SysCommand["PGREP"] + " -x " + application)
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
//
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

			RunCommand(application)
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

// GetCurTime returns the current time as int (in 24-hour format, e.g., 231305)
func GetCurTime() int {
	curTime, _ := strconv.Atoi(To24H(time.Now()))
	return curTime
}

// To24H converts 12-hour time to 24-hour time, returning a string (e.g., "231305")
func To24H(value time.Time) string {
	return value.Format("150405")
}
