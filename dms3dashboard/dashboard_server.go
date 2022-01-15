// Package dms3dash server implements a dms3server-based metrics dashboard for all dms3clients
//
package dms3dash

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

// InitDashboardServer configs the library and server configuration for the dashboard
//
func InitDashboardServer(configPath string, dm *DeviceMetrics) {

	dms3libs.LogDebug(filepath.Base(dms3libs.GetFunctionName()))

	dashboardConfig = new(tomlTables)
	dms3libs.LoadComponentConfig(&dashboardConfig, filepath.Join(configPath, "dms3dashboard", "dms3dashboard.toml"))
	dms3libs.CheckFileLocation(configPath, "dms3dashboard", &dashboardConfig.Server.FileLocation, dashboardConfig.Server.Filename)

	dashboardData = new(deviceData)
	dm.appendServerMetrics()

	go dashboardConfig.Server.startDashboard(configPath)

}

// SendDashboardRequest manages dashboard requests and receipt of client device data
//
func SendDashboardRequest(conn net.Conn) {

	dms3libs.LogDebug(filepath.Base(dms3libs.GetFunctionName()))

	if DashboardEnable {
		sendDashboardEnableState(conn, "1")
		receiveDashboardData(conn)
	} else {
		sendDashboardEnableState(conn, "0")
	}

}

// startDashboard initializes and starts an HTTP server, serving the client dash on the server
//
func (dash *serverKeyValues) startDashboard(configPath string) {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	funcs := template.FuncMap{
		"ModVal":         dms3libs.ModVal,
		"FormatDateTime": dms3libs.FormatDateTime,
		"iconStatus":     iconStatus,
		"iconType":       iconType,
		"deviceType":     deviceType,
		"clientCount":    clientCount,
		"showEventCount": showEventCount,
	}

	tmpl := template.Must(template.New(dash.Filename).Funcs(funcs).ParseFiles(filepath.Join(dash.FileLocation, dash.Filename)))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(filepath.Join(configPath, "dms3dashboard", "assets")))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		dashboardData = &deviceData{
			Title:   dash.Title,
			Devices: dashboardData.Devices,
		}

		dashboardData.updateServerMetrics()

		if err := tmpl.Execute(w, dashboardData); err != nil {
			dms3libs.LogFatal(err.Error())
		}

	})

	if err := http.ListenAndServe(":"+fmt.Sprint(dash.Port), nil); err != nil {
		dms3libs.LogFatal(err.Error())
	}

}

// updateServerMetrics updates dynamic dashboard data of the server, triggered
// initially on dashboard start and subsequent webpage refreshes
//
func (dd *deviceData) updateServerMetrics() {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	for i := range dd.Devices {

		if dd.Devices[i].Platform.Type == Server {
			dd.Devices[i].Period.LastReport = time.Now()
			dd.Devices[i].Period.Uptime = dms3libs.Uptime(dd.Devices[i].Period.StartTime)
			dd.Devices[i].Platform.Kernel = dms3libs.DeviceKernel()
		} else {
			// check for and remove dead (non-reporting) client devices
			lastUpdate := dms3libs.SecondsSince(dd.Devices[i].Period.LastReport)
			missingDeviceLimit := uint32((dd.Devices[i].Period.CheckInterval * dashboardConfig.Server.DeviceStatus.Missing))

			if lastUpdate > missingDeviceLimit {
				dms3libs.LogInfo("Non-reporting remote device timeout reached: removing " + dd.Devices[i].Platform.Hostname + " client")
				dd.Devices = append(dd.Devices[:i], dd.Devices[i+1:]...)
				break
			}

		}

	}

}

// sendDashboardEnableState asks clients to send client info based on dashboard state
//
func sendDashboardEnableState(conn net.Conn, enableState string) {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	if _, err := conn.Write([]byte(enableState)); err != nil {
		dms3libs.LogFatal(err.Error())
	} else {
		dms3libs.LogInfo("Sent dashboard enable state as: " + enableState)
	}

}

