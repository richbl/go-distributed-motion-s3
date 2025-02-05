// Package dms3build compiler configuration structures and variables
package dms3build

import "github.com/richbl/go-distributed-motion-s3/dms3libs"

// structComponent contains component details
type structComponent struct {
	srcName        string // component source filename
	exeName        string // compiled filename
	dirName        string // location for components in release folder
	configFilename string // relevant config (TOML) file
	compile        bool   // whether component should be compiled
}

var components = []structComponent{
	{
		srcName:        "cmd/dms3client/dms3client.go",
		exeName:        dms3libs.DMS3Client,
		dirName:        dms3libs.DMS3Client,
		configFilename: dms3libs.DMS3TOML,
		compile:        true,
	},
	{
		srcName:        "cmd/dms3server/dms3server.go",
		exeName:        dms3libs.DMS3Server,
		dirName:        dms3libs.DMS3Server,
		configFilename: "dms3server.toml",
		compile:        true,
	},
	{
		srcName:        "cmd/dms3mail/dms3mail.go",
		exeName:        dms3libs.DMS3Mail,
		dirName:        dms3libs.DMS3Mail,
		configFilename: "dms3mail.toml",
		compile:        true,
	},
	{
		srcName:        "cmd/install_dms3/install_dms3.go",
		exeName:        "install_dms3",
		dirName:        "dms3build",
		configFilename: "dms3build.toml",
		compile:        true,
	},
	{
		srcName:        "cmd/dms3client_remote_installer/dms3client_remote_installer.go",
		exeName:        "dms3client_remote_installer",
		dirName:        "dms3build",
		configFilename: "",
		compile:        true,
	},
	{
		srcName:        "cmd/dms3server_remote_installer/dms3server_remote_installer.go",
		exeName:        "dms3server_remote_installer",
		dirName:        "dms3build",
		configFilename: "",
		compile:        true,
	},
	{
		srcName:        "",
		exeName:        dms3libs.DMS3Libs,
		dirName:        dms3libs.DMS3Libs,
		configFilename: "dms3libs.toml",
		compile:        false,
	},
	{
		srcName:        "",
		exeName:        dms3libs.DMS3Dashboard,
		dirName:        dms3libs.DMS3Dashboard,
		configFilename: "dms3dashboard.toml",
		compile:        false,
	},
}

// platformType represents all available platform types
type platformType string

// platform types
const (
	linuxArm6  platformType = "linuxArm6"
	linuxArm7  platformType = "linuxArm7"
	linuxArm8  platformType = "linuxArm8"
	linuxAMD64 platformType = "linuxAMD64"
)

// structPlatform contains build environment/platform details
type structPlatform struct {
	dirName     string
	compileTags string
}

// BuildEnv contains platform build details
var BuildEnv = map[platformType]structPlatform{
	linuxArm6: {
		dirName:     "linux_arm6",
		compileTags: "GOOS=linux GOARCH=arm GOARM=6",
	},
	linuxArm7: {
		dirName:     "linux_arm7",
		compileTags: "GOOS=linux GOARCH=arm GOARM=7",
	},
	linuxArm8: {
		dirName:     "linux_arm8",
		compileTags: "GOOS=linux GOARCH=arm64",
	},
	linuxAMD64: {
		dirName:     "linux_amd64",
		compileTags: "GOOS=linux GOARCH=amd64",
	},
}
