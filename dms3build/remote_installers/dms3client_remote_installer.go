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

	// stop existing systemd service (if running)
	dms3libs.RunCommand("systemctl stop dms3client.service")

	// move binary files into binaryInstallDir
	dms3libs.CopyFile("dms3_release/go_dms3client", filepath.Join(binaryInstallDir, "go_dms3client"))
	_, err := dms3libs.RunCommand("chmod +x " + filepath.Join(binaryInstallDir, "go_dms3client"))
	dms3libs.CheckErr(err)

	dms3libs.CopyFile("dms3_release/go_dms3mail", filepath.Join(binaryInstallDir, "go_dms3mail"))
	_, err = dms3libs.RunCommand("chmod +x " + filepath.Join(binaryInstallDir, "go_dms3mail"))
	dms3libs.CheckErr(err)

	// copy configuration files into configInstallDir
	dms3libs.MkDir(configInstallDir)
	dms3libs.CopyDir("dms3_release/dms3client", configInstallDir)
	dms3libs.CopyDir("dms3_release/dms3dashboard", configInstallDir)
	dms3libs.CopyDir("dms3_release/dms3libs", configInstallDir)
	dms3libs.CopyDir("dms3_release/dms3mail", configInstallDir)
	dms3libs.RmDir("dms3_release")

	// restart systemd service
	dms3libs.RunCommand("systemctl start dms3client.service")

}
