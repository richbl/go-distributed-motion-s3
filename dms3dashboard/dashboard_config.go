// Package dms3dash dashboard configuration structures and variables
//
package dms3dash

import "time"

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
	DeviceStatus *serverDeviceStatus
}

// serverDeviceStatus represents the device status cycle in the TOML file
type serverDeviceStatus struct {
	Caution int
	Danger  int
	Missing int
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
	Environment string
	Kernel      string
}

// DeviceTime represents device time/duration metrics
type DeviceTime struct {
	CheckInterval int
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
