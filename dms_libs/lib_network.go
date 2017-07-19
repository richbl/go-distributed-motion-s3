package dmslibs

import (
	"strconv"
	"sync"
)

// PingHosts uses native ping command to ping the address range passed in to freshen the local arp cache
func PingHosts(ipBase string, ipRange []int) {
	var wg sync.WaitGroup
	cmd := SysCommand["PING"] + " -q -W 1 -c 1 " + ipBase

	for i := ipRange[0]; i < ipRange[1]; i++ {
		wg.Add(1)

		// allow threaded system command calls to finish asynchronously
		go func(i int) {
			defer wg.Done()
			RunCommand(cmd + strconv.Itoa(i))
		}(i)
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

	res, err := RunCommand(SysCommand["ARP"] + " -n | " + SysCommand["GREP"] + " -E '" + macListRegex + "'")

	if err != nil {
		Info.Println(SysCommand["ARP"], "command failed")
	}

	return (len(string(res)) > 0)
}

// FindMacs2 is a slower but non-threaded alternative to calling PingHosts + FindMacs, using the arp-scan utility
func FindMacs2(macsToFind []string, ipBase string, ipRange []int) bool {
	macListRegex := ""

	for i := 0; i < len(macsToFind); i++ {
		macListRegex += macsToFind[i]

		if i < len(macsToFind)-1 {
			macListRegex += "|"
		}

	}

	res, err := RunCommand("sudo " + SysCommand["ARPSCAN"] + " -q " + ipBase + strconv.Itoa(ipRange[0]) + "-" + ipBase + strconv.Itoa(ipRange[1]) + " | " + SysCommand["GREP"] + " -E '" + macListRegex + "'")

	if err != nil {
		Info.Println(SysCommand["ARPSCAN"], "command failed")
	}

	return (len(string(res)) > 0)
}
