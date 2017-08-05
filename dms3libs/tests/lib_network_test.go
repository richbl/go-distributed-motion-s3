package dms3libs_test

import (
	"go-distributed-motion-s3/dms3libs"
	"os/exec"
	"testing"
)

func init() {
	dms3libs.LoadLibConfig("../../dms3libs/lib_config.toml")
}

func TestPingHosts(t *testing.T) {

	// ACTION: set active IP net address details
	testIPBase := "10.10.10."
	testIPRange := []int{100, 150}

	dms3libs.PingHosts(testIPBase, testIPRange)

}

func TestFindMacs(t *testing.T) {

	// ACTION: set active network interface (e.g., eth0)
	curInterface := "wlp2s0"

	var localMAC []string

	// determine local MAC address for testing
	cmd := dms3libs.LibConfig.SysCommands["CAT"] + " /sys/class/net/" + curInterface + "/address"
	if res, err := exec.Command("bash", "-c", cmd).Output(); err != nil {
		t.Error("Unable to determine local MAC address for testing")
	} else {
		localMAC = append(localMAC, string(res))
	}

	if !dms3libs.FindMacs(localMAC) {
		t.Error("FindMacs failed to find local MAC address")
	}

}
