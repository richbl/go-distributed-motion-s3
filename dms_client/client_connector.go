package client

import (
	"fmt"
	"go_server/dms_libs"
	"net"
	"strconv"
	"time"
)

// StartClient periodically attempts to connect to the server (based on CheckInterval)
func StartClient(ServerIP string, ServerPort int) {

	for {
		if conn, err := net.Dial("tcp", ServerIP+":"+fmt.Sprint(ServerPort)); err != nil {
			dmslibs.LogInfo(err.Error())
		} else {
			// server connection established
			defer conn.Close()
			dmslibs.LogInfo("OPEN connection from: " + conn.RemoteAddr().String())
			go processClientRequest(conn)
		}

		time.Sleep(CheckInterval * time.Second)
	}

}

// processClientRequest reads from the connection and processes motion detector state
func processClientRequest(conn net.Conn) {

	dmslibs.LogDebug(dmslibs.GetFunctionName())
	buf := make([]byte, 8)

	if n, err := conn.Read(buf); err != nil {
		dmslibs.LogInfo(err.Error())
	} else {
		state, _ := strconv.Atoi(string(buf[:n]))
		dmslibs.MotionDetector.SetState(dmslibs.MotionDetectorState(state))
		ProcessMotionDetectorState(dmslibs.MotionDetector.State())
	}

	dmslibs.LogInfo("CLOSE connection from: " + conn.RemoteAddr().String())
	conn.Close()

}
