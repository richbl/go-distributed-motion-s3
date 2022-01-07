// Package dms3libs OS provides operating system information services for dms3 device components
//
package dms3libs

import (
	"os"
	"runtime"
	"syscall"
)

// DeviceHostname returns the name of the local machine
//
func DeviceHostname() string {

	name, err := os.Hostname()
	CheckErr(err)
	return name

}

// DeviceOS returns the operating system of the local machine
//
func DeviceOS() string {
	return runtime.GOOS
}

// DevicePlatform returns the CPU architecture of the local machine
//
func DevicePlatform() string {
	return runtime.GOARCH
}

// DeviceKernel returns the current kernel in use on the local machine
//
func DeviceKernel() string {

	utsName, error := uname()
	CheckErr(error)

	var len int
	var buf [65]byte

	for ; utsName.Release[len] != 0; len++ {
		buf[len] = uint8(utsName.Release[len])
	}

	return string(buf[:len])
}

// uname returns the Utsname struct used to query system settings
//
func uname() (*syscall.Utsname, error) {

	uts := &syscall.Utsname{}

	if err := syscall.Uname(uts); err != nil {
		return nil, err
	}

	return uts, nil

}
