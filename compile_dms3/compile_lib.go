package dms3compile

import (
	"fmt"
	"go-distributed-motion-s3/dms3libs"
	"path/filepath"
)

// structPlatform contains build environment/platform details
type structPlatform struct {
	envName     string
	dirName     string
	compileTags string
}

// structComponent contains component details
type structComponent struct {
	srcName        string
	exeName        string
	dirName        string
	configFilename string
	compile        bool
}

var buildEnv = []structPlatform{
	{
		envName:     "linuxArm7",
		dirName:     "linux_arm7",
		compileTags: "GOOS=linux GOARCH=arm GOARM=7",
	},
	{
		envName:     "linuxAMD64",
		dirName:     "linux_amd64",
		compileTags: "GOOS=linux GOARCH=amd64",
	},
}

var components = []structComponent{
	{
		srcName:        "go_dms3client.go",
		exeName:        "go_dms3client",
		dirName:        "dms3client",
		configFilename: "dms3client.toml",
		compile:        true,
	},
	{
		srcName:        "go_dms3server.go",
		exeName:        "go_dms3server",
		dirName:        "dms3server",
		configFilename: "dms3server.toml",
		compile:        true,
	},
	{
		srcName:        "go_dms3mail.go",
		exeName:        "go_dms3mail",
		dirName:        "dms3mail",
		configFilename: "dms3mail.toml",
		compile:        true,
	},
	{
		srcName:        "",
		exeName:        "dms3libs",
		dirName:        "dms3libs",
		configFilename: "dms3libs.toml",
		compile:        false,
	},
}

// BuildReleaseFolder creates the directory structure for each platform passed into it
func BuildReleaseFolder(releaseDir string) {

	dms3libs.RmDir(releaseDir)

	for itr := range buildEnv {
		fmt.Print("creating release folder for " + buildEnv[itr].envName + " platform... ")
		dms3libs.MkDir(filepath.Join(releaseDir, buildEnv[itr].dirName))
		fmt.Println("success")
	}

	for itr := range components {
		fmt.Print("creating release folder for " + components[itr].exeName + " component... ")
		dirName := components[itr].dirName

		if components[itr].dirName == "dms3server" {
			dirName = filepath.Join(dirName, "media")
		}

		dms3libs.MkDir(filepath.Join(releaseDir, dirName))
		fmt.Println("success")
	}

	fmt.Println()

}

// BuildComponents compiles dms3 components for each platform passed into it
func BuildComponents(releaseDir string) {

	for itr := range buildEnv {
		fmt.Print("building dms3 components for " + buildEnv[itr].envName + " platform... ")

		for jtr := range components {

			if components[jtr].compile {
				_, err := dms3libs.RunCommand("env " + buildEnv[itr].compileTags + " go build -o " + filepath.Join(releaseDir, buildEnv[itr].dirName) + "/" + components[jtr].exeName + " " + components[jtr].srcName)
				dms3libs.CheckErr(err)
			}
		}

		fmt.Println("success")
	}
	fmt.Println()

}

// CopyServiceDaemons copies daemons into release folder
func CopyServiceDaemons(releaseDir string) {

	fmt.Print("copying dms3 service daemons into " + releaseDir + " folder... ")
	dms3libs.CopyFile("dms3client/daemons/systemd/dms3client.service", filepath.Join(releaseDir, "dms3client/dms3client.service"))
	dms3libs.CopyFile("dms3server/daemons/systemd/dms3server.service", filepath.Join(releaseDir, "dms3server/dms3server.service"))
	fmt.Println("success")
	fmt.Println()
}

// CopyMediaFiles copies dms3server media files into release folder
func CopyMediaFiles(releaseDir string) {

	fmt.Print("copying dms3server media files (WAV) into " + releaseDir + " folder... ")
	dms3libs.CopyFile("dms3server/media/motion_start.wav", filepath.Join(releaseDir, "dms3server/media/motion_start.wav"))
	dms3libs.CopyFile("dms3server/media/motion_stop.wav", filepath.Join(releaseDir, "dms3server/media/motion_stop.wav"))
	fmt.Println("success")
	fmt.Println()

}

// CopyConfigFiles copies dms3server media files into release folder
func CopyConfigFiles(releaseDir string) {

	fmt.Print("copying dms3 component config files (TOML) into " + releaseDir + " folder... ")

	for itr := range components {
		dms3libs.CopyFile(components[itr].configFilename, filepath.Join(releaseDir, components[itr].dirName+"/"+components[itr].configFilename))
	}

	fmt.Println("success")
	fmt.Println()

}

// CopyReleaseFolder copies the release folder into /etc/distributed-motion-s3
func CopyReleaseFolder(releaseDir string, confDir string) {

	fmt.Print("copying " + releaseDir + " folder into " + confDir + " folder... ")
	dms3libs.RmDir(confDir)
	dms3libs.CopyDir(releaseDir, confDir)

	for itr := range buildEnv {
		dms3libs.RmDir(filepath.Join(confDir, buildEnv[itr].dirName))
	}

	fmt.Println("success")
	fmt.Println()

}

// CopyComponents copies dms3 components into /usr/local/bin
func CopyComponents(releaseDir string, execDir string, platform string) {

	for itr := range buildEnv {

		if buildEnv[itr].envName == platform {
			fmt.Print("copying " + buildEnv[itr].envName + " dms3 components into " + execDir + " folder (root permissions expected)... ")
			dms3libs.CopyDir(filepath.Join(releaseDir, buildEnv[itr].dirName), execDir)
			fmt.Println("success")
			fmt.Println()
		}

	}

}
