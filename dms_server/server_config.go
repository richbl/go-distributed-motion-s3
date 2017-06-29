package server

import (
	"go_server/libs"
)

// Logging enable (1) or disable (0) application logging
// NOTE: passing in 2 sets logging to STDOUT (useful when debugging or running as daemon)
const Logging = 2

// LogFilename is the logging filename
// Ignored if LOGGING == 0
const LogFilename = "dms_server.log"

// LogLocation is the location of logfile (full path)
// By default, this is in the local folder (e.g., /etc/distributed_motion_surveillance/dms_server)
// Ignored if LOGGING == 0
var LogLocation = libconfig.GetPath()

// PlayAudio enables (1) or disables (0) the play-back of audio on motion daemon start/stop
const PlayAudio = 1

// ServerPort is the port on which to run the motion server
const ServerPort = 1965

// AudioMotionStart is the audio file played when the motion daemon is activated
// By default, this is in the local folder (e.g., /etc/motion_surveillance/motion_monitor)
// Ignored if PLAY_AUDIO == 0
var AudioMotionStart = libconfig.GetPath() + "/media/motion_start.wav"

// AudioMotionStop is the audio file played when the motion daemon is deactivated
// By default, this is in the local folder (e.g., /etc/motion_surveillance/motion_monitor)
// Ignored if PLAY_AUDIO == 0
var AudioMotionStop = libconfig.GetPath() + "/media/motion_stop.wav"

// ScanForTime enables (1) or disables (0) motion daemon based on time-of-day
const ScanForTime = 0

// CheckInterval is the interval (in seconds) between checks for change to motion_state
const CheckInterval = 15

// AlwaysOnRange is the start and end times (24-hour format) for motion to always be enabled,
// regardless of current IoT device existence on the LAN
// Ignored if SCAN_FOR_TIME == 0
var AlwaysOnRange = [...]string{"2300", "0400"}

// IPBase is the first three address octets defining the LAN (e.g., 10.10.10.) where devices will be
// scanned for to determine when motion should be run
const IPBase = "10.10.10."

// IPRange is the fourth address octet defined as a range (e.g., 100..254)
var IPRange = [...]int{100, 254}

// MacsToFind are the MAC addresses of IoT device(s) to search for on the LAN
var MacsToFind = [...]string{"24:da:9b:0d:53:8f", "f8:cf:c5:d2:bb:9e"}
