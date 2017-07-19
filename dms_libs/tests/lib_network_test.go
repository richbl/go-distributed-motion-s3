package dmslibs_test

import (
	"go_server/dms_libs"
	"os/exec"
	"testing"
)

func TestPingHosts(t *testing.T) {

	// ACTION: set active IP net address details
	testIPBase := "10.10.10."
	testIPRange := []int{100, 150}

	dmslibs.PingHosts(testIPBase, testIPRange)
}

func TestFindMacs(t *testing.T) {

	// ACTION: set active network interface (e.g., eth0)
	curInterface := "wlp2s0"

	var localMAC []string

	// determine local MAC address for testing
	cmd := dmslibs.SysCommand["CAT"] + " /sys/class/net/" + curInterface + "/address"
	res, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		t.Error("Unable to determine local MAC address for testing")
	}

	localMAC = append(localMAC, string(res))

	if !dmslibs.FindMacs(localMAC) {
		t.Error("FindMacs failed to find local MAC address")
	}
}
