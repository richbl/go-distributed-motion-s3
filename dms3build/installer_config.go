package dms3build

type structDevice struct {
	User                string
	Server              string
	SSHPassword         string
	RemoteAdminPassword string
	Port                string
	Platform            PlatformType
}

// Clients represent dms3 smart device component platforms
var Clients = []structDevice{
	{
		User:                "pi",
		Server:              "picam-alpha.local",
		SSHPassword:         "", // SSH certificate
		RemoteAdminPassword: "acm44655!!",
		Port:                "22",
		Platform:            LinuxArm7,
	},
	{
		User:                "pi",
		Server:              "picam-beta.local",
		SSHPassword:         "", // SSH certificate
		RemoteAdminPassword: "acm44655!!",
		Port:                "22",
		Platform:            LinuxArm7,
	},
}

// Servers represent dms3 server component platforms
var Servers = []structDevice{
	{
		User:                "richbl",
		Server:              "backup.local",
		SSHPassword:         "", // SSH certificate
		RemoteAdminPassword: "acm44655!!",
		Port:                "22",
		Platform:            LinuxAMD64,
	},
}
