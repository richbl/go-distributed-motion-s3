package server

import (
	"go_server/dms_libs"
	"time"
)

var checkIntervalTimestamp = dmslibs.GetCurTime()

// DetermineMotionDetectorState determines whether to start the motion detector application based
// device presence/time logic
//
func DetermineMotionDetectorState() dmslibs.MotionDetectorState {

	dmslibs.LogDebug(dmslibs.GetFunctionName())

	if checkIntervalExpired() {

		if timeInRange() || !deviceOnLAN() {
			return setMotionDetectorState(dmslibs.Start)
		}

		return setMotionDetectorState(dmslibs.Stop)
	}

	return dmslibs.MotionDetector.State()

}

// setMotionDetectorState sets the state read by device clients to starts/stop the motion detector
// applications
//
func setMotionDetectorState(state dmslibs.MotionDetectorState) dmslibs.MotionDetectorState {

	dmslibs.LogDebug(dmslibs.GetFunctionName())

	if dmslibs.MotionDetector.State() != state {
		dmslibs.MotionDetector.SetState(state)

		if PlayAudio == 1 {

			switch state {
			case dmslibs.Start:
				dmslibs.PlayAudio(AudioMotionDetectorStart)
			case dmslibs.Stop:
				dmslibs.PlayAudio(AudioMotionDetectorStop)
			}

		}
	}

	return state

}

// checkIntervalExpired determines if last check interval (in seconds) has expired
func checkIntervalExpired() bool {

	dmslibs.LogDebug(dmslibs.GetFunctionName())
	curTime := dmslibs.GetCurTime()

	if (curTime - checkIntervalTimestamp) >= CheckInterval {
		checkIntervalTimestamp = curTime
		return true
	}

	return false

}

// timeInRange checks to see if the current time is within the bounds of the 'always on' range
// (if that ScanForTime option is enabled)
//
func timeInRange() bool {

	dmslibs.LogDebug(dmslibs.GetFunctionName())

	if ScanForTime == 0 {
		return false
	}

	return calcDataRange()

}

// calcDataRange checks to see if the configured time range crosses into the next day, and
// determines time range accordingly
//
func calcDataRange() bool {

	dmslibs.LogDebug(dmslibs.GetFunctionName())
	const Start = 0
	const End = 1

	curTime := dmslibs.To24H(time.Now())

	if AlwaysOnRange[Start] > AlwaysOnRange[End] {
		return curTime >= AlwaysOnRange[Start] || curTime < AlwaysOnRange[End]
	}

	return curTime >= AlwaysOnRange[Start] && curTime < AlwaysOnRange[End]

}

// deviceOnLAN checks to see if device MACs exist on LAN (first freshens local arp cache to
// guarantee good results)
//
func deviceOnLAN() bool {

	dmslibs.LogDebug(dmslibs.GetFunctionName())
	dmslibs.PingHosts(IPBase, IPRange)
	return dmslibs.FindMacs(MacsToFind)

}
