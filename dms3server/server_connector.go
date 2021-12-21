// Package dms3server connector initializes the dms3server device component
//
package dms3server

import (
	"fmt"
	"net"
	"path/filepath"
	"strconv"

	dms3dash "github.com/richbl/go-distributed-motion-s3/dms3dashboard"
	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

// Init configs the library and configuration for dms3server
func Init(configPath string) {

	dms3libs.SetUptime(&startTime)

	dms3libs.LoadLibConfig(filepath.Join(configPath, "dms3libs/dms3libs.toml"))
	dms3libs.LoadComponentConfig(&ServerConfig, filepath.Join(configPath, "dms3server/dms3server.toml"))

	dms3libs.SetLogFileLocation(ServerConfig.Logging)
	dms3libs.CreateLogger(ServerConfig.Logging)

	setMediaLocation(configPath, ServerConfig)

	dms3dash.InitDashboardServer(configPath, configDashboardServerMetrics())
	startServer(ServerConfig.Server.Port)

}

// configDashboardServerMetrics initializes the DeviceMetrics struct used by dms3dashboard
func configDashboardServerMetrics() *dms3dash.DeviceMetrics {

	dm := &dms3dash.DeviceMetrics{
		Platform: dms3dash.DevicePlatform{
			Type: dms3dash.Client,
		},
		Period: dms3dash.DeviceTime{
			StartTime:     startTime,
			CheckInterval: ServerConfig.Server.CheckInterval,
		},
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
//
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