// receiveDashboardData receives and parses client dashboard metrics
//
func receiveDashboardData(conn net.Conn) {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	updatedDeviceMetrics := new(DeviceMetrics)
	buf := make([]byte, 1024)

	if n, err := conn.Read(buf); err != nil {
		dms3libs.LogFatal(err.Error())
	} else {
		decBuf := bytes.NewBuffer(buf[:n]) // gob decoding of client metrics

		if err := gob.NewDecoder(decBuf).Decode(updatedDeviceMetrics); err != nil {
			dms3libs.LogFatal(err.Error())
		}

		updatedDeviceMetrics.updateDeviceMetrics()
	}

}

// updateDeviceMetrics adds new devices to the dashboard list, or updates existing device
// metrics, where Hostname is the unique key
//
func (udm *DeviceMetrics) updateDeviceMetrics() {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	// scan for existing client device
	for i := range dashboardData.Devices {

		if dashboardData.Devices[i].Platform.Hostname == udm.Platform.Hostname {

			if dashboardData.Devices[i].Platform.Type == Client {
				dashboardData.Devices[i].EventCount = udm.EventCount
				dashboardData.Devices[i].Period.LastReport = udm.Period.LastReport
				dashboardData.Devices[i].Period.Uptime = udm.Period.Uptime
				dashboardData.Devices[i].Platform.Kernel = udm.Platform.Kernel
				return
			}

		}

	}

	// add new client device and (optionally) resort device order
	dashboardData.Devices = append(dashboardData.Devices, *udm)

	if dashboardConfig.Server.ReSort {
		resortDashboardDevices()
	}

}

// resortDashboardDevices re-sorts all dashboard devices alphabetically
//
func resortDashboardDevices() {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	sort.Slice(dashboardData.Devices, func(i, j int) bool {
		switch strings.Compare(dashboardData.Devices[i].Platform.Hostname, dashboardData.Devices[j].Platform.Hostname) {
		case -1:
			return true
		case 1:
			return false
		}
		return dashboardData.Devices[i].Platform.Hostname > dashboardData.Devices[j].Platform.Hostname
	})
}

// appendServerMetrics appends the server to the dashboard list
//
func (dm *DeviceMetrics) appendServerMetrics() {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	serverData := new(DeviceMetrics)
	*serverData = *dm
	serverData.Platform.Type = Server
	serverData.Platform.Hostname = dms3libs.DeviceHostname()
	serverData.Platform.Environment = dms3libs.DeviceOS() + " " + dms3libs.DevicePlatform()
	serverData.Platform.Kernel = dms3libs.DeviceKernel()

	dashboardData.Devices = append(dashboardData.Devices, *serverData)

}

// iconStatus is an HTML template function that returns the CSS string representing icon color,
// depending on the last time the client reported status to the server, relative to the client's
// CheckInterval
//
func iconStatus(index int) string {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	lastUpdate := dms3libs.SecondsSince(dashboardData.Devices[index].Period.LastReport)
	checkInterval := dashboardData.Devices[index].Period.CheckInterval

	warningLimit := uint32((checkInterval * dashboardConfig.Server.DeviceStatus.Caution))
	dangerLimit := uint32((checkInterval * dashboardConfig.Server.DeviceStatus.Danger))

	switch {
	case lastUpdate < warningLimit:
		return "icon-success"
	case (lastUpdate >= warningLimit) && (lastUpdate < dangerLimit):
		return "icon-warning"
	case lastUpdate >= dangerLimit:
		return "icon-danger"
	default:
		return ""
	}

}

// iconType is an HTML template function that returns an icon based on device type
//
func iconType(index int) string {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	switch dashboardData.Devices[index].Platform.Type {
	case Client:
		return "icon-raspberry-pi"
	case Server:
		return "icon-server2"
	default:
		return ""
	}

}

// deviceType is an HTML template function that returns a string based on device type
//
func deviceType(index int) string {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	switch dashboardData.Devices[index].Platform.Type {
	case Client:
		return "client"
	case Server:
		return "server"
	default:
		return ""
	}

}

// clientCount is an HTML template function that returns the current count of dms3clients
// reporting to the server
//
func clientCount() int {
	return len(dashboardData.Devices) - 1
}

// showEventCount is an HTML template function that returns whether to display client event count
//
func showEventCount(index int) bool {
	return dashboardData.Devices[index].ShowEventCount
}
