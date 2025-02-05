// this script will be copied to the dms3 device component platform, executed,
// and then deleted automatically
//
// NOTE: must be run with admin privileges on the remote device
package main

import (
	"path/filepath"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

func main() {

	// NOTE: this component is run on the remote client device (per dms3build.toml configuration)

	// load libs config file from dms3_release folder on remote device
	dms3libs.LoadLibConfig(filepath.Join(dms3libs.DMS3Release, dms3libs.DMS3Config, dms3libs.DMS3Libs, dms3libs.DMS3TOML))

	binaryInstallDir := filepath.Join(string(filepath.Separator), "usr", "local", "bin")
	configInstallDir := filepath.Join(string(filepath.Separator), "etc", "distributed-motion-s3")
	logDir := filepath.Join(string(filepath.Separator), "var", "log", "dms3")

	// move binary files into binaryInstallDir
	dms3libs.CopyFile(filepath.Join(dms3libs.DMS3Release, "cmd", dms3libs.DMS3Client), filepath.Join(binaryInstallDir, dms3libs.DMS3Client))
	dms3libs.CopyFile(filepath.Join(dms3libs.DMS3Release, "cmd", dms3libs.DMS3Mail), filepath.Join(binaryInstallDir, dms3libs.DMS3Mail))

	// create log folder
	dms3libs.MkDir(logDir)

	// copy configuration files into configInstallDir
	dms3libs.MkDir(configInstallDir)
	dms3libs.CopyDir(filepath.Join(dms3libs.DMS3Release, dms3libs.DMS3Config, dms3libs.DMS3Client), configInstallDir)
	dms3libs.CopyDir(filepath.Join(dms3libs.DMS3Release, dms3libs.DMS3Config, dms3libs.DMS3Dashboard), configInstallDir)
	dms3libs.CopyDir(filepath.Join(dms3libs.DMS3Release, dms3libs.DMS3Config, dms3libs.DMS3Libs), configInstallDir)
	dms3libs.CopyDir(filepath.Join(dms3libs.DMS3Release, dms3libs.DMS3Config, dms3libs.DMS3Mail), configInstallDir)

	dms3libs.RmDir(dms3libs.DMS3Release)

}
