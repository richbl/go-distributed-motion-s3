// Package dms3dash server implements a dms3server-based metrics dashboard for all dms3clients
//
package dms3dash

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"go-distributed-motion-s3/dms3libs"
	"html/template"
	"log"
	"net"
	"net/http"
	"path"
	"path/filepath"
	"time"
)

// InitDashboardServer configs the library and server configuration for the dashboard
func InitDashboardServer(configPath string, dm *DeviceMetrics) {

	dashboardConfig = new(tomlTables)
	dms3libs.LoadComponentConfig(&dashboardConfig, filepath.Join(configPath, "dms3dashboard/dms3dashboard.toml"))

	if dashboardConfig.Server.Enable == true {
		dashboardConfig.Server.setDashboardFileLocation(configPath)
		dashboardData = new(deviceData)
		dm.appendServerMetrics()
		go dashboardConfig.Server.startDashboard()
	}

}

// SendDashboardRequest manages dashboard requests and receipt of client device data
func SendDashboardRequest(conn net.Conn) {

	dashboardConfig.Server.sendDashboardEnableState(conn)

	if dashboardConfig.Server.Enable == true {
		dashboardConfig.Server.receiveDashboardData(conn)
	}

}

// setDashboardFileLocation sets the location of the HTML file used when displaying the dashboard
func (dash *serverKeyValues) setDashboardFileLocation(configPath string) {

	relPath := filepath.Join(configPath, "dms3dashboard")
	devPath := filepath.Join(path.Dir(dms3libs.GetPackageDir()), "dms3dashboard")
	fail := false

	if !dms3libs.IsFile(filepath.Join(dash.FileLocation, dash.Filename)) {

		// if no location set, set to release folder, else set to development folder
		if dash.FileLocation == "" {

			if dms3libs.IsFile(filepath.Join(relPath, dash.Filename)) {
				dash.FileLocation = relPath
			} else if dms3libs.IsFile(filepath.Join(devPath, dash.Filename)) {
				dash.FileLocation = devPath
			} else {
				fail = true
			}

		} else {
			fail = true
		}

		if fail {
			log.Fatalln("unable to set dashboard location... check TOML configuration file")
		}
	}

}

// startDashboard intializes and starts an HTTP server, serving the client dash on the server
func (dash *serverKeyValues) startDashboard() {

	funcs := template.FuncMap{
		"ModVal":         dms3libs.ModVal,
		"FormatDateTime": dms3libs.FormatDateTime,
		"iconStatus":     iconStatus,
		"iconType":       iconType,
		"clientCount":    clientCount,
	}

	tmpl := template.Must(template.New(dash.Filename).Funcs(funcs).ParseFiles(filepath.Join(dash.FileLocation, dash.Filename)))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("dms3dashboard/assets"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		dashboardData = &deviceData{
			Title:   dash.Title,
			Clients: dashboardData.Clients,
		}

		dashboardData.updateServerMetrics()
		tmpl.Execute(w, dashboardData)

	})

	http.ListenAndServe(":"+fmt.Sprint(dash.Port), nil)
}

// updateServerMetrics updates dynamic dashboard data of the server
func (dd *deviceData) updateServerMetrics() {

	dd.Clients[0].LastReport = time.Now()
	dd.Clients[0].Uptime = dms3libs.Uptime(dd.Clients[0].StartTime)

}

// sendDashboardEnableState asks clients to send client info based on dashboard state
func (dash *serverKeyValues) sendDashboardEnableState(conn net.Conn) {

	state := "0"

	if dash.Enable {
		state = "1"
	}

	if _, err := conn.Write([]byte(state)); err != nil {
		dms3libs.LogFatal(err.Error())
	} else {
		dms3libs.LogInfo("Sent dashboard enable state as: " + state)
	}
}

// receiveDashboardData receives and parses client dashboard metrics
func (dash *serverKeyValues) receiveDashboardData(conn net.Conn) {

	newClientMetrics := new(DeviceMetrics)
	buf := make([]byte, 1024)

	if n, err := conn.Read(buf); err != nil {
		dms3libs.LogFatal(err.Error())
	} else {
		// gob decoding of client metrics
		decBuf := bytes.NewBuffer(buf[:n])
		err = gob.NewDecoder(decBuf).Decode(newClientMetrics)

		newClientMetrics.appendClientMetrics()
	}

}

// appendClientMetrics adds new clients to the dashboard list, or updates existing client
// metrics, where Hostname is the unique key
//
func (dm *DeviceMetrics) appendClientMetrics() {

	for i := range dashboardData.Clients {

		if dashboardData.Clients[i].Type == Client {

			if dashboardData.Clients[i].Hostname == dm.Hostname {
				dashboardData.Clients[i].EventCount = dm.EventCount
				dashboardData.Clients[i].LastReport = dm.LastReport
				dashboardData.Clients[i].Uptime = dm.Uptime
				return
			}

		}

	}

	dashboardData.Clients = append(dashboardData.Clients, *dm)

}

//  appendServerMetrics appends the server to the dashboard list
func (dm *DeviceMetrics) appendServerMetrics() {

	serverData := new(DeviceMetrics)
	*serverData = *dm
	serverData.Hostname = dms3libs.DeviceHostname()
	serverData.Environment = dms3libs.DeviceOS() + " " + dms3libs.DevicePlatform()
	serverData.Kernel = dms3libs.DeviceKernel()

	dashboardData.Clients = append(dashboardData.Clients, *serverData)

}

// iconStatus is an HTML template function that returns the CSS string representing icon color,
// depending on the last time the client reported status to the server, relative to the client's
// CheckInterval
//
func iconStatus(index int) string {

	seconds := dms3libs.SecondsSince(dashboardData.Clients[index].LastReport)
	interval := dashboardData.Clients[index].CheckInterval
	warningLimit := uint32((interval * 2))
	dangerLimit := uint32((interval * 4))

	switch {
	case seconds < warningLimit:
		return "icon-success"
	case (seconds >= warningLimit) && (seconds < dangerLimit):
		return "icon-warning"
	case seconds >= dangerLimit:
		return "icon-danger"
	default:
		return ""
	}

}

// iconType is an HTML template function that returns an icon based on device type
func iconType(index int) string {

	switch dashboardData.Clients[index].Type {
	case Client:
		return "icon-raspberry-pi"
	case Server:
		return "icon-server2"
	default:
		return ""
	}

}

// clientCount is an HTML template function that returns the current count of dms3clients
// reporting to the server
//
func clientCount() int {
	return len(dashboardData.Clients) - 1
}
