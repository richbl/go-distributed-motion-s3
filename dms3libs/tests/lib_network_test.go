package dms3libs_test

import (
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

func init() {
	dms3libs.LoadLibConfig(filepath.Join("..", "..", "config", "dms3libs.toml"))
}

func TestPingHosts(t *testing.T) {

	// ACTION: set active IP net address details
	testIPBase := "10.10.10."
	testIPRange := []int{100, 150}

	dms3libs.PingHosts(testIPBase, testIPRange)

}

func TestFindMacs(t *testing.T) {

	// ACTION: set active network interface (e.g., eth0)
	curInterface := "wlp9s0"

	var localMAC []string

	// determine local MAC address for testing
	cmd := dms3libs.LibConfig.SysCommands["CAT"] + " /sys/class/net/" + curInterface + "/address"

	if res, err := exec.Command(dms3libs.LibConfig.SysCommands["BASH"], "-c", cmd).Output(); err != nil {
		t.Error("Unable to determine local MAC address for testing")
	} else {
		localMAC = append(localMAC, string(res))
	}

	if !dms3libs.FindMacs(localMAC) {
		t.Error("FindMacs failed to find local MAC address")
	}

}
