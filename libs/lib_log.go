package libconfig

import (
	"fmt"
	"runtime"
)

// GetFunctionName comment
func GetFunctionName() string {
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// PrintFunctionName comment
func PrintFunctionName() {
	fmt.Println("BEGIN #" + GetFunctionName())
}
