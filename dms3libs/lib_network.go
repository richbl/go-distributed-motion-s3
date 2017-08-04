package dms3libs

import (
	"strconv"
	"sync"
)

// PingHosts uses native ping command to ping the address range passed in to freshen the local
// arp cache
//
func PingHosts(ipBase string, ipRange []int) {

	var wg sync.WaitGroup
	cmd := LibConfig.SysCommands.PING + " -q -W 1 -c 1 " + ipBase

	for i := ipRange[0]; i < ipRange[1]; i++ {
		wg.Add(1)

		// allow threaded system command calls to finish asynchronously
		go func(i int, w *sync.WaitGroup) {
			defer w.Done()
			RunCommand(cmd + strconv.Itoa(i))
		}(i, &wg)
	}

	wg.Wait()

}

// FindMacs uses arp to find mac addressed passed in, returning true if any mac passed in is found
// (e.g., mac1 | mac2 | mac3)
//
func FindMacs(macsToFind []string) bool {
	macListRegex := ""

	for i := 0; i < len(macsToFind); i++ {
		macListRegex += macsToFind[i]

		if i < len(macsToFind)-1 {
			macListRegex += "|"
		}

	}

	res, err := RunCommand(LibConfig.SysCommands.ARP + " -n | " + LibConfig.SysCommands.GREP + " -E '" + macListRegex + "'")

	if err != nil {
		LogInfo(LibConfig.SysCommands.ARP + " command code: " + err.Error())
	}

	return (len(string(res)) > 0)

}
