// this script will be copied to the dms3 device component platform, executed, and
// then deleted automatically
//
// NOTE: must be run with admin privileges on the remote device
//
package main

import (
	"path/filepath"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

func main() {

	binaryInstallDir := "/usr/local/bin/"
	configInstallDir := "/etc/distributed-motion-s3"
	logDir := "/var/log/dms3"

	// stop existing service (if running)
	_, err := dms3libs.RunCommand(dms3libs.LibConfig.SysCommands["SERVICE"] + " dms3server stop")
	dms3libs.CheckErr(err)

	// move binary files into binaryInstallDir
	dms3libs.CopyFile("dms3_release/dms3server", filepath.Join(binaryInstallDir, "dms3server"))
	_, err = dms3libs.RunCommand(dms3libs.LibConfig.SysCommands["CHMOD"] + " +x " + filepath.Join(binaryInstallDir, "dms3server"))
	dms3libs.CheckErr(err)

	// create log folder
	dms3libs.MkDir(logDir)

	// copy configuration files into configInstallDir
	dms3libs.MkDir(configInstallDir)
	dms3libs.CopyDir("dms3_release/dms3server", configInstallDir)
	dms3libs.CopyDir("dms3_release/dms3dashboard", configInstallDir)
	dms3libs.CopyDir("dms3_release/dms3libs", configInstallDir)
	dms3libs.RmDir("dms3_release")

	// restart service
	_, err = dms3libs.RunCommand(dms3libs.LibConfig.SysCommands["SERVICE"] + " dms3server start")
	dms3libs.CheckErr(err)

}
