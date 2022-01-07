// Package dms3libs log provides logging services for dms3 device components
//
package dms3libs

import (
	"log"
	"os"
	"path/filepath"
)

// global logging objects
var (
	Fatal        *log.Logger
	Info         *log.Logger
	Debug        *log.Logger
	loggingLevel int
)

// StructLogging is used for dms3 logging
type StructLogging struct {
	LogLevel    int
	LogDevice   int
	LogFilename string
	LogLocation string
}

func init() {

	var f *os.File
	Fatal = log.New(f, "FATAL: ", log.Lshortfile|log.LstdFlags)
	Info = log.New(f, "INFO: ", log.Lshortfile|log.LstdFlags)
	Debug = log.New(f, "DEBUG: ", log.Lshortfile|log.LstdFlags)

}

// CreateLogger creates an application log file (1) or redirects to STDOUT (2) based on logDevice
func CreateLogger(logger *StructLogging) {

	var (
		f   *os.File
		err error
	)

	if logger.LogLevel > 0 {
		loggingLevel = logger.LogLevel

		switch logger.LogDevice {
		case 0:
			f = os.Stdout
		case 1:
			{
				if f, err = os.OpenFile(filepath.Join(logger.LogLocation, logger.LogFilename), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644); err != nil {
					LogFatal(err.Error())
				}
			}
		}

	}

	Fatal.SetOutput(f)
	Info.SetOutput(f)
	Debug.SetOutput(f)

}

// LogFatal generates a fatal log message based on loggingLevel
func LogFatal(msg string) {

	if loggingLevel >= 1 {
		Fatal.Fatalln(msg)
	}

	os.Exit(1)

}

// LogInfo generates an informational log message based on loggingLevel
func LogInfo(msg string) {

	if loggingLevel >= 2 {
		Info.Println(msg)
	}

}

// LogDebug generates a debug log message based on loggingLevel
func LogDebug(msg string) {

	if loggingLevel >= 4 {
		Debug.Println(msg)
	}

}
