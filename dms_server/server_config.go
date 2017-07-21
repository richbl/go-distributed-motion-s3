package server

import (
	"go_server/dms_libs"
)

// LogLevel sets the log levels for application logging using the following table:
//
//		0 - OFF, no logging
//		1 - FATAL, report fatal events
//		2 - INFO, report informational events
//    4 - DEBUG, report debugging events
//
const LogLevel = 2

// LogDevice determines to what device logging should be set using the following table:
//
//		0 - STDOUT (terminal)
//		1 - log file
//
// Ignored if LogLevel == 0
//
const LogDevice = 0

// LogFilename is the logging filename
// Ignored if LogLevel == 0 or LogDevice == 0
//
const LogFilename = "dms_server.log"

// LogLocation is the location of logfile (full path)
// By default, this is in the local folder (e.g., /etc/distributed_motion_surveillance/dms_server)
// Ignored if LogLevel == 0 or LogDevice == 0
//
var LogLocation = dmslibs.GetPackageDir()

// PlayAudio enables (1) or disables (0) the play-back of audio on motion detector application
// start/stop
//
const PlayAudio = 1

// ServerPort is the port on which to run the motion server
const ServerPort = 1965

// AudioMotionDetectorStart is the audio file played when the motion detector application is
// activated... by default, this is in the local folder (e.g., /etc/motion_surveillance/motion_monitor)
// Ignored if PlayAudio == 0
//
var AudioMotionDetectorStart = dmslibs.GetPackageDir() + "/media/motion_start.wav"

// AudioMotionDetectorStop is the audio file played when the motion detector application is
// deactivated... by default, this is in the local folder (e.g., /etc/motion_surveillance/motion_monitor)
// Ignored if PlayAudio == 0
//
var AudioMotionDetectorStop = dmslibs.GetPackageDir() + "/media/motion_stop.wav"

// ScanForTime enables (1) or disables (0) motion detector application based on time-of-day
const ScanForTime = 0

// CheckInterval is the interval (in seconds) between checks for change to motion_state
const CheckInterval = 15

// AlwaysOnRange is the start and end times (24-hour format) for motion to always be enabled,
// regardless of current IoT device existence on the LAN
// Ignored if ScanForTime == 0
//
var AlwaysOnRange = []string{"2300", "0400"}

// IPBase is the first three address octets defining the LAN (e.g., 10.10.10.) where devices will
// be scanned for to determine when motion should be run
//
const IPBase = "10.10.10."

// IPRange is the fourth address octet defined as a range (e.g., 100..254)
var IPRange = []int{100, 254}

// MacsToFind are the MAC addresses of IoT device(s) to search for on the LAN
var MacsToFind = []string{"24:da:9b:0d:53:8f", "f8:cf:c5:d2:bb:9e"}
