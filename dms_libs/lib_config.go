package dmslibs

// SysCommand provides a location mapping of required system commands
var SysCommand = map[string]string{
	"APLAY":   "/usr/bin/aplay",
	"ARP":     "/usr/sbin/arp",
	"ARPSCAN": "/usr/bin/arp-scan", // TODO not currently used
	"CAT":     "/bin/cat",
	"GREP":    "/bin/grep",
	"PGREP":   "/usr/bin/pgrep",
	"PING":    "/bin/ping",
}

// MotionDetector is the motion detector application run on the clients
var MotionDetector = structMotionDetector{
	Location: "/usr/bin/motion",
	Command:  "motion",
	structMotionDetectorState: structMotionDetectorState{
		State: Stop,
	},
}

// possible states for the motion detector application
const (
	Start MotionDetectorState = iota
	Stop
)

type structMotionDetector struct {
	Location string
	Command  string
	structMotionDetectorState
}

type structMotionDetectorState struct {
	State MotionDetectorState
}

// MotionDetectorState is the type used to define the state (Start or Stop) of the motion detector application
type MotionDetectorState int
