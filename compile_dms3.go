package main

import (
	"go-distributed-motion-s3/compile_dms3"
	"os"
)

func main() {

	const (
		execDir    = "/usr/local/bin"             // set to anywhere seen by $PATH
		confDir    = "/etc/distributed-motion-s3" // default location to store config files (*.toml)
		releaseDir = "dms3_release"               // release folder containing all platform builds
	)

	// build platform-specific components into release folder
	//
	dms3compile.BuildReleaseFolder(releaseDir)
	dms3compile.BuildComponents(releaseDir)

	// copy service daemons into release folder
	//
	dms3compile.CopyServiceDaemons(releaseDir)

	// copy dms3server media files into release folder
	//
	dms3compile.CopyMediaFiles(releaseDir)

	// copy TOML files into release folder
	//
	dms3compile.CopyConfigFiles(releaseDir)

	os.Exit(0) // TODO no need to copy release locally

	// copy release folder into /etc/distributed-motion-s3
	//
	dms3compile.CopyReleaseFolder(releaseDir, confDir)

	// copy compiled dms3 components into /usr/local/bin
	//
	dms3compile.CopyComponents(releaseDir, execDir, "linuxAMD64")
	// dms3compile.CopyComponents(releaseDir, execDir, build, "linuxArm7")

}
