// this script will be copied to the dms3 device component platform, executed, and
// then deleted automatically
//
// NOTE: must be run with admin privileges
//
package main

import (
	"go-distributed-motion-s3/dms3libs"
	"path/filepath"
)

func main() {

	binaryInstallDir := "/usr/local/bin/"
	configInstallDir := "/etc/distributed-motion-s3"
	logDir := "/var/log/dms3"

	// stop existing upstart service (if running)
	dms3libs.RunCommand("service dms3server stop")

	// move binary files into binaryInstallDir
	dms3libs.CopyFile("dms3_release/go_dms3server", filepath.Join(binaryInstallDir, "go_dms3server"))
	_, err := dms3libs.RunCommand("chmod +x " + filepath.Join(binaryInstallDir, "go_dms3server"))
	dms3libs.CheckErr(err)

	// create log folder
	dms3libs.MkDir(logDir)

	// copy configuration files into configInstallDir
	dms3libs.MkDir(configInstallDir)
	dms3libs.CopyDir("dms3_release/dms3server", configInstallDir)
	dms3libs.CopyDir("dms3_release/dms3dashboard", configInstallDir)
	dms3libs.CopyDir("dms3_release/dms3libs", configInstallDir)
	dms3libs.RmDir("dms3_release")

	// restart upstart service
	dms3libs.RunCommand("service dms3server start")

}
