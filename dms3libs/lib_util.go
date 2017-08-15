package dms3libs

import (
	"fmt"
	"log"
	"os"
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

// StripRet strips the rightmost byte from the byte array
func StripRet(value []byte) []byte {

	if len(value) <= 1 {
		return value
	}

	return value[:len(value)-1]

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

// Format24H formats 24-hour time to six places (HHMMSS)
func Format24H(time string) string {
	return rightPadToLen(time, "0", 6)
}

func rightPadToLen(s string, padStr string, pLen int) string {
	return s + strings.Repeat(padStr, pLen-len(s))
}

// CheckErr does simple error management (no logging dependencies)
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
