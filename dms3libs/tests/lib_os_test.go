package dms3libs_test

import (
	"testing"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

func TestDeviceHostname(t *testing.T) {

	val := dms3libs.DeviceHostname()

	if val != "" {
		t.Log("Success, devicehost is", val)
	} else {
		t.Error("Failure. Unable to find devicehost")
	}

}

func TestDeviceOSName(t *testing.T) {

	val := dms3libs.DeviceOSName()

	if val != "" {
		t.Log("Success, device OS name is", val)
	} else {
		t.Error("Failure. Unable to find device OS name")
	}

}

func TestGetDeviceDetails(t *testing.T) {

	val := dms3libs.GetDeviceDetails(dms3libs.Sysname)

	if val != "" {
		t.Log("Success, device platform is", val)
	} else {
		t.Error("Failure. Unable to find device details")
	}

}
