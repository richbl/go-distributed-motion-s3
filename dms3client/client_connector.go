package dms3client

import (
	"fmt"
	"go-distributed-motion-s3/dms3libs"
	"net"
	"strconv"
	"time"
)

// Init configs the library and configuration for dms3client
func Init() {

	dms3libs.LoadLibConfig("/etc/distributed-motion-s3/dms3libs/dms3libs.toml")
	dms3libs.LoadComponentConfig(&clientConfig, "/etc/distributed-motion-s3/dms3client/dms3client.toml")

	dms3libs.SetLogFileLocation(clientConfig.Logging)
	dms3libs.CreateLogger(clientConfig.Logging)

	StartClient(clientConfig.Server.IP, clientConfig.Server.Port)

}

// StartClient periodically attempts to connect to the server (based on CheckInterval)
func StartClient(ServerIP string, ServerPort int) {

	for {
		if conn, err := net.Dial("tcp", ServerIP+":"+fmt.Sprint(ServerPort)); err != nil {
			dms3libs.LogInfo(err.Error())
		} else {
			dms3libs.LogInfo("OPEN connection from: " + conn.RemoteAddr().String())
			go processClientRequest(conn)
		}

		time.Sleep(time.Duration(clientConfig.Server.CheckInterval) * time.Second)
	}

}

// processClientRequest reads from the connection and processes motion detector state
func processClientRequest(conn net.Conn) {

	dms3libs.LogDebug(dms3libs.GetFunctionName())
	buf := make([]byte, 8)

	if n, err := conn.Read(buf); err != nil {
		dms3libs.LogInfo(err.Error())
	} else {
		val, _ := strconv.Atoi(string(buf[:n]))
		state := dms3libs.MotionDetectorState(val)

		if dms3libs.MotionDetector.SetState(state) {
			ProcessMotionDetectorState()
			dms3libs.LogInfo("Received motion detector state as: " + strconv.Itoa(int(state)))
		} else {
			dms3libs.LogInfo("Unanticipated motion detector state: ignored")
		}

	}

	dms3libs.LogInfo("CLOSE connection from: " + conn.RemoteAddr().String())
	conn.Close()

}
