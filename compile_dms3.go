package main

import (
	"go-distributed-motion-s3/dms3build"
)

// compiles dms3 components into platform-specific Go binary executables and copies configuration
// and media files into a dms3_release folder
//
// the dms3_release folder is then used as the base object for performing dms3 component
// installation on dms3client(s) and dms3server(s) (see install_dms3.go for details)
//
func main() {

	const releaseDir = "dms3_release"

	// build platform-specific components into release folder
	dms3build.BuildReleaseFolder(releaseDir)
	dms3build.BuildComponents(releaseDir)

	// copy service daemons into release folder
	dms3build.CopyServiceDaemons(releaseDir)

	// copy remote installer scripts into release folder
	dms3build.CopyInstallerFiles(releaseDir)

	// copy dms3server media files into release folder
	dms3build.CopyMediaFiles(releaseDir)

	// copy TOML files into release folder
	dms3build.CopyConfigFiles(releaseDir)

}
