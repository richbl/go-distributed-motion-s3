package dmslibs

import (
	"os/exec"
	"strconv"
	"sync"
)

// PingHosts uses native ping command to ping the address range passed in to freshen the local arp cache
func PingHosts(ipBase string, ipRange []int) {

	var wg sync.WaitGroup
	cmd := SysCommands["PING"] + " -q -W 1 -c 1 " + ipBase
	for i := ipRange[0]; i <= ipRange[1]; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			exec.Command("bash", "-c", cmd+strconv.Itoa(i)).Output()
		}()

	}

	wg.Wait()
}

// FindMacs uses arp to find mac addressed passed in, returning true if any mac passed in is found (e.g., mac1 | mac2 | mac3)
func FindMacs(macsToFind []string) bool {

	macListRegex := ""
	for i := 0; i < len(macsToFind); i++ {
		macListRegex += macsToFind[i]
		if i < len(macsToFind)-1 {
			macListRegex += "|"
		}
	}

	cmd := SysCommands["ARP"] + " -n | " + SysCommands["GREP"] + " -E '" + macListRegex + "'"
	res, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		Info.Println(SysCommands["ARP"], "command failed")
	}

	return (len(string(res)) > 0)
}

// FindMacs2 is a non-threaded alternative to calling PingHosts + FindMacs, and uses the arp-scan utility
func FindMacs2(macsToFind []string, ipBase string, ipRange []int) bool {

	macListRegex := ""
	for i := 0; i < len(macsToFind); i++ {
		macListRegex += macsToFind[i]
		if i < len(macsToFind)-1 {
			macListRegex += "|"
		}
	}

	cmd := "sudo " + SysCommands["ARPSCAN"] + " -q " + ipBase + strconv.Itoa(ipRange[0]) + "-" + ipBase + strconv.Itoa(ipRange[1]) + " | " + SysCommands["GREP"] + " -E '" + macListRegex + "'"
	res, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		Info.Println(SysCommands["ARPSCAN"], "command failed")
	}

	return (len(string(res)) > 0)
}
