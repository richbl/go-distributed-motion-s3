package dms3libs_test

import (
	"go-distributed-motion-s3/dms3libs"
	"testing"
)

func init() {
	dms3libs.LoadLibConfig("../../config/dms3libs.toml")
}

func TestConfiguration(t *testing.T) {

	for k, v := range dms3libs.LibConfig.SysCommands {

		if dms3libs.IsFile(v) {
			t.Log(k, "confirmed at", v)
		} else {
			t.Error(k, "not found at", v)
		}

	}

}
