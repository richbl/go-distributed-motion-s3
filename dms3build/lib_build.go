package dms3build

import (
	"fmt"
	"go-distributed-motion-s3/dms3libs"
	"os"
	"path/filepath"

	"github.com/mrgleam/easyssh"
)

// BuildReleaseFolder creates the directory structure for each platform passed into it
func BuildReleaseFolder(releaseDir string) {

	dms3libs.RmDir(releaseDir)

	for itr := range BuildEnv {
		fmt.Print("Creating release folder for " + BuildEnv[itr].envName + " platform... ")
		dms3libs.MkDir(filepath.Join(releaseDir, BuildEnv[itr].DirName))
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

	for itr := range BuildEnv {
		fmt.Print("Building dms3 components for " + BuildEnv[itr].envName + " platform... ")

		for jtr := range components {

			if components[jtr].compile {
				_, err := dms3libs.RunCommand("env " + BuildEnv[itr].compileTags + " go build -o " + filepath.Join(releaseDir, BuildEnv[itr].DirName) + "/" + components[jtr].exeName + " " + components[jtr].srcName)
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

// CopyInstallerFiles copies dms3s remote installer scripts into release folder
func CopyInstallerFiles(releaseDir string) {

	fmt.Print("Copying dms3 remote installer scripts into " + releaseDir + " folder... ")
	dms3libs.CopyFile("dms3build/dms3client_remote_installer.sh", filepath.Join(releaseDir, "dms3build/dms3client_remote_installer.sh"))
	dms3libs.CopyFile("dms3build/dms3server_remote_installer.sh", filepath.Join(releaseDir, "dms3build/dms3server_remote_installer.sh"))
	fmt.Println("Success")
	fmt.Println()

}

// CopyConfigFiles copies dms3server media files into release folder
func CopyConfigFiles(releaseDir string) {

	fmt.Print("Copying dms3 component config files (TOML) into " + releaseDir + " folder... ")

	for itr := range components {
		dms3libs.CopyFile(filepath.Join("config", components[itr].configFilename), filepath.Join(releaseDir, components[itr].dirName+"/"+components[itr].configFilename))
	}

	fmt.Println("Success")
	fmt.Println()

}

// RemoteMkDir creates a new folder over SSH with permissions passed in
func RemoteMkDir(ssh *easyssh.MakeConfig, newPath string) {

	_, _, _, err := ssh.Run("mkdir -p "+newPath, 5)
	dms3libs.CheckErr(err)

}

// RemoteCopyDir copies a directory over SSH from srcDir to destDir
func RemoteCopyDir(ssh *easyssh.MakeConfig, srcDir string, destDir string) {

	fmt.Print("Copying folder " + srcDir + " to " + ssh.User + "@" + ssh.Server + ":" + destDir + "... ")
	dirTree := dms3libs.WalkDir(srcDir)

	// create directory tree...
	for dirName, dirType := range dirTree {

		if dirType == 0 {
			RemoteMkDir(ssh, destDir+dirName[len(srcDir):])
		}

	}

	// ...then copy files into directory tree
	for dirName, dirType := range dirTree {

		if dirType == 1 {
			ssh.Scp(dirName, destDir+dirName[len(srcDir):])
		}

	}

	fmt.Println("Success")

}

// RemoteRunCommand runs a command via the SSH protocol
func RemoteRunCommand(ssh *easyssh.MakeConfig, command string) {

	fmt.Print("Running command " + "'" + command + "' on " + ssh.User + "@" + ssh.Server + "... ")
	_, _, _, err := ssh.Run(command, 5)
	dms3libs.CheckErr(err)
	fmt.Println("Success")

}

// RemoteCopyFile copies a file from src to a remote dest using SCP
func RemoteCopyFile(ssh *easyssh.MakeConfig, srcFile string, destFile string) {

	fmt.Print("Copying file " + srcFile + " to " + destFile + " on " + ssh.User + "@" + ssh.Server + "... ")
	err := ssh.Scp(srcFile, destFile)
	dms3libs.CheckErr(err)
	fmt.Println("Success")

}

// ExecFilePath returns the absolute path to the project root (default: go-distributed-motion-s3)
func ExecFilePath() string {

	ex, _ := os.Executable()
	return filepath.Dir(ex)

}

// IsRunningRelease checks if the installer is running against the project source or release folder
func IsRunningRelease() bool {

	dir, _ := filepath.Abs(ExecFilePath())

	if filepath.Base(filepath.Dir(dir)) == "dms3_release" {
		return true
	}
	return false

}
