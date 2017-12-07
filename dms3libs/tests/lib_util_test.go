package dms3libs_test

import (
	"go-distributed-motion-s3/dms3libs"
	"testing"
)

func init() {
	dms3libs.LoadLibConfig("../../config/dms3libs.toml")
}

func TestPrintFuncName(t *testing.T) {

	dms3libs.LogDebug(dms3libs.GetFunctionName())

}

func TestGetPackageDir(t *testing.T) {

	dms3libs.GetPackageDir()

}

func TestStripRet(t *testing.T) {

	testArray := []byte{50, 40, 30, 20, 10}

	res := dms3libs.StripRet(testArray)

	if len(res) != len(testArray)-1 {
		t.Error("command failed")
	}

}
