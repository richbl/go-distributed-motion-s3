package dmslibs

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
)

// GetFunctionName uses reflection (runtime) to return current function name
func GetFunctionName() string {
	pc := make([]uintptr, 10)

	// get program counter index (call stack)
	// set to 3, as this function is wrapped by PrintFuncName()
	runtime.Callers(3, pc)
	fn := runtime.FuncForPC(pc[0])
	return fn.Name()
}

// PrintFuncName wraps GetFunctionName()
func PrintFuncName() {
	fmt.Println("BEGIN: " + GetFunctionName())
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
func RunCommand(cmd string) (res []byte, err error) {
	return exec.Command("bash", "-c", cmd).Output()
}
