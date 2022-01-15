package dms3libs_test

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

func TestGetFunctionName(t *testing.T) {

	val := filepath.Base(dms3libs.GetFunctionName())

	if val != "" {
		t.Log("Success, function name is", val)
	} else {
		t.Error("Failure. Unable to get function name")
	}

}

func TestGetPackageDir(t *testing.T) {

	val := dms3libs.GetPackageDir()

	if val != "" {
		t.Log("Success, package dir is", val)
	} else {
		t.Error("Failure. Unable to get package dir")
	}

}

func TestStripRet(t *testing.T) {

	testArray := []byte{50, 40, 30, 20, 10}
	res := dms3libs.StripRet(testArray)

	if len(res) != len(testArray)-1 {
		t.Error("function failed")
	}

}

func TestSetUptime(t *testing.T) {

	now := time.Now()
	then := now
	dms3libs.SetUptime(&then)

	if now.Sub(then).Seconds() < 0.1 {
		t.Log("Success")
	} else {
		t.Error("Test for SetUptime failed")
	}

}

func TestUptime(t *testing.T) {

	testTime := new(time.Time)
	dms3libs.SetUptime(testTime)
	time.Sleep(1000 * time.Millisecond) // force uptime value
	val := dms3libs.Uptime(*testTime)

	if val != "000d:00h:00m:00s" {
		t.Log("Success, uptime is", val)
	} else {
		t.Error("Failure. Unable to get uptime")
	}

}

func TestSecondsSince(t *testing.T) {

	testTime := time.Now()
	time.Sleep(1000 * time.Millisecond)
	val := dms3libs.SecondsSince(testTime)

	if val >= 1 {
		t.Log("Success, seconds since", testTime, "is", val)
	} else {
		t.Error("Failure. Unable to get seconds since", testTime)
	}

}

func TestTo24H(t *testing.T) {

	testTime := time.Now()
	val := dms3libs.To24H(testTime)

	if val == testTime.Format("150405") {
		t.Log("Success, current time to 24H format is", val)
	} else {
		t.Error("Failure. Unable to convert current time")
	}

}

func TestFormatDateTime(t *testing.T) {

	testTime := time.Now()
	val := dms3libs.FormatDateTime(testTime)

	if val == testTime.Format("2006-01-02 at 15:04:05") {
		t.Log("Success, current time formatted is", val)
	} else {
		t.Error("Failure. Unable to format time")
	}

}

func TestModVal(t *testing.T) {

	dividend := 30
	divisor := 10
	val := dms3libs.ModVal(dividend, divisor)

	if val == 0 {
		t.Log("Success, ModVal(", dividend, ",", divisor, ") is", val)
	} else {
		t.Error("Failure. ModVal(", dividend, ",", divisor, ") yielded", val)
	}

}

func TestGetImageDimensions(t *testing.T) {

	imageFile := filepath.Join(dms3libs.GetPackageDir(), "lib_util_test.jpg")
	w, h := dms3libs.GetImageDimensions(imageFile)

	if w != 100 && h != 100 {
		t.Error("Failure. Dimensions do not match image")
	}

}
