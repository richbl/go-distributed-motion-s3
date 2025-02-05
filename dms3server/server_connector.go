// Package dms3server connector initializes the dms3server device component
package dms3server

import (
	"fmt"
	"net"
	"path"
	"path/filepath"
	"strconv"

	dms3dash "github.com/richbl/go-distributed-motion-s3/dms3dashboard"
	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

// Init configs the library and configuration for dms3server
func Init(configPath string) {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	dms3libs.LoadLibConfig(filepath.Join(configPath, dms3libs.DMS3Libs, "dms3libs.toml"))
	dms3libs.LoadComponentConfig(&ServerConfig, filepath.Join(configPath, dms3libs.DMS3Server, "dms3server.toml"))

	dms3libs.SetLogFileLocation(ServerConfig.Logging)
	dms3libs.CreateLogger(ServerConfig.Logging)

	dms3libs.LogInfo("dms3server " + dms3libs.GetProjectVersion() + " started")

	setMediaLocation(configPath, ServerConfig)
	dms3dash.DashboardEnable = ServerConfig.Server.EnableDashboard

	if dms3dash.DashboardEnable {
		dms3dash.InitDashboardServer(configPath, ServerConfig.Server.CheckInterval)
	}

	startServer(ServerConfig.Server.Port)

}

// startServer starts the TCP server
func startServer(serverPort int) {

	if listener, err := net.Listen("tcp", ":"+fmt.Sprint(serverPort)); err != nil {
		dms3libs.LogFatal(err.Error())
	} else {
		dms3libs.LogInfo("TCP server started")
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

	dms3libs.LogDebug(filepath.Base(dms3libs.GetFunctionName()))

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

// setMediaLocation sets the location where audio files are located for motion detection
// application start/stop events
func setMediaLocation(configPath string, config *structSettings) {

	media := []mediaPath{
		{&config.Audio.PlayMotionStart, filepath.Join(dms3libs.DMS3Server, "media", "motion_start.wav")},
		{&config.Audio.PlayMotionStop, filepath.Join(dms3libs.DMS3Server, "media", "motion_stop.wav")},
	}

	for _, mp := range media {
		checkMediaLocations(configPath, mp)
	}
}

// checkMediaLocations identifies the possible location of media (audio) files (release or
// development depending on the runtime environment)
func checkMediaLocations(configPath string, mp mediaPath) {

	relPath := filepath.Join(configPath, mp.mediaLocation)
	devPath := filepath.Join(path.Dir(dms3libs.GetPackageDir()), mp.mediaLocation)

	if mp.configLocation != nil && *mp.configLocation != "" {
		return // already set
	}

	var possiblePaths = []string{relPath, devPath}
	for _, path := range possiblePaths {
		if dms3libs.IsFile(path) {
			*mp.configLocation = path
			return
		}
	}

	dms3libs.LogFatal("unable to set media location... check TOML configuration file")
}
