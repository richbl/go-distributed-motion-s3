// this script will be copied to the dms3 device component platform, executed,
// and then deleted automatically
//
// NOTE: must be run with admin privileges on the remote device
package main

import (
	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

func main() {

	// NOTE: this component is run on the remote client device (per dms3build.toml configuration)

	// Specific files and directories for the client
	binFiles := []string{
		dms3libs.DMS3Client,
		dms3libs.DMS3Mail,
	}
	configDirs := []string{
		dms3libs.DMS3Client,
		dms3libs.DMS3Dashboard,
		dms3libs.DMS3Libs,
		dms3libs.DMS3Mail,
	}

	// Call the shared installer function
	dms3libs.DeviceInstaller(binFiles, configDirs)
}
