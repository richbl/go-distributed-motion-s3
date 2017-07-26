package dms3server

import (
	"go-distributed-motion-s3/dms3libs"
	"time"
)

var checkIntervalTimestamp = dms3libs.GetCurTime()

// DetermineMotionDetectorState determines whether to start the motion detector application based
// device presence/time logic
//
func DetermineMotionDetectorState() dms3libs.MotionDetectorState {

	dms3libs.LogDebug(dms3libs.GetFunctionName())

	if checkIntervalExpired() {

		if timeInRange() || !deviceOnLAN() {
			return setMotionDetectorState(dms3libs.Start)
		}

		return setMotionDetectorState(dms3libs.Stop)
	}

	return dms3libs.MotionDetector.State()

}

// setMotionDetectorState sets the state read by device clients to starts/stop the motion detector
// applications
//
func setMotionDetectorState(state dms3libs.MotionDetectorState) dms3libs.MotionDetectorState {

	dms3libs.LogDebug(dms3libs.GetFunctionName())

	if dms3libs.MotionDetector.State() != state {
		dms3libs.MotionDetector.SetState(state)

		if PlayAudio == 1 {

			switch state {
			case dms3libs.Start:
				dms3libs.PlayAudio(AudioMotionDetectorStart)
			case dms3libs.Stop:
				dms3libs.PlayAudio(AudioMotionDetectorStop)
			}

		}
	}

	return state

}

// checkIntervalExpired determines if last check interval (in seconds) has expired
func checkIntervalExpired() bool {

	dms3libs.LogDebug(dms3libs.GetFunctionName())
	curTime := dms3libs.GetCurTime()

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

	dms3libs.LogDebug(dms3libs.GetFunctionName())

	if ScanForTime == 0 {
		return false
	}

	return calcDataRange()

}

// calcDataRange checks to see if the configured time range crosses into the next day, and
// determines time range accordingly
//
func calcDataRange() bool {

	dms3libs.LogDebug(dms3libs.GetFunctionName())
	const Start = 0
	const End = 1

	curTime := dms3libs.To24H(time.Now())

	if AlwaysOnRange[Start] > AlwaysOnRange[End] {
		return curTime >= AlwaysOnRange[Start] || curTime < AlwaysOnRange[End]
	}

	return curTime >= AlwaysOnRange[Start] && curTime < AlwaysOnRange[End]

}

// deviceOnLAN checks to see if device MACs exist on LAN (first freshens local arp cache to
// guarantee good results)
//
func deviceOnLAN() bool {

	dms3libs.LogDebug(dms3libs.GetFunctionName())
	dms3libs.PingHosts(IPBase, IPRange)
	return dms3libs.FindMacs(MacsToFind)

}
