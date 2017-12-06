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

		if ServerConfig.PlayAudio == 1 {

<<<<<<< Updated upstream
			switch state {
			case dms3libs.Start:
				dms3libs.PlayAudio(ServerConfig.AudioMotionDetectorStart)
			case dms3libs.Stop:
				dms3libs.PlayAudio(ServerConfig.AudioMotionDetectorStop)
			}

=======
	if serverConfig.Audio.Enable {

		switch state {
		case dms3libs.Start:
			dms3libs.PlayAudio(serverConfig.Audio.PlayMotionStart)
		case dms3libs.Stop:
			dms3libs.PlayAudio(serverConfig.Audio.PlayMotionStop)
>>>>>>> Stashed changes
		}
	}

	return state

}

// checkIntervalExpired determines if last check interval (in seconds) has expired
func checkIntervalExpired() bool {

	dms3libs.LogDebug(dms3libs.GetFunctionName())
	curTime := dms3libs.GetCurTime()

<<<<<<< Updated upstream
	if (curTime - checkIntervalTimestamp) >= ServerConfig.CheckInterval {
		checkIntervalTimestamp = curTime
=======
	if time.Since(checkIntervalTimestamp).Seconds() >= float64(serverConfig.Server.CheckInterval) {
		checkIntervalTimestamp = time.Now()
>>>>>>> Stashed changes
		return true
	}

	return false

}

// timeInRange checks to see if the current time is within the bounds of the 'always on' range
// (if the ScanForTime option is enabled)
//
func timeInRange() bool {

	dms3libs.LogDebug(dms3libs.GetFunctionName())

<<<<<<< Updated upstream
	if ServerConfig.ScanForTime {
		return calcDataRange()
=======
	if serverConfig.AlwaysOn.Enable {
		return calcTimeRange()
>>>>>>> Stashed changes
	}

	return false

}

// calcDataRange checks to see if the configured time range crosses into the next day, and
// determines time range accordingly
//
func calcDataRange() bool {

	dms3libs.LogDebug(dms3libs.GetFunctionName())

	curTime := dms3libs.To24H(time.Now())

<<<<<<< Updated upstream
	startTime := dms3libs.Format24H(ServerConfig.AlwaysOnRange[0])
	endTime := dms3libs.Format24H(ServerConfig.AlwaysOnRange[1])
=======
	startTime := dms3libs.Format24H(serverConfig.AlwaysOn.TimeRange[0])
	endTime := dms3libs.Format24H(serverConfig.AlwaysOn.TimeRange[1])
>>>>>>> Stashed changes

	if startTime > endTime {
		return curTime >= startTime || curTime < endTime
	}

	return curTime >= startTime && curTime < endTime

}

// deviceOnLAN checks to see if device MACs exist on LAN (first freshens local arp cache to
// guarantee good results)
//
func deviceOnLAN() bool {

	dms3libs.LogDebug(dms3libs.GetFunctionName())
<<<<<<< Updated upstream
	dms3libs.PingHosts(ServerConfig.IPBase, ServerConfig.IPRange)
	return dms3libs.FindMacs(ServerConfig.MacsToFind)
=======
	dms3libs.PingHosts(serverConfig.UserProxy.IPBase, serverConfig.UserProxy.IPRange)
	return dms3libs.FindMacs(serverConfig.UserProxy.MacsToFind)
>>>>>>> Stashed changes

}
