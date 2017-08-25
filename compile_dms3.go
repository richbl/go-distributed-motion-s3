package main

import (
	"go-distributed-motion-s3/dms3build"
)

func main() {

	const (
		execDir    = "/usr/local/bin"             // set to anywhere seen by $PATH
		confDir    = "/etc/distributed-motion-s3" // default location to store config files (*.toml)
		releaseDir = "dms3_release"               // release folder containing all platform builds
	)

	// build platform-specific components into release folder
	//
	dms3build.BuildReleaseFolder(releaseDir)
	dms3build.BuildComponents(releaseDir)

	// copy service daemons into release folder
	//
	dms3build.CopyServiceDaemons(releaseDir)

	// copy dms3server media files into release folder
	//
	dms3build.CopyMediaFiles(releaseDir)

	// copy TOML files into release folder
	//
	dms3build.CopyConfigFiles(releaseDir)

	// copy release folder into /etc/distributed-motion-s3
	//
	// dms3build.CopyReleaseFolder(releaseDir, confDir)

	// copy compiled dms3 components into /usr/local/bin
	//
	// dms3build.CopyComponents(releaseDir, execDir, "linuxAMD64")
	// dms3build.CopyComponents(releaseDir, execDir, build, "linuxArm7")

}
