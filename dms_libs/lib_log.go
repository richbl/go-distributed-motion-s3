package dmslibs

import (
	"log"
	"os"
)

var (
	// Info is the global logger for informational alerts
	Info *log.Logger

	// Error is the global logger for error alerts
	Error *log.Logger
)

func init() {
	var f *os.File
	Info = log.New(f, "INFO: ", log.Lshortfile|log.LstdFlags)
	Error = log.New(f, "ERROR: ", log.Lshortfile|log.LstdFlags)
}

// CreateLogger creates an application log file (1) or redirects to STDOUT (2) based on loggerType
func CreateLogger(loggerType int, logLocation string, logFilename string) {

	var (
		f   *os.File
		err error
	)

	switch loggerType {
	case 1:
		{
			// TODO rotating log with size cap?
			f, err = os.OpenFile(logFilename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				log.Fatalln(err)
			}
		}
	case 2:
		{
			f = os.Stdout
		}
	default:
		// no log
	}

	Info.SetOutput(f)
	Error.SetOutput(f)
}
