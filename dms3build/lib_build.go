// Package dms3build library
//
package dms3build

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/mrgleam/easyssh"
	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

// BuildReleaseFolder creates the directory structure for each platform passed into it
func BuildReleaseFolder() {

	dms3libs.RmDir("dms3_release")

	for platformType := range BuildEnv {
		fmt.Print("Creating release folder for " + BuildEnv[platformType].dirName + " platform... ")
		dms3libs.MkDir(filepath.Join("dms3_release", BuildEnv[platformType].dirName))
		fmt.Println("Success")
	}

	for itr := range components {
		fmt.Print("Creating release folder for " + components[itr].exeName + " component... ")
		dirName := components[itr].dirName

		if components[itr].dirName == "dms3server" {
			dirName = filepath.Join(dirName, "media")
		}

		dms3libs.MkDir(filepath.Join("dms3_release", dirName))
		fmt.Println("Success")
	}

	fmt.Println()

}

// BuildComponents compiles dms3 components for each platform passed into it
func BuildComponents() {

	for platformType := range BuildEnv {
		fmt.Print("Building dms3 components for " + BuildEnv[platformType].dirName + " platform... ")

		for jtr := range components {

			if components[jtr].compile {
				_, err := dms3libs.RunCommand("env " + BuildEnv[platformType].compileTags + " go build -o " + filepath.Join("dms3_release", BuildEnv[platformType].dirName) + "/" + components[jtr].exeName + " " + components[jtr].srcName)
				dms3libs.CheckErr(err)
			}

		}

		fmt.Println("Success")
	}
	fmt.Println()

}

// CopyServiceDaemons copies daemons into release folder
func CopyServiceDaemons() {

	fmt.Print("Copying dms3 service daemons into dms3_release folder... ")
	dms3libs.CopyFile("dms3client/daemons/systemd/dms3client.service", filepath.Join("dms3_release", "dms3client/dms3client.service"))
	dms3libs.CopyFile("dms3server/daemons/systemd/dms3server.service", filepath.Join("dms3_release", "dms3server/dms3server.service"))
	fmt.Println("Success")
	fmt.Println()
}

// CopyMediaFiles copies dms3server media files into release folder
func CopyMediaFiles() {

	fmt.Print("Copying dms3server media files (WAV) into dms3_release folder... ")
	dms3libs.CopyFile("dms3server/media/motion_start.wav", filepath.Join("dms3_release", "dms3server/media/motion_start.wav"))
	dms3libs.CopyFile("dms3server/media/motion_stop.wav", filepath.Join("dms3_release", "dms3server/media/motion_stop.wav"))
	fmt.Println("Success")
	fmt.Println()

}

// CopyDashboardFiles copies the dms3dashboard html file into release folder
func CopyDashboardFiles() {

	fmt.Print("Copying dms3dashboard file (HTML) into dms3_release folder... ")
	dms3libs.CopyFile("dms3dashboard/dashboard.html", filepath.Join("dms3_release", "dms3dashboard/dashboard.html"))
	fmt.Println("Success")

	fmt.Print("Copying dms3dashboard assets into dms3_release folder... ")
	dms3libs.CopyDir("dms3dashboard/assets", filepath.Join("dms3_release", "dms3dashboard"))
	fmt.Println("Success")
	fmt.Println()

}

// CopyConfigFiles copies dms3server media files into release folder
func CopyConfigFiles() {

	fmt.Print("Copying dms3 component config files (TOML) into dms3_release folder... ")

	for itr := range components {

		if components[itr].configFilename != "" {
			dms3libs.CopyFile(filepath.Join("config", components[itr].configFilename), filepath.Join("dms3_release", components[itr].dirName+"/"+components[itr].configFilename))
		}

	}

	fmt.Println("Success")
	fmt.Println()

}

// ReleasePath sets the installer release path based on whether called from source project or
// from a binary release folder
//
func ReleasePath() map[string]string {

	paths := make(map[string]string)

	if isRunningRelease() {
		base := filepath.Dir(filepath.Dir(execFilePath()))
		paths["configFolder"] = filepath.Join(base, "dms3_release/dms3build")
		paths["releaseFolder"] = filepath.Join(base, "dms3_release")
	} else {
		base, _ := os.Getwd()
		paths["configFolder"] = filepath.Join(base, "config")
		paths["releaseFolder"] = filepath.Join(base, "dms3_release")
	}

	return paths

}

// ConfirmReleaseFolder checks for the existence of the release folder
func ConfirmReleaseFolder(releasePath string) {

	if !dms3libs.IsFile(releasePath) {
		dms3libs.CheckErr(errors.New("No release folder found"))
	}

}

