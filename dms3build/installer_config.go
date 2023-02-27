// Package dms3build installer configuration structures and variables
package dms3build

// BuildConfig contains installation configuration settings read from TOML file
var BuildConfig *structSettings

// client-side configuration parameters
type structSettings struct {
	Clients map[string]structDevice
	Servers map[string]structDevice
}

type structDevice struct {
	User                string
	DeviceName          string
	SSHPassword         string
	RemoteAdminPassword string
	Port                int
	Platform            platformType
}
