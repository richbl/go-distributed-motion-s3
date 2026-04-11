// Package dms3libs OS provides operating system information services for dms3 device components
package dms3libs

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"unsafe"
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

	utsName, err := uname()
	CheckErr(err)

	var rawString string

	// Slice the arrays ([:]) to pass them into the generic extraction helper
	switch element {
	case Sysname:
		rawString = extractString(utsName.Sysname[:])
	case Machine:
		rawString = extractString(utsName.Machine[:])
	case Release:
		rawString = extractString(utsName.Release[:])
	default:
		LogFatal("invalid DeviceDetails element passed in")
	}

	return strings.ToLower(rawString)
}

// extractString converts a null-terminated integer array (int8 or uint8) into a string
func extractString[T ~int8 | ~uint8](arr []T) string {

	if len(arr) == 0 {
		return ""
	}

	// Reinterpret the underlying memory directly as a byte slice, avoiding explicit numeric
	// type casts
	byteSlice := unsafe.Slice((*byte)(unsafe.Pointer(&arr[0])), len(arr))

	// Find the null terminator
	length := bytes.IndexByte(byteSlice, 0)
	if length == -1 {
		length = len(byteSlice) // Safe fallback if no null-terminator is present
	}

	return string(byteSlice[:length])
}

// uname returns the Utsname struct used to query system settings
func uname() (*syscall.Utsname, error) {

	uts := &syscall.Utsname{}

	if err := syscall.Uname(uts); err != nil {
		return nil, fmt.Errorf("uname call failed: %w", err)
	}

	return uts, nil
}
