// Package dms3dash client implements client services for a dms3server-based metrics dashboard
package dms3dash

import (
	"bytes"
	"encoding/gob"
	"net"
	"path/filepath"
	"time"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

var dashboardClientMetrics *DeviceMetrics

// InitDashboardClient loads configuration and assigns the dashboard client profile (sets static client metrics)
func InitDashboardClient(configPath string, checkInterval uint16) {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	dashboardConfig = new(tomlTables)
	dms3libs.LoadComponentConfig(&dashboardConfig, filepath.Join(configPath, dms3libs.DMS3Dashboard, "dms3dashboard.toml"))

	dashboardClientMetrics = &DeviceMetrics{
		Platform: DevicePlatform{
			Type:        Client,
			Hostname:    dms3libs.GetDeviceHostname(),
			OSName:      dms3libs.GetDeviceOSName(),
			Environment: dms3libs.GetDeviceDetails(dms3libs.Sysname) + " " + dms3libs.GetDeviceDetails(dms3libs.Machine),
			Kernel:      dms3libs.GetDeviceDetails(dms3libs.Release),
		},
		Period: DeviceTime{
			CheckInterval: checkInterval,
			StartTime:     time.Now(),
			Uptime:        "",
			LastReport:    time.Now(),
		},
		ShowEventCount: false,
		EventCount:     0,
	}

	dashboardClientMetrics.checkImagesFolder()
}

// ReceiveDashboardRequest receives server requests and returns data
func ReceiveDashboardRequest(conn net.Conn) {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	if receiveDashboardEnableState(conn) {
		sendDashboardData(conn)
	}

}

// receiveDashboardEnableState parses the server dashboard state notification, returning true
// if the dashboard state is enabled
func receiveDashboardEnableState(conn net.Conn) bool {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	buf := make([]byte, 16)
	var n int
	var err error

	if n, err = conn.Read(buf); err != nil {
		dms3libs.LogFatal(err.Error())
		return false
	}

	val := string(buf[:n])
	dms3libs.LogInfo("Received dashboard enable state as: " + val)

	return (val == "1")
}

// sendDashboardData sends dashboard info to server
func sendDashboardData(conn net.Conn) {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	// update client metrics
	dashboardClientMetrics.Period.LastReport = time.Now()
	dashboardClientMetrics.Period.Uptime = dms3libs.Uptime(dashboardClientMetrics.Period.StartTime)

	// calls not needed,as these values do not change (unless client reboots)
	//
	// dashboardClientMetrics.Platform.OSName = dms3libs.GetDeviceOSName()
	// dashboardClientMetrics.Platform.Environment = dms3libs.GetDeviceDetails(dms3libs.Sysname) + " " + dms3libs.GetDeviceDetails(dms3libs.Machine)
	// dashboardClientMetrics.Platform.Kernel = dms3libs.GetDeviceDetails(dms3libs.Release)

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

// checkImagesFolder confirms the location of the motion-triggered image/movie files managed by
// the motion detector application (if installed), and used in displaying client metrics in the
// dashboard
func (dash *DeviceMetrics) checkImagesFolder() {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	if dms3libs.IsFile(dashboardConfig.Client.ImagesFolder) {
		dashboardClientMetrics.ShowEventCount = true
	} else {

		if dashboardConfig.Client.ImagesFolder == "" {
			dashboardClientMetrics.ShowEventCount = false
		} else {
			dms3libs.LogFatal("Unable to find motion detector application images folder... check TOML configuration file")
		}

	}

}
