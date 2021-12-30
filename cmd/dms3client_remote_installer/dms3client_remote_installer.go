// this script will be copied to the dms3 device component platform, executed,
// and then deleted automatically
//
// NOTE: must be run with admin privileges on the remote device
//
package main

import (
	"path/filepath"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

func main() {

	// NOTE: this component is run on the remote client device (per dms3build.toml configuration)

	// load libs config file from dms3_release folder on remote device
	dms3libs.LoadLibConfig(filepath.Join("dms3_release", "config", "dms3libs", "dms3libs.toml"))

	binaryInstallDir := filepath.Join(string(filepath.Separator), "usr", "local", "bin")
	configInstallDir := filepath.Join(string(filepath.Separator), "etc", "distributed-motion-s3")
	logDir := filepath.Join(string(filepath.Separator), "var", "log", "dms3")

	// move binary files into binaryInstallDir
	dms3libs.CopyFile(filepath.Join("dms3_release", "cmd", "dms3client"), filepath.Join(binaryInstallDir, "dms3client"))
	dms3libs.CopyFile(filepath.Join("dms3_release", "cmd", "dms3mail"), filepath.Join(binaryInstallDir, "dms3mail"))

	// create log folder
	dms3libs.MkDir(logDir)

	// copy configuration files into configInstallDir
	dms3libs.MkDir(configInstallDir)
	dms3libs.CopyDir(filepath.Join("dms3_release", "config", "dms3client"), configInstallDir)
	dms3libs.CopyDir(filepath.Join("dms3_release", "config", "dms3dashboard"), configInstallDir)
	dms3libs.CopyDir(filepath.Join("dms3_release", "config", "dms3libs"), configInstallDir)
	dms3libs.CopyDir(filepath.Join("dms3_release", "config", "dms3mail"), configInstallDir)

	dms3libs.RmDir("dms3_release")

}
