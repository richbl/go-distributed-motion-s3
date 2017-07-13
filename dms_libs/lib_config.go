package dmslibs

// SysCommands provides a location mapping of required system/application commands
var SysCommands = map[string]string{
	"APLAY":   "/usr/bin/aplay",
	"ARP":     "/usr/sbin/arp",
	"ARPSCAN": "/usr/bin/arp-scan", // TODO not currently used
	"GREP":    "/bin/grep",
	"PGREP":   "/usr/bin/pgrep",
	"PING":    "/bin/ping",
	"MOTION":  "/usr/bin/motion",
}
