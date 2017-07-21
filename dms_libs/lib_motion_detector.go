package dmslibs

// MotionDetector is the motion detector application run on the clients
var MotionDetector = structMotionDetector{
	Location: "/usr/bin/motion",
	Command:  "motion",
	State:    Stop,
}

// possible states for the motion detector application
const (
	Start MotionDetectorState = iota
	Stop
)

type structMotionDetector struct {
	Location string
	Command  string
	State    MotionDetectorState
}

// MotionDetectorState is the type used to define the state (Start or Stop) of the motion
// detector application
//
type MotionDetectorState int
