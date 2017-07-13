package dmslibs_test

import (
	"fmt"
	"go_server/dms_libs"
	"testing"
)

func TestConfiguration(t *testing.T) {
	for k, v := range dmslibs.SysCommands {
		if dmslibs.IsFile(v) {
			fmt.Println(k, "confirmed at", v)
		} else {
			t.Error(k, "not found at", v)
		}
	}
}
