// Package dms3dash client implements client services for a dms3server-based metrics dashboard
//
package dms3dash

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	"path/filepath"
	"time"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

var dashboardClientMetrics *DeviceMetrics

// InitDashboardClient loads configuration and assigns the dashboard client profile (sets
//  static client metrics)
//
func InitDashboardClient(configPath string, dm *DeviceMetrics) {

	dashboardConfig = new(tomlTables)
	dms3libs.LoadComponentConfig(&dashboardConfig, filepath.Join(configPath, "dms3dashboard", "dms3dashboard.toml"))

	dashboardClientMetrics = &DeviceMetrics{
		Platform: DevicePlatform{
			Type:        dm.Platform.Type,
			Hostname:    dms3libs.DeviceHostname(),
			Environment: dms3libs.DeviceOS() + " " + dms3libs.DevicePlatform(),
			Kernel:      dms3libs.DeviceKernel(),
		},
		Period: DeviceTime{
			StartTime:     dm.Period.StartTime,
			CheckInterval: dm.Period.CheckInterval,
		},
	}

	dashboardClientMetrics.checkImagesFolder()

}

// ReceiveDashboardRequest receives server requests and returns data
func ReceiveDashboardRequest(conn net.Conn) {

	if receiveDashboardEnableState(conn) {
		sendDashboardData(conn)
	}

}

// checkImagesFolder confirms the location of the motion-triggered image/movie files managed by
// the motion detector application (if installed), and used in displaying client metrics in the
// dashboard
//
func (dash *DeviceMetrics) checkImagesFolder() {

	if dms3libs.IsFile(dashboardConfig.Client.ImagesFolder) {
		dashboardClientMetrics.ShowEventCount = true
	} else {

		if dashboardConfig.Client.ImagesFolder == "" {
			dashboardClientMetrics.ShowEventCount = false
		} else {
			log.Fatalln("unable to find motion detector application images folder... check TOML configuration file")
		}

	}

}

// receiveDashboardEnableState parses the server dashboard state notification, returning true
// if the dashboard state is enabled
//
func receiveDashboardEnableState(conn net.Conn) bool {

	buf := make([]byte, 16)

	if n, err := conn.Read(buf); err != nil {
		dms3libs.LogFatal(err.Error())
		return false
	} else {
		val := string(buf[:n])
		dms3libs.LogInfo("Received dashboard enable state as: " + val)
		return (val == "1")
	}

}

// sendDashboardData sends dashboard info to server
func sendDashboardData(conn net.Conn) {

	// update client metrics
	dashboardClientMetrics.Period.LastReport = time.Now()
	dashboardClientMetrics.Period.Uptime = dms3libs.Uptime(dashboardClientMetrics.Period.StartTime)

	if dashboardClientMetrics.ShowEventCount {
		dashboardClientMetrics.EventCount = dms3libs.CountFilesInDir(dashboardConfig.Client.ImagesFolder)
	}

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
