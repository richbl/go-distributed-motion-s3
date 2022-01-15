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
//
func BuildReleaseFolder() {

	dms3libs.RmDir("dms3_release")

	for platformType := range BuildEnv {
		fmt.Println("Creating release folder for " + BuildEnv[platformType].dirName + " platform...")
		dms3libs.MkDir(filepath.Join("dms3_release", "cmd", BuildEnv[platformType].dirName))
	}

	for itr := range components {
		fmt.Println("Creating release folder for " + components[itr].exeName + " component...")
		dirName := components[itr].dirName

		if components[itr].dirName == "dms3server" {
			dirName = filepath.Join(dirName, "media")
		}

		dms3libs.MkDir(filepath.Join("dms3_release", "config", dirName))
	}

}

// BuildComponents compiles dms3 components for each platform passed into it
//
func BuildComponents() {

	for platformType := range BuildEnv {
		fmt.Println("Building dms3 components for " + BuildEnv[platformType].dirName + " platform...")

		for jtr := range components {

			if components[jtr].compile {

				// dms3 installer (install_dms3) is only built once (natively on linuxAMD64) into dms3_release root
				if components[jtr].exeName == "install_dms3" {

					if platformType == linuxAMD64 {
						_, err := dms3libs.RunCommand(dms3libs.LibConfig.SysCommands["ENV"] + " " + BuildEnv[platformType].compileTags + " go build -o " + filepath.Join("dms3_release", "cmd", components[jtr].exeName) + " " + components[jtr].srcName)
						dms3libs.CheckErr(err)
					}

				} else {
					_, err := dms3libs.RunCommand(dms3libs.LibConfig.SysCommands["ENV"] + " " + BuildEnv[platformType].compileTags + " go build -o " + filepath.Join("dms3_release", "cmd", BuildEnv[platformType].dirName, components[jtr].exeName) + " " + components[jtr].srcName)
					dms3libs.CheckErr(err)
				}

			}

		}

	}

}

// CopyServiceDaemons copies daemons into release folder
//
func CopyServiceDaemons() {

	fmt.Println("Copying dms3 service daemons into dms3_release folder...")

	dms3libs.CopyFile(filepath.Join("dms3client", "daemons", "systemd", "dms3client.service"), filepath.Join("dms3_release", "config", "dms3client", "dms3client.service"))
	dms3libs.CopyFile(filepath.Join("dms3server", "daemons", "systemd", "dms3server.service"), filepath.Join("dms3_release", "config", "dms3server", "dms3server.service"))

}

// CopyMediaFiles copies dms3server media files into release folder
//
func CopyMediaFiles() {

	fmt.Println("Copying dms3server media files (WAV) into dms3_release folder...")

	dms3libs.CopyFile(filepath.Join("dms3server", "media", "motion_start.wav"), filepath.Join("dms3_release", "config", "dms3server", "media", "motion_start.wav"))
	dms3libs.CopyFile(filepath.Join("dms3server", "media", "motion_stop.wav"), filepath.Join("dms3_release", "config", "dms3server", "media", "motion_stop.wav"))

}

// CopyComponents copies component html files and assets into the release folder
func CopyComponents(component string) {

	fmt.Println("Copying " + component + " file (HTML) into dms3_release folder...")
	dms3libs.CopyFile(filepath.Join(component, component+".html"), filepath.Join("dms3_release", "config", component, component+".html"))

	fmt.Println("Copying " + component + " assets into dms3_release folder...")
	dms3libs.CopyDir(filepath.Join(component, "assets"), filepath.Join("dms3_release", "config", component))

}

// CopyConfigFiles copies config files into release folder
//
func CopyConfigFiles() {

	fmt.Println("Copying dms3 component config files (TOML) into dms3_release folder...")

	for itr := range components {

		if components[itr].configFilename != "" {
			dms3libs.CopyFile(filepath.Join("config", components[itr].configFilename), filepath.Join("dms3_release", "config", components[itr].dirName, components[itr].configFilename))
		}

	}

}

