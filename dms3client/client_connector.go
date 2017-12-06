package dms3client

import (
	"fmt"
	"go-distributed-motion-s3/dms3dashboard"
	"go-distributed-motion-s3/dms3libs"
	"net"
	"strconv"
	"time"
)

// Init configs the library, configuration, and dashboard for dms3client
func Init() {

	dms3libs.SetUptime(&startTime)

	dms3libs.LoadLibConfig("/etc/distributed-motion-s3/dms3libs/dms3libs.toml")
	LoadClientConfig("/etc/distributed-motion-s3/dms3client/dms3client.toml")

	cfg := ClientConfig.Logging
	dms3libs.CreateLogger(cfg.LogLevel, cfg.LogDevice, cfg.LogLocation, cfg.LogFilename)

<<<<<<< Updated upstream
	StartClient(ClientConfig.ServerIP, ClientConfig.ServerPort)
=======
	dms3dash.InitDashboardClient(configDashboardClientMetrics())
	StartClient(clientConfig.Server.IP, clientConfig.Server.Port)
>>>>>>> Stashed changes

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

		time.Sleep(time.Duration(ClientConfig.CheckInterval) * time.Second)
	}

}

// processClientRequest reads from the connection and processes dashboard and motion detector
// application state
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
<<<<<<< Updated upstream
		state, _ := strconv.Atoi(string(buf[:n]))
		dms3libs.MotionDetector.SetState(dms3libs.MotionDetectorState(state))
		ProcessMotionDetectorState(dms3libs.MotionDetector.State())
=======
		val, _ := strconv.Atoi(string(buf[:n]))
		state := dms3libs.MotionDetectorState(val)

		if dms3libs.MotionDetector.SetState(state) {
			ProcessMotionDetectorState()
			dms3libs.LogInfo("Received motion detector state as: " + strconv.Itoa(int(state)))
		} else {
			dms3libs.LogInfo("Received unanticipated motion detector state: ignored")
		}

>>>>>>> Stashed changes
	}

}
