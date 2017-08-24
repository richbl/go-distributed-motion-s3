package dms3build

type structDevice struct {
	User     string
	Server   string
	Password string
	Port     string
}

// Clients represent dms3 smart device component platforms
var Clients = []structDevice{
	{
		User:     "pi",
		Server:   "picam-alpha.local",
		Password: "", // ssh cert used
		Port:     "22",
	},
	{
		User:     "pi",
		Server:   "picam-beta.local",
		Password: "", // ssh cert used
		Port:     "22",
	},
}

// Servers represent dms3 server component platforms
var Servers = []structDevice{
	{
		User:     "richbl",
		Server:   "backup.local",
		Password: "", // ssh cert used
		Port:     "22",
	},
}
