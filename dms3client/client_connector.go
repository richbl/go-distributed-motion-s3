// Package dms3client connector initializes the dms3client device component
package dms3client

import (
	"fmt"
	"net"
	"path/filepath"
	"strconv"
	"time"

	dms3dash "github.com/richbl/go-distributed-motion-s3/dms3dashboard"
	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

// Init configs the library, configuration, and dashboard for dms3client
func Init(configPath string) {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	dms3libs.LoadLibConfig(filepath.Join(configPath, dms3libs.DMS3Libs, "dms3libs.toml"))
	dms3libs.LoadComponentConfig(&clientConfig, filepath.Join(configPath, dms3libs.DMS3Client, dms3libs.DMS3TOML))

	dms3libs.SetLogFileLocation(clientConfig.Logging)
	dms3libs.CreateLogger(clientConfig.Logging)

	dms3libs.LogInfo("dms3client " + dms3libs.GetProjectVersion() + " started")

	dms3dash.InitDashboardClient(configPath, clientConfig.Server.CheckInterval)
	startClient(clientConfig.Server.IP, clientConfig.Server.Port)

}

// startClient periodically attempts to connect to the server (based on CheckInterval)
func startClient(serverIP string, serverPort int) {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	for {
		if conn, err := net.Dial("tcp", serverIP+":"+fmt.Sprint(serverPort)); err != nil {
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
func processClientRequest(conn net.Conn) {

	dms3libs.LogDebug(filepath.Base(dms3libs.GetFunctionName()))

	dms3dash.ReceiveDashboardRequest(conn)
	receiveMotionDetectorState(conn)

	dms3libs.LogInfo("CLOSE connection from: " + conn.RemoteAddr().String())
	conn.Close()

}

// receiveMotionDetectorState receives motion detector state from the server
func receiveMotionDetectorState(conn net.Conn) {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	buf := make([]byte, 8)

	// receive motion detector application state

	var n int
	var err error

	if n, err = conn.Read(buf); err != nil {
		dms3libs.LogInfo(err.Error())

		return
	}

	val, err := strconv.Atoi(string(buf[:n]))
	if err != nil {
		dms3libs.LogInfo("Error converting motion detector state to integer: " + err.Error())
	}

	state := dms3libs.MotionDetectorState(val)

	if dms3libs.MotionDetector.SetState(state) {
		ProcessMotionDetectorState()
		dms3libs.LogInfo("Received motion detector state as: " + strconv.Itoa(int(state)))

		return
	}

	dms3libs.LogInfo("Received unanticipated motion detector state: ignored")

}
