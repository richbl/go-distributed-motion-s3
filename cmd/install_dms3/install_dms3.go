// installs dms3 components (dms3client, dms3server, and dms3mail) and supporting configuration,
// service daemons, and media files onto the specified dms3 device component platforms (see
// dms3build.toml for a list of platforms to install onto)
//
// this installer depends on a local dms3_release folder created through dms3 compilation (see
// compile_dms3.go for details)
package main

import (
	"path/filepath"

	"github.com/richbl/go-distributed-motion-s3/dms3build"
	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

func main() {

	releasePath := "dms3_release"

	// confirm existence of dms3_release folder based on releasePath
	dms3build.ConfirmReleaseFolder(releasePath)

	// read configuration in from dms3build TOML
	dms3libs.LoadComponentConfig(&dms3build.BuildConfig, filepath.Join(releasePath, "config", "dms3build", "dms3build.toml"))

	// install components onto device platforms
	dms3build.InstallClientComponents(releasePath)
	dms3build.InstallServerComponents(releasePath)

}
