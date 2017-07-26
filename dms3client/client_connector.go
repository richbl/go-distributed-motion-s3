package dms3client

import (
	"fmt"
	"go-distributed-motion-s3/dms3libs"
	"net"
	"strconv"
	"time"
)

// StartClient periodically attempts to connect to the server (based on CheckInterval)
func StartClient(ServerIP string, ServerPort int) {

	for {
		if conn, err := net.Dial("tcp", ServerIP+":"+fmt.Sprint(ServerPort)); err != nil {
			dms3libs.LogInfo(err.Error())
		} else {
			// server connection established
			defer conn.Close()
			dms3libs.LogInfo("OPEN connection from: " + conn.RemoteAddr().String())
			go processClientRequest(conn)
		}

		time.Sleep(CheckInterval * time.Second)
	}

}

// processClientRequest reads from the connection and processes motion detector state
func processClientRequest(conn net.Conn) {

	dms3libs.LogDebug(dms3libs.GetFunctionName())
	buf := make([]byte, 8)

	if n, err := conn.Read(buf); err != nil {
		dms3libs.LogInfo(err.Error())
	} else {
		state, _ := strconv.Atoi(string(buf[:n]))
		dms3libs.MotionDetector.SetState(dms3libs.MotionDetectorState(state))
		ProcessMotionDetectorState(dms3libs.MotionDetector.State())
	}

	dms3libs.LogInfo("CLOSE connection from: " + conn.RemoteAddr().String())
	conn.Close()

}
