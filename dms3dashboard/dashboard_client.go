// Package dms3dash client implements a dms3server-based metrics dashboard for all dms3clients
//
package dms3dash

import (
	"bytes"
	"encoding/gob"
	"go-distributed-motion-s3/dms3libs"
	"net"
	"path/filepath"
	"time"
)

var dashboardClientMetrics *DeviceMetrics

// InitDashboardClient loads configuration and assigns the dashboard client profile (sets
//	static client metrics)
//
func InitDashboardClient(configPath string, dm *DeviceMetrics) {

	dashboardConfig = new(tomlTables)
	dms3libs.LoadComponentConfig(&dashboardConfig, filepath.Join(configPath, "dms3dashboard/dms3dashboard.toml"))

	dashboardClientMetrics = &DeviceMetrics{
		Hostname:      dms3libs.DeviceHostname(),
		Environment:   dms3libs.DeviceOS() + " " + dms3libs.DevicePlatform(),
		Kernel:        dms3libs.DeviceKernel(),
		CheckInterval: dm.CheckInterval,
		StartTime:     dm.StartTime,
	}

}

// ReceiveDashboardRequest receives server requests and returns data
func ReceiveDashboardRequest(conn net.Conn) {

	if receiveDashboardEnableState(conn) == true {
		sendDashboardData(conn)
	}

}

// receiveDashboardEnableState parses the server dashboard state notification, returning true
// if the dashboard state is enabled
//
func receiveDashboardEnableState(conn net.Conn) bool {

	buf := make([]byte, 16)
	n, err := conn.Read(buf)

	if err != nil {
		dms3libs.LogFatal(err.Error())
	}

	val := string(buf[:n])
	dms3libs.LogInfo("Received dashboard enable state as: " + val)
	return (val == "1")

}

// sendDashboardData sends dashboard info to server
func sendDashboardData(conn net.Conn) {

	// update client metrics
	dashboardClientMetrics.LastReport = time.Now()
	dashboardClientMetrics.Uptime = dms3libs.Uptime(dashboardClientMetrics.StartTime)
	dashboardClientMetrics.EventCount = dms3libs.CountFilesInDir(dashboardConfig.Client.ImagesFolder)

	// gob encoding of client metrics
	encBuf := new(bytes.Buffer)

	if err := gob.NewEncoder(encBuf).Encode(dashboardClientMetrics); err != nil {
		dms3libs.LogFatal(err.Error())
	}

	if _, err := conn.Write(encBuf.Bytes()); err != nil {
		dms3libs.LogFatal(err.Error())
	} else {
		dms3libs.LogInfo("Sent client dashboard data")
	}

}
