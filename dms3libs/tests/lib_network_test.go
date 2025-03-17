package dms3libs_test

import (
	"errors"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

var (
	ErrNoMacFound = errors.New("no MAC address found")
)

func init() {
	dms3libs.LoadLibConfig(filepath.Join("..", "..", dms3libs.DMS3Config, dms3libs.DMS3libsTOML))
}

func TestPingHosts(t *testing.T) {

	// ACTION: set active IP net address details
	testIPBase := "10.10.10."
	testIPRange := []int{100, 150}

	if err := dms3libs.PingHosts(testIPBase, testIPRange); err != nil {
		t.Error("PingHosts failed to ping hosts", err)
	}

}

// TestFindMacs tests the FindMacs function by first finding a local MAC address to establish
// a true case, and then passing it into the FindMacs() function
func TestFindMacs(t *testing.T) {

	var localMAC []string

	// Get a local MAC address for testing
	if res, err := getMACAddress(); err != nil {
		t.Error("Unable to determine local MAC address for testing")
	} else {
		localMAC = append(localMAC, res)
	}

	// Test the FindMacs function
	if !dms3libs.FindMacs(localMAC) {
		t.Error("FindMacs failed to find local MAC address")
	}

}

// getMACAddress returns the MAC address of the first device found on the network
func getMACAddress() (string, error) {

	// Execute the `ip neigh` command
	var out []byte
	var err error
	out, err = dms3libs.RunCommand(dms3libs.LibConfig.SysCommands["IP"] + " neigh")
	if err != nil {
		return "", err
	}

	// Define a regex pattern to match MAC addresses
	macRegex := regexp.MustCompile(`([0-9a-fA-F]{2}[:-]){5}([0-9a-fA-F]{2})`)
	matches := macRegex.FindAllString(string(out), -1)

	// Check if any MAC addresses were found
	if len(matches) == 0 {
		return "", ErrNoMacFound
	}

	return strings.TrimSpace(matches[0]), nil
}
