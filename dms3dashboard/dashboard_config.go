// Package dms3dash dashboard configuration structures and variables
package dms3dash

import (
	"time"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

var DashboardEnable bool

var dashboardConfig *tomlTables
var dashboardData *deviceData

// tomlTables represents the TOML table(s)
type tomlTables struct {
	Client *clientKeyValues
	Server *serverKeyValues
}

// clientKeyValues represents the k-v pairs in the TOML file
type clientKeyValues struct {
	ImagesFolder string
}

// serverKeyValues represents the k-v pairs in the TOML file
type serverKeyValues struct {
	Port         int
	Filename     string
	FileLocation string
	Title        string
	ReSort       bool
	ServerFirst  bool
	DeviceStatus *serverDeviceStatus
}

// serverDeviceStatus represents the device status cycle in the TOML file
type serverDeviceStatus struct {
	Caution uint32
	Danger  uint32
	Missing uint32
}

// deviceData represents dashboard elements from all devices
type deviceData struct {
	Title   string
	Devices []DeviceMetrics
}

// DeviceMetrics represents device data presented on the dashboard
type DeviceMetrics struct {
	Platform       DevicePlatform
	Period         DeviceTime
	ShowEventCount bool
	EventCount     int
}

// DevicePlatform represents the physical device platform environment
type DevicePlatform struct {
	Type        dashboardDeviceType
	Hostname    string
	OSName      string
	Environment string
	Kernel      string
}

// DeviceTime represents device time/duration metrics
type DeviceTime struct {
	CheckInterval uint16
	StartTime     time.Time
	Uptime        string
	LastReport    time.Time
}

// dashboardDeviceType defines the dashboard device type
type dashboardDeviceType int

// types of DMS3 devices
const (
	Client dashboardDeviceType = iota
	Server
)

// initializeDeviceMetrics is a helper function to initialize a DeviceMetrics struct.
func initializeDeviceMetrics(deviceType dashboardDeviceType, checkInterval uint16) *DeviceMetrics {
	return &DeviceMetrics{
		Platform: DevicePlatform{
			Type:        deviceType,
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
}
