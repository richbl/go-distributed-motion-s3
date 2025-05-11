// Package dms3libs util provides utility services for dms3 device components
package dms3libs

import (
	"fmt"
	"image"
	_ "image/jpeg" // Import JPEG decoder
	"log"
	"os"
	"path"
	"runtime"
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

// Uptime returns uptime for the application process
func Uptime(startTime time.Time) string {
	return fmtDuration(time.Since(startTime))
}

// SecondsSince returns seconds passed since value passed
func SecondsSince(value time.Time) uint32 {
	return uint32(time.Since(value).Seconds())
}

// To24H converts 12-hour time to 24-hour time, returning a string (e.g., "231305")
func To24H(value time.Time) string {
	return value.Format("150405")
}

// Format24H formats 24-hour time to six places (HHMMSS)
func Format24H(time string) string {
	return rightPadToLen(time, "0", 6)
}

// FormatDateTime formats time to "date at time"
func FormatDateTime(value time.Time) string {
	return value.Format("2006-01-02 at 15:04:05")
}

// ModVal returns the remainder of number/val passed in
func ModVal(number int, val int) int {
	return number % val
}

// CheckErr does simple error management (no logging dependencies)
func CheckErr(err error) {

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}

// GetProjectVersion returns the current project version string to the caller
func GetProjectVersion() string {
	return "1.4.4"
}

// GetImageDimensions returns the (width, height) of an image passed in
func GetImageDimensions(imagePath string) (int, int) {

	var file *os.File
	var err error
	var img image.Config

	if file, err = os.Open(imagePath); err != nil {
		LogFatal(err.Error())
	}

	defer file.Close()

	if img, _, err = image.DecodeConfig(file); err != nil {
		LogFatal(err.Error())
	}

	return img.Width, img.Height

}

// rightPadToLen pads a string to pLen places with padStr
func rightPadToLen(s string, padStr string, pLen int) string {
	return s + strings.Repeat(padStr, pLen-len(s))
}

// fmtDuration formats a duration value to show days/hours/minutes/seconds
func fmtDuration(d time.Duration) string {

	days := (d / time.Hour) / 24
	d -= days * (time.Hour * 24)
	hours := d / time.Hour
	d -= hours * time.Hour
	minutes := d / time.Minute
	d -= minutes * time.Minute
	seconds := d / time.Second

	return fmt.Sprintf("%03dd:%02dh:%02dm:%02ds", days, hours, minutes, seconds)
}
