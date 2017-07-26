package dms3libs_test

import (
	"fmt"
	"go-distributed-motion-s3/dms3libs"
	"testing"
)

func TestConfiguration(t *testing.T) {

	for k, v := range dms3libs.SysCommand {

		if dms3libs.IsFile(v) {
			fmt.Println(k, "confirmed at", v)
		} else {
			t.Error(k, "not found at", v)
		}

	}

}
