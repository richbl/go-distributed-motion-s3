// Package dms3client connector initializes the dms3client device component
//
package dms3client

import (
	"fmt"
	"go-distributed-motion-s3/dms3dashboard"
	"go-distributed-motion-s3/dms3libs"
	"net"
	"path/filepath"
	"strconv"
	"time"
)

// Init configs the library, configuration, and dashboard for dms3client
func Init(configPath string) {

	dms3libs.SetUptime(&startTime)

	dms3libs.LoadLibConfig(filepath.Join(configPath, "dms3libs/dms3libs.toml"))
	dms3libs.LoadComponentConfig(&clientConfig, filepath.Join(configPath, "dms3client/dms3client.toml"))

	dms3libs.SetLogFileLocation(clientConfig.Logging)
	dms3libs.CreateLogger(clientConfig.Logging)

	dms3dash.InitDashboardClient(configPath, configDashboardClientMetrics())
	startClient(clientConfig.Server.IP, clientConfig.Server.Port)

}

// configDashboardClientMetrics initializes the DeviceMetrics struct used by dms3dashboard
func configDashboardClientMetrics() *dms3dash.DeviceMetrics {

	dm := &dms3dash.DeviceMetrics{
		CheckInterval: clientConfig.Server.CheckInterval,
		StartTime:     startTime,
		Type:          dms3dash.Client,
	}

	return dm

}

// startClient periodically attempts to connect to the server (based on CheckInterval)
func startClient(ServerIP string, ServerPort int) {

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

// processClientRequest reads from the connection and processes dashboard and motion detector
// application state
//
func processClientRequest(conn net.Conn) {

	dms3libs.LogDebug(dms3libs.GetFunctionName())

	dms3dash.ReceiveDashboardRequest(conn)
	receiveMotionDetectorState(conn)

	dms3libs.LogInfo("CLOSE connection from: " + conn.RemoteAddr().String())
	conn.Close()

}

// receiveMotionDetectorState receives motion detector state from the server
func receiveMotionDetectorState(conn net.Conn) {

	buf := make([]byte, 8)

	// receive motion detector application state

	if n, err := conn.Read(buf); err != nil {
		dms3libs.LogInfo(err.Error())
	} else {
		val, _ := strconv.Atoi(string(buf[:n]))
		state := dms3libs.MotionDetectorState(val)

		if dms3libs.MotionDetector.SetState(state) {
			ProcessMotionDetectorState()
			dms3libs.LogInfo("Received motion detector state as: " + strconv.Itoa(int(state)))
		} else {
			dms3libs.LogInfo("Received unanticipated motion detector state: ignored")
		}

	}

}