// ConfirmReleaseFolder checks for the existence of the release folder
//
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

	fmt.Println("Installing dms3client components onto remote device(s) identified in dms3build.toml...")

	for _, client := range BuildConfig.Clients {

		ssh = &easyssh.MakeConfig{
			User:     client.User,
			Server:   client.DeviceName,
			Password: client.SSHPassword,
			Port:     strconv.Itoa(client.Port),
		}

		// copy dms3 release folder components to remote device platform
		remoteMkDir(ssh, "dms3_release")
		remoteCopyDir(ssh, filepath.Join(releasePath, "config", "dms3client"), filepath.Join("dms3_release", "config", "dms3client"))
		remoteCopyDir(ssh, filepath.Join(releasePath, "config", "dms3libs"), filepath.Join("dms3_release", "config", "dms3libs"))
		remoteCopyDir(ssh, filepath.Join(releasePath, "config", "dms3mail"), filepath.Join("dms3_release", "config", "dms3mail"))
		remoteCopyDir(ssh, filepath.Join(releasePath, "config", "dms3dashboard"), filepath.Join("dms3_release", "config", "dms3dashboard"))

		// copy dms3 release file components to remote device platform
		remoteMkDir(ssh, filepath.Join("dms3_release", "cmd"))
		remoteCopyFile(ssh, filepath.Join(releasePath, "cmd", BuildEnv[client.Platform].dirName, "dms3client"), filepath.Join("dms3_release", "cmd", "dms3client"))
		remoteCopyFile(ssh, filepath.Join(releasePath, "cmd", BuildEnv[client.Platform].dirName, "dms3mail"), filepath.Join("dms3_release", "cmd", "dms3mail"))
		remoteCopyFile(ssh, filepath.Join(releasePath, "cmd", BuildEnv[client.Platform].dirName, "dms3client_remote_installer"), "dms3client_remote_installer")

		// run client installer, then remove on completion
		remoteRunCommand(ssh, "echo "+client.RemoteAdminPassword+" | sudo -S "+"."+string(filepath.Separator)+"dms3client_remote_installer")
		remoteRunCommand(ssh, "rm dms3client_remote_installer")

	}

}

// InstallServerComponents installs dms3server components onto device platforms identified in
// the dms3build.toml configuration file
//
func InstallServerComponents(releasePath string) {

	var ssh *easyssh.MakeConfig

	fmt.Println("Installing dms3server components onto remote device(s) identified in dms3build.toml...")

	for _, server := range BuildConfig.Servers {

		ssh = &easyssh.MakeConfig{
			User:     server.User,
			Server:   server.DeviceName,
			Password: server.SSHPassword,
			Port:     strconv.Itoa(server.Port),
		}

		// copy dms3 release folder components to remote device platform
		remoteMkDir(ssh, "dms3_release")
		remoteCopyDir(ssh, filepath.Join(releasePath, "config", "dms3server"), filepath.Join("dms3_release", "config", "dms3server"))
		remoteCopyDir(ssh, filepath.Join(releasePath, "config", "dms3libs"), filepath.Join("dms3_release", "config", "dms3libs"))
		remoteCopyDir(ssh, filepath.Join(releasePath, "config", "dms3dashboard"), filepath.Join("dms3_release", "config", "dms3dashboard"))

		remoteMkDir(ssh, filepath.Join("dms3_release", "cmd"))
		remoteCopyFile(ssh, filepath.Join(filepath.Join(releasePath, "cmd", BuildEnv[server.Platform].dirName), "dms3server"), filepath.Join("dms3_release", "cmd", "dms3server"))
		remoteCopyFile(ssh, filepath.Join(filepath.Join(releasePath, "cmd", BuildEnv[server.Platform].dirName), "dms3server_remote_installer"), "dms3server_remote_installer")

		// run server installer, then remove on completion
		remoteRunCommand(ssh, "echo "+server.RemoteAdminPassword+" | sudo -S "+"."+string(filepath.Separator)+"dms3server_remote_installer")
		remoteRunCommand(ssh, "rm dms3server_remote_installer")
	}

}

// remoteMkDir creates a new folder over SSH with permissions passed in
//
func remoteMkDir(ssh *easyssh.MakeConfig, newPath string) {
	remoteRunCommand(ssh, "mkdir -p "+newPath)
}

// remoteCopyDir copies a directory over SSH from srcDir to destDir
//
func remoteCopyDir(ssh *easyssh.MakeConfig, srcDir string, destDir string) {

	fmt.Println("Copying folder " + srcDir + " to " + ssh.User + "@" + ssh.Server + ":" + destDir + "...")

	dirTree := dms3libs.WalkDir(srcDir)

	// create directory tree...
	for dirName, dirType := range dirTree {

		if dirType == 0 {
			remoteMkDir(ssh, destDir+dirName[len(srcDir):])
		}

	}

	// ...then copy files into directory tree
	for fileName, dirType := range dirTree {

		if dirType == 1 {
			remoteCopyFile(ssh, fileName, destDir+fileName[len(srcDir):])
		}

	}

}

// remoteRunCommand runs a command via the SSH protocol
//
func remoteRunCommand(ssh *easyssh.MakeConfig, command string) {

	fmt.Println("Running remote command " + "'" + command + "' on " + ssh.User + "@" + ssh.Server + "...")

	_, _, _, err := ssh.Run(command, 5)
	dms3libs.CheckErr(err)

}

// remoteCopyFile copies a file from src to a remote dest file using SCP
//
func remoteCopyFile(ssh *easyssh.MakeConfig, srcFile string, destFile string) {

	fmt.Println("Copying file " + srcFile + " to " + ssh.User + "@" + ssh.Server + ":" + destFile + "...")

	srcAttrib, err := os.Stat(srcFile)
	dms3libs.CheckErr(err)

	err = ssh.Scp(srcFile, destFile)
	dms3libs.CheckErr(err)

	// ssh.Scp() does not set file attribs, so reset them on remote device
	remoteRunCommand(ssh, "chmod 0"+strconv.FormatUint(uint64(srcAttrib.Mode()), 8)+" "+destFile)

}