// InstallClientComponents installs dms3client components onto device platforms identified in
// the dms3build.toml configuration file
//
func InstallClientComponents(releasePath string) {

	var ssh *easyssh.MakeConfig

	for _, client := range BuildConfig.Clients {

		ssh = &easyssh.MakeConfig{
			User:     client.User,
			Server:   client.DeviceName,
			Password: client.SSHPassword,
			Port:     strconv.Itoa(client.Port),
		}

		// copy dms3 release folder components to remote device platform
		remoteMkDir(ssh, "dms3_release")

		remoteCopyDir(ssh, filepath.Join(releasePath, "dms3client"), filepath.Join("dms3_release", "dms3client"))
		remoteCopyDir(ssh, filepath.Join(releasePath, "dms3libs"), filepath.Join("dms3_release", "dms3libs"))
		remoteCopyDir(ssh, filepath.Join(releasePath, "dms3mail"), filepath.Join("dms3_release", "dms3mail"))

		remoteMkDir(ssh, filepath.Join("dms3_release", "dms3dashboard"))
		remoteCopyFile(ssh, filepath.Join(filepath.Join(releasePath, "dms3dashboard"), "dms3dashboard.toml"), filepath.Join(filepath.Join("dms3_release", "dms3dashboard"), "dms3dashboard.toml"))

		remoteCopyDir(ssh, filepath.Join(filepath.Join(releasePath, BuildEnv[client.Platform].dirName), "dms3client"), filepath.Join("dms3_release", "dms3client"))
		remoteCopyDir(ssh, filepath.Join(filepath.Join(releasePath, BuildEnv[client.Platform].dirName), "dms3mail"), filepath.Join("dms3_release", "dms3mail"))
		remoteCopyDir(ssh, filepath.Join(filepath.Join(releasePath, BuildEnv[client.Platform].dirName), "dms3client_remote_installer"), "dms3client_remote_installer")
		remoteRunCommand(ssh, "chmod +x dms3client_remote_installer")

		// run client installer, then remove on completion
		remoteRunCommand(ssh, "echo '"+client.RemoteAdminPassword+"' | sudo -S ./dms3client_remote_installer")
		remoteRunCommand(ssh, "rm dms3client_remote_installer")
		fmt.Println("")

	}

}

// InstallServerComponents installs dms3server components onto device platforms identified in
// the dms3build.toml configuration file
//
func InstallServerComponents(releasePath string) {

	var ssh *easyssh.MakeConfig

	for _, server := range BuildConfig.Servers {

		ssh = &easyssh.MakeConfig{
			User:     server.User,
			Server:   server.DeviceName,
			Password: server.SSHPassword,
			Port:     strconv.Itoa(server.Port),
		}

		// copy dms3 release folder components to remote device platform
		remoteMkDir(ssh, "dms3_release")

		remoteCopyDir(ssh, filepath.Join(releasePath, "dms3server"), filepath.Join("dms3_release", "dms3server"))
		remoteCopyDir(ssh, filepath.Join(releasePath, "dms3libs"), filepath.Join("dms3_release", "dms3libs"))
		remoteCopyDir(ssh, filepath.Join(releasePath, "dms3dashboard"), filepath.Join("dms3_release", "dms3dashboard"))

		remoteCopyDir(ssh, filepath.Join(filepath.Join(releasePath, BuildEnv[server.Platform].dirName), "dms3server"), filepath.Join("dms3_release", "dms3server"))
		remoteCopyDir(ssh, filepath.Join(filepath.Join(releasePath, BuildEnv[server.Platform].dirName), "dms3server_remote_installer"), "dms3server_remote_installer")
		remoteRunCommand(ssh, "chmod +x dms3server_remote_installer")

		// run server installer, then remove on completion
		remoteRunCommand(ssh, "echo '"+server.RemoteAdminPassword+"' | sudo -S ./dms3server_remote_installer")
		remoteRunCommand(ssh, "rm dms3server_remote_installer")
		fmt.Println("")

	}

}

// remoteMkDir creates a new folder over SSH with permissions passed in
func remoteMkDir(ssh *easyssh.MakeConfig, newPath string) {

	_, _, _, err := ssh.Run("mkdir -p "+newPath, 5)
	dms3libs.CheckErr(err)

}

// remoteCopyDir copies a directory over SSH from srcDir to destDir
func remoteCopyDir(ssh *easyssh.MakeConfig, srcDir string, destDir string) {

	fmt.Print("Copying folder " + srcDir + " to " + ssh.User + "@" + ssh.Server + ":" + destDir + "... ")
	dirTree := dms3libs.WalkDir(srcDir)

	// create directory tree...
	for dirName, dirType := range dirTree {

		if dirType == 0 {
			remoteMkDir(ssh, destDir+dirName[len(srcDir):])
		}

	}

	// ...then copy files into directory tree
	for dirName, dirType := range dirTree {

		if dirType == 1 {
			remoteCopyFile(ssh, dirName, destDir+dirName[len(srcDir):])
		}

	}

	fmt.Println("Success")

}

// remoteRunCommand runs a command via the SSH protocol
func remoteRunCommand(ssh *easyssh.MakeConfig, command string) {

	fmt.Print("Running command " + "'" + command + "' on " + ssh.User + "@" + ssh.Server + "... ")
	_, _, _, err := ssh.Run(command, 5)
	dms3libs.CheckErr(err)
	fmt.Println("Success")

}

// remoteCopyFile copies a file from src to a remote dest using SCP
func remoteCopyFile(ssh *easyssh.MakeConfig, srcFile string, destFile string) {

	fmt.Print("Copying file " + srcFile + " to " + destFile + " on " + ssh.User + "@" + ssh.Server + "... ")
	err := ssh.Scp(srcFile, destFile)
	dms3libs.CheckErr(err)
	fmt.Println("Success")

}

// execFilePath returns the absolute path to the project root (default: go-distributed-motion-s3)
func execFilePath() string {

	ex, _ := os.Executable()
	return filepath.Dir(ex)

}

// isRunningRelease checks if the executable was called from the dms3_release folder
func isRunningRelease() bool {

	dir, _ := filepath.Abs(execFilePath())
	return filepath.Base(filepath.Dir(dir)) == "dms3_release"

}
