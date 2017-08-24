package dms3build

import (
	"fmt"
	"go-distributed-motion-s3/dms3libs"
	"path/filepath"
	"strings"

	"github.com/hypersleep/easyssh"
)

// BuildReleaseFolder creates the directory structure for each platform passed into it
func BuildReleaseFolder(releaseDir string) {

	dms3libs.RmDir(releaseDir)

	for itr := range buildEnv {
		fmt.Print("Creating release folder for " + buildEnv[itr].envName + " platform... ")
		dms3libs.MkDir(filepath.Join(releaseDir, buildEnv[itr].dirName))
		fmt.Println("Success")
	}

	for itr := range components {
		fmt.Print("Creating release folder for " + components[itr].exeName + " component... ")
		dirName := components[itr].dirName

		if components[itr].dirName == "dms3server" {
			dirName = filepath.Join(dirName, "media")
		}

		dms3libs.MkDir(filepath.Join(releaseDir, dirName))
		fmt.Println("Success")
	}

	fmt.Println()

}

// BuildComponents compiles dms3 components for each platform passed into it
func BuildComponents(releaseDir string) {

	for itr := range buildEnv {
		fmt.Print("Building dms3 components for " + buildEnv[itr].envName + " platform... ")

		for jtr := range components {

			if components[jtr].compile {
				_, err := dms3libs.RunCommand("env " + buildEnv[itr].compileTags + " go build -o " + filepath.Join(releaseDir, buildEnv[itr].dirName) + "/" + components[jtr].exeName + " " + components[jtr].srcName)
				dms3libs.CheckErr(err)
			}
		}

		fmt.Println("Success")
	}
	fmt.Println()

}

// CopyServiceDaemons copies daemons into release folder
func CopyServiceDaemons(releaseDir string) {

	fmt.Print("Copying dms3 service daemons into " + releaseDir + " folder... ")
	dms3libs.CopyFile("dms3client/daemons/systemd/dms3client.service", filepath.Join(releaseDir, "dms3client/dms3client.service"))
	dms3libs.CopyFile("dms3server/daemons/systemd/dms3server.service", filepath.Join(releaseDir, "dms3server/dms3server.service"))
	fmt.Println("Success")
	fmt.Println()
}

// CopyMediaFiles copies dms3server media files into release folder
func CopyMediaFiles(releaseDir string) {

	fmt.Print("Copying dms3server media files (WAV) into " + releaseDir + " folder... ")
	dms3libs.CopyFile("dms3server/media/motion_start.wav", filepath.Join(releaseDir, "dms3server/media/motion_start.wav"))
	dms3libs.CopyFile("dms3server/media/motion_stop.wav", filepath.Join(releaseDir, "dms3server/media/motion_stop.wav"))
	fmt.Println("Success")
	fmt.Println()

}

// CopyConfigFiles copies dms3server media files into release folder
func CopyConfigFiles(releaseDir string) {

	fmt.Print("Copying dms3 component config files (TOML) into " + releaseDir + " folder... ")

	for itr := range components {
		dms3libs.CopyFile(components[itr].configFilename, filepath.Join(releaseDir, components[itr].dirName+"/"+components[itr].configFilename))
	}

	fmt.Println("Success")
	fmt.Println()

}

// CopyReleaseFolder copies the release folder into /etc/distributed-motion-s3
func CopyReleaseFolder(releaseDir string, confDir string) {

	fmt.Print("Copying " + releaseDir + " folder into " + confDir + " folder... ")
	dms3libs.RmDir(confDir)
	dms3libs.CopyDir(releaseDir, confDir)

	for itr := range buildEnv {
		dms3libs.RmDir(filepath.Join(confDir, buildEnv[itr].dirName))
	}

	fmt.Println("Success")
	fmt.Println()

}

// CopyComponents copies dms3 components into /usr/local/bin
func CopyComponents(releaseDir string, execDir string, platform string) {

	for itr := range buildEnv {

		if buildEnv[itr].envName == platform {
			fmt.Print("Copying " + buildEnv[itr].envName + " dms3 components into " + execDir + " folder (root permissions expected)... ")
			dms3libs.CopyDir(filepath.Join(releaseDir, buildEnv[itr].dirName), execDir)
			fmt.Println("Success")
			fmt.Println()
		}

	}

}

// RemoteMkDir creates a new folder over SSH with permissions passed in
func RemoteMkDir(ssh *easyssh.MakeConfig, newPath string) {

	_, _, _, err := ssh.Run("mkdir -p "+newPath, 5)
	dms3libs.CheckErr(err)

}

// RemoteCopyDir copies a directory over SSH from srcDir to destDir
func RemoteCopyDir(ssh *easyssh.MakeConfig, srcDir string, destDir string) {

	dirTree := dms3libs.WalkDir(srcDir)

	// create directory tree...
	for dirName, dirType := range dirTree {

		if dirType == 0 {
			RemoteMkDir(ssh, destDir+"/"+strings.TrimLeft(dirName, srcDir))
		}

	}

	// ...then copy files into directory tree
	for dirName, dirType := range dirTree {

		if dirType == 1 {
			ssh.Scp(dirName, dirName)
		}

	}

}
