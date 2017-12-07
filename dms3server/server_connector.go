package dms3server

import (
	"fmt"
	"go-distributed-motion-s3/dms3dashboard"
	"go-distributed-motion-s3/dms3libs"
	"net"
	"strconv"
)

// Init configs the library and configuration for dms3server
func Init() {

	dms3libs.SetUptime(&startTime)

	dms3libs.LoadLibConfig("/etc/distributed-motion-s3/dms3libs/dms3libs.toml")
	dms3libs.LoadComponentConfig(&serverConfig, "/etc/distributed-motion-s3/dms3server/dms3server.toml")

	dms3libs.SetLogFileLocation(serverConfig.Logging)
	dms3libs.CreateLogger(serverConfig.Logging)
	setMediaLocation(serverConfig)

	dms3dash.InitDashboardServer(configDashboardServerMetrics())
	startServer(serverConfig.Server.Port)

}

// configDashboardServerMetrics initializes the DeviceMetrics struct used by dms3dashboard
func configDashboardServerMetrics() *dms3dash.DeviceMetrics {

	dm := &dms3dash.DeviceMetrics{
		CheckInterval: serverConfig.Server.CheckInterval,
		StartTime:     startTime,
		Type:          dms3dash.Server,
	}

	return dm

}

// startServer starts the TCP server
func startServer(serverPort int) {

	if listener, error := net.Listen("tcp", ":"+fmt.Sprint(serverPort)); error != nil {
		dms3libs.LogFatal(error.Error())
	} else {
		defer listener.Close()
		serverLoop(listener)
	}

}

// serverLoop starts a loop to listen for clients, spawning a separate processing thread on
// dms3client connect
func serverLoop(listener net.Listener) {

	for {

		if conn, err := listener.Accept(); err != nil {
			dms3libs.LogFatal(err.Error())
		} else {
			dms3libs.LogInfo("OPEN connection from: " + conn.RemoteAddr().String())
			go processClient(conn)
		}

	}

}

// processClient passes motion detector application state to all dms3client listeners
func processClient(conn net.Conn) {

	dms3libs.LogDebug(dms3libs.GetFunctionName())

	dms3dash.SendDashboardRequest(conn)
	sendMotionDetectorState(conn)

	dms3libs.LogInfo("CLOSE connection from: " + conn.RemoteAddr().String())
	conn.Close()

}

// sendMotionDetectorState sends detector state to clients
func sendMotionDetectorState(conn net.Conn) {

	state := strconv.Itoa(int(DetermineMotionDetectorState()))

	if _, err := conn.Write([]byte(state)); err != nil {
		dms3libs.LogFatal(err.Error())
	} else {
		dms3libs.LogInfo("Sent motion detector state as: " + state)
	}

}
