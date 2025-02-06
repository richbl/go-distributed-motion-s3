// this script will be copied to the dms3 device component platform, executed,
// and then deleted automatically
//
// NOTE: must be run with admin privileges on the remote device
package main

import (
	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

func main() {

	// NOTE: this component is run on the remote server device (per dms3build.toml configuration)

	// Specific files and directories for the server
	binFiles := []string{
		dms3libs.DMS3Server,
	}
	configDirs := []string{
		dms3libs.DMS3Server,
		dms3libs.DMS3Dashboard,
		dms3libs.DMS3Libs,
	}

	// Call the shared Setup function
	dms3libs.DeviceInstaller(binFiles, configDirs)

}
