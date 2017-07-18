package dmslibs

import (
	"log"
	"os"
)

var (
	// Fatal is the global logger for fatal error alerts
	Fatal *log.Logger

	// Info is the global logger for informational alerts
	Info *log.Logger

	// Debug is the global logger for debugger alerts
	Debug *log.Logger

	loggingLevel int
)

func init() {
	var f *os.File
	Fatal = log.New(f, "FATAL: ", log.Lshortfile|log.LstdFlags)
	Info = log.New(f, "INFO: ", log.Lshortfile|log.LstdFlags)
	Debug = log.New(f, "DEBUG: ", log.Lshortfile|log.LstdFlags)
}

// CreateLogger creates an application log file (1) or redirects to STDOUT (2) based on logDevice
func CreateLogger(logLevel int, logDevice int, logLocation string, logFilename string) {

	var (
		f   *os.File
		err error
	)

	if logLevel > 0 {
		loggingLevel = logLevel

		switch logDevice {
		case 0:
			f = os.Stdout
		case 1:
			{
				f, err = os.OpenFile(logFilename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
				if err != nil {
					log.Fatalln(err)
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

	if loggingLevel == 4 {
		Debug.Println(msg)
	}

}
