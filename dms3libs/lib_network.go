// Package dms3libs network provides networking services for dms3 device components
package dms3libs

import (
	"os/exec"
	"strconv"
	"sync"
)

// PingHosts uses native ping command to ping the address range passed in to freshen the local
// arp cache
func PingHosts(ipBase string, ipRange []int) {

	var wg sync.WaitGroup
	cmd := LibConfig.SysCommands["PING"] + " -q -W 1 -c 1 " + ipBase

	for i := ipRange[0]; i <= ipRange[1]; i++ {
		wg.Add(1)

		// allow threaded system command calls to finish asynchronously
		go func(i int, w *sync.WaitGroup) {
			defer w.Done()
			_, _ = RunCommand(cmd + strconv.Itoa(i))
		}(i, &wg)
	}

	wg.Wait()

}

// FindMacs uses 'ip neigh' to find mac address(es) passed in, returning true if any mac passed in is found
// (e.g., mac1|mac2|mac3)
func FindMacs(macsToFind []string) bool {

	macListRegex := ""

	for i := 0; i < len(macsToFind); i++ {
		macListRegex += macsToFind[i]

		if i < len(macsToFind)-1 {
			macListRegex += "|"
		}

	}

	if _, err := RunCommand(LibConfig.SysCommands["IP"] + " neigh | " + LibConfig.SysCommands["GREP"] + " -iE '" + macListRegex + "'"); err != nil {

		switch err.(type) {
		case *exec.ExitError:
			// no ip found
			LogInfo(LibConfig.SysCommands["IP"] + " command: no device mac address found")
		default:
			// fatal command error
			LogFatal("Failed to run '" + LibConfig.SysCommands["IP"] + " neigh | " + LibConfig.SysCommands["GREP"] + " -iE '" + macListRegex + "': " + err.Error())
		}

		return false
	} else {
		LogInfo(LibConfig.SysCommands["IP"] + " command: device mac address found")
		return true
	}

}
