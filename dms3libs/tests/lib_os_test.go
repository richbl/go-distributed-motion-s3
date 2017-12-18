package dms3libs_test

import (
	"go-distributed-motion-s3/dms3libs"
	"testing"
)

func init() {
	dms3libs.LoadLibConfig("../../config/dms3libs.toml")
}

func TestDeviceHostname(t *testing.T) {

	val := dms3libs.DeviceHostname()

	if val != "" {
		t.Log("Success, devicehost is", val)
	} else {
		t.Error("Failure. Unable to find devicehost")
	}

}

func TestDeviceOS(t *testing.T) {

	val := dms3libs.DeviceOS()

	if val != "" {
		t.Log("Success, device OS is", val)
	} else {
		t.Error("Failure. Unable to find deviceOS")
	}

}

func TestDevicePlatform(t *testing.T) {

	val := dms3libs.DevicePlatform()

	if val != "" {
		t.Log("Success, device platform is", val)
	} else {
		t.Error("Failure. Unable to find device platform")
	}

}

func TestDeviceKernel(t *testing.T) {

	val := dms3libs.DeviceKernel()

	if val != "" {
		t.Log("Success, device kernel is", val)
	} else {
		t.Error("Failure. Unable to find device kernel")
	}

}
