// Package dms3dash server implements a dms3server-based metrics dashboard for all dms3clients
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
func InitDashboardServer(configPath string, checkInterval uint16) {

	dms3libs.LogDebug(filepath.Base(dms3libs.GetFunctionName()))

	dashboardConfig = new(tomlTables)
	dms3libs.LoadComponentConfig(&dashboardConfig, filepath.Join(configPath, dms3libs.DMS3Dashboard, dms3libs.DMS3dashboardTOML))
	dms3libs.CheckFileLocation(configPath, dms3libs.DMS3Dashboard, &dashboardConfig.Server.FileLocation, dashboardConfig.Server.Filename)

	dashboardData = &deviceData{
		Title:   "",
		Devices: []DeviceMetrics{},
	}

	// create initial server device entry in set of all dashboard devices
	serverData := initializeDeviceMetrics(Server, checkInterval)

	dashboardData.Devices = append(dashboardData.Devices, *serverData)
	go dashboardConfig.Server.startDashboard(configPath)
}

// SendDashboardRequest manages dashboard requests and receipt of client device data
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
func (dash *serverKeyValues) startDashboard(configPath string) {
	dms3libs.LogDebug(filepath.Base(dms3libs.GetFunctionName()))

	// Template functions
	funcs := template.FuncMap{
		"ModVal":         dms3libs.ModVal,
		"FormatDateTime": dms3libs.FormatDateTime,
		"iconStatus":     iconStatus,
		"iconType":       iconType,
		"clientCount":    clientCount,
		"showEventCount": showEventCount,
	}

	// Parse template
	tmpl := template.Must(template.New(dash.Filename).Funcs(funcs).ParseFiles(filepath.Join(dash.FileLocation, dash.Filename)))

	// File server for static assets
	fs := http.FileServer(http.Dir(filepath.Join(configPath, dms3libs.DMS3Dashboard, "assets")))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Dashboard handler
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		handleDashboardRequest(w, tmpl, dash.Title)
	})

	// Configure HTTP server with timeouts
	server := &http.Server{
		Addr:         ":" + fmt.Sprint(dash.Port),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server
	dms3libs.LogInfo(fmt.Sprintf("Starting dashboard server on port %d", dash.Port))
	if err := server.ListenAndServe(); err != nil {
		dms3libs.LogFatal(err.Error())
	}
}

// handleDashboardRequest processes requests for the dashboard
func handleDashboardRequest(w http.ResponseWriter, tmpl *template.Template, title string) {

	dashboardData := &deviceData{
		Title:   title,
		Devices: dashboardData.Devices,
	}

	dashboardData.updateServerMetrics()

	if err := tmpl.Execute(w, dashboardData); err != nil {
		dms3libs.LogFatal(err.Error())
	}

}

// updateServerMetrics updates dynamic dashboard data of the server, triggered
// initially on dashboard start and subsequent webpage refreshes
func (dd *deviceData) updateServerMetrics() {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	for i := range dd.Devices {

		if dd.Devices[i].Platform.Type == Server {
			dd.Devices[i].Period.LastReport = time.Now()
			dd.Devices[i].Period.Uptime = dms3libs.Uptime(dd.Devices[i].Period.StartTime)
		} else {
			// check for and remove dead (non-reporting) client devices
			lastUpdate := dms3libs.SecondsSince(dd.Devices[i].Period.LastReport)
			missingDeviceLimit := uint32(dd.Devices[i].Period.CheckInterval) * dashboardConfig.Server.DeviceStatus.Missing

			if lastUpdate > missingDeviceLimit {
				dms3libs.LogInfo("Non-reporting remote device timeout reached: removing " + dd.Devices[i].Platform.Hostname + " client")
				dd.Devices = append(dd.Devices[:i], dd.Devices[i+1:]...)

				break
			}

		}
	}

}

// sendDashboardEnableState asks clients to send client info based on dashboard state
func sendDashboardEnableState(conn net.Conn, enableState string) {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	if _, err := conn.Write([]byte(enableState)); err != nil {
		dms3libs.LogFatal(err.Error())
	} else {
		dms3libs.LogInfo("Sent dashboard enable state as: " + enableState)
	}

}

// receiveDashboardData receives and parses client dashboard metrics
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
func (udm *DeviceMetrics) updateDeviceMetrics() {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	// scan for existing client device
	for i := range dashboardData.Devices {

		if dashboardData.Devices[i].Platform.Hostname == udm.Platform.Hostname {

			if dashboardData.Devices[i].Platform.Type == Client {
				dashboardData.Devices[i].EventCount = udm.EventCount
				dashboardData.Devices[i].Period.LastReport = udm.Period.LastReport
				dashboardData.Devices[i].Period.Uptime = udm.Period.Uptime
				dashboardData.Devices[i].Platform.OSName = udm.Platform.OSName
				dashboardData.Devices[i].Platform.Environment = udm.Platform.Environment
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

	if dashboardConfig.Server.ServerFirst {

		// place server in first dashboard position
		for i := range dashboardData.Devices {

			if dashboardData.Devices[i].Platform.Type == Server {
				server := dashboardData.Devices[i]
				dashboardData.Devices = append(dashboardData.Devices[:i], dashboardData.Devices[i+1:]...)
				dashboardData.Devices = append([]DeviceMetrics{server}, dashboardData.Devices...)
			}

		}

	}

}

// iconStatus is an HTML template function that returns the CSS string representing icon color,
// depending on the last time the client reported status to the server, relative to the client's
// CheckInterval
func iconStatus(index int) string {

	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

	lastUpdate := dms3libs.SecondsSince(dashboardData.Devices[index].Period.LastReport)
	checkInterval := dashboardData.Devices[index].Period.CheckInterval
	warningLimit := uint32(checkInterval) * dashboardConfig.Server.DeviceStatus.Caution
	dangerLimit := uint32(checkInterval) * dashboardConfig.Server.DeviceStatus.Danger

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
// func deviceType(index int) string {

// 	dms3libs.LogDebug(filepath.Base((dms3libs.GetFunctionName())))

// 	switch dashboardData.Devices[index].Platform.Type {
// 	case Client:
// 		return "client"
// 	case Server:
// 		return "server"
// 	default:
// 		return ""
// 	}

// }

// clientCount is an HTML template function that returns the current count of dms3clients
// reporting to the server
func clientCount() int {
	return len(dashboardData.Devices) - 1
}

// showEventCount is an HTML template function that returns whether to display client event count
func showEventCount(index int) bool {
	return dashboardData.Devices[index].ShowEventCount
}
