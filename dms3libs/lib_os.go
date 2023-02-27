// Package dms3libs OS provides operating system information services for dms3 device components
package dms3libs

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
)

// DeviceDetails defines the set of available device details available in GetDeviceDetails
type DeviceDetails int

// types of DMS3 devices
const (
	Sysname DeviceDetails = iota
	Machine
	Release
)

// GetDeviceHostname returns the name of the local machine
func GetDeviceHostname() string {

	name, err := os.Hostname()
	CheckErr(err)
	return name

}

// GetDeviceOSName returns the OS release name (NAME) and version ID (VERSION_ID) from a parse of
// the /etc/os-release file found in most Linux-based distributions
func GetDeviceOSName() string {

	result := "OS unknown"

	if file, err := os.Open(filepath.Join(string(filepath.Separator), "etc", "os-release")); err == nil {

		defer file.Close()
		scanner := bufio.NewScanner(file)

		nameRegx := regexp.MustCompile(`^NAME=(.*)$`)
		versionIDRegx := regexp.MustCompile(`^VERSION_ID=(.*)$`)
		osName := ""
		osVersion := ""

		for scanner.Scan() {

			if res := nameRegx.FindStringSubmatch(scanner.Text()); res != nil {
				osName = strings.Trim(res[1], `"`)
			} else if res := versionIDRegx.FindStringSubmatch(scanner.Text()); res != nil {
				osVersion = strings.Trim(res[1], `"`)
			}

		}

		if osName != "" && osVersion != "" {
			result = strings.ToLower(osName + " " + osVersion)
		}

	}

	return result
}

// GetDeviceDetails returns device details of the local machine
func GetDeviceDetails(element DeviceDetails) string {

	utsName, error := uname()
	CheckErr(error)

	var len int
	var buf [65]byte

	switch element {
	case Sysname:
		for ; utsName.Sysname[len] != 0; len++ {
			buf[len] = uint8(utsName.Sysname[len])
		}
	case Machine:
		for ; utsName.Machine[len] != 0; len++ {
			buf[len] = uint8(utsName.Machine[len])
		}
	case Release:
		for ; utsName.Release[len] != 0; len++ {
			buf[len] = uint8(utsName.Release[len])
		}
	default:
		LogFatal("invalid DeviceDetails element passed in")
	}

	return strings.ToLower(string(buf[:len]))
}

// uname returns the Utsname struct used to query system settings
func uname() (*syscall.Utsname, error) {

	uts := &syscall.Utsname{}

	if err := syscall.Uname(uts); err != nil {
		return nil, err
	}

	return uts, nil

}
