// compiles dms3 components into platform-specific Go binary executables and copies configuration
// and media files into a dms3_release folder
//
// the dms3_release folder is then used as the base object for performing dms3 component
// installation on dms3client(s) and dms3server(s) device platforms (see install_dms3.go for
// details)
package main

import (
	"path/filepath"

	"github.com/richbl/go-distributed-motion-s3/dms3build"
	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

func main() {

	dms3libs.LoadLibConfig(filepath.Join(dms3libs.DMS3Config, dms3libs.DMS3libsTOML))

	// create release folder
	dms3build.BuildReleaseFolder()

	// build platform-specific components into release folder
	dms3build.BuildComponents()

	// copy service daemons into release folder
	dms3build.CopyServiceDaemons()

	// copy dms3server media files into release folder
	dms3build.CopyMediaFiles()

	// copy dms3dashboard html file and assets into release folder
	dms3build.CopyComponents(dms3libs.DMS3Dashboard)

	// copy dms3mail html file and assets into release folder
	dms3build.CopyComponents(dms3libs.DMS3Mail)

	// copy TOML files into release folder
	dms3build.CopyConfigFiles()

}
