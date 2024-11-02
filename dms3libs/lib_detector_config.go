// Package dms3libs provides motion detector application services for dms3 device components
package dms3libs

// MotionDetector represents the motion detector application running on the clients (e.g., motion)
var MotionDetector = motionDetectorStruct{
	state: Stop,
}

// states of the motion detector application
const (
	Stop MotionDetectorState = iota
	Start
)

// MotionDetectorStruct encapsulates the state of the motion detector
type motionDetectorStruct struct {
	state MotionDetectorState
}

// MotionDetectorState defines the motion detector application state type
type MotionDetectorState int

// State returns the motion detector application state
func (s motionDetectorStruct) State() MotionDetectorState {
	return s.state
}

// SetState sets the motion detector application state
func (s *motionDetectorStruct) SetState(state MotionDetectorState) bool {

	switch state {
	case Start, Stop:
		{
			s.state = state
			return true
		}
	default:
		return false
	}

}
