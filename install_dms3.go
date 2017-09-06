package main

import (
	"fmt"
	"go-distributed-motion-s3/dms3build"
	"go-distributed-motion-s3/dms3libs"
	"os"
	"path/filepath"

	"github.com/mrgleam/easyssh"
)

// installs dms3 components (dms3client, dms3server, and dms3mail) and supporting configuration,
// service daemons, and media files onto the specified dms3 device component platforms (see
// installer_config.go for list of platforms to install onto)
//
// this installer depends on a local dms3_release folder created through dms3 compilation (see
// compile_dms3.go for details)
//
func main() {

	var ssh *easyssh.MakeConfig
	paths := make(map[string]string)

	// determine if running from the release folder or from source project
	if dms3build.IsRunningRelease() {
		base := filepath.Dir(filepath.Dir(dms3build.ExecFilePath()))
		paths["configFolder"] = filepath.Join(base, "dms3_release/dms3build")
		paths["releaseFolder"] = filepath.Join(base, "dms3_release")
	} else {
		base, _ := os.Getwd()
		paths["configFolder"] = filepath.Join(base, "config")
		paths["releaseFolder"] = filepath.Join(base, "dms3_release")
	}

	if !dms3libs.IsFile(paths["releaseFolder"]) {
		fmt.Println("No release folder found")
		os.Exit(1)
	}

	dms3libs.LoadComponentConfig(&dms3build.BuildConfig, paths["configFolder"]+"/dms3build.toml")

	// dms3client component installation
	//
	for _, client := range dms3build.BuildConfig.Clients {

		ssh = &easyssh.MakeConfig{
			User:     client.User,
			Server:   client.DeviceName,
			Password: client.SSHPassword,
			Port:     client.Port,
		}

		// copy dms3 release folder to remote device platform
		dms3build.RemoteCopyDir(ssh, paths["releaseFolder"], "dms3_release")
		dms3build.RemoteRunCommand(ssh, "chmod +x dms3_release/"+dms3build.BuildEnv[client.Platform].DirName+"/go_dms3client")
		dms3build.RemoteRunCommand(ssh, "chmod +x dms3_release/"+dms3build.BuildEnv[client.Platform].DirName+"/go_dms3mail")

		// copy client installer to remote device platform
		dms3build.RemoteCopyFile(ssh, paths["releaseFolder"]+"/"+dms3build.BuildEnv[client.Platform].DirName+"/dms3client_remote_installer", "dms3client_remote_installer")
		dms3build.RemoteRunCommand(ssh, "chmod +x dms3client_remote_installer")

		// run client installer, then remove on completion
		dms3build.RemoteRunCommand(ssh, "echo '"+client.RemoteAdminPassword+"' | sudo -S ./dms3client_remote_installer")
		dms3build.RemoteRunCommand(ssh, "rm dms3client_remote_installer")
		fmt.Println("")

	}

	// dms3server component installation
	//
	for _, server := range dms3build.BuildConfig.Servers {

		ssh = &easyssh.MakeConfig{
			User:     server.User,
			Server:   server.DeviceName,
			Password: server.SSHPassword,
			Port:     server.Port,
		}

		// copy dms3 release folder to remote device platform
		dms3build.RemoteCopyDir(ssh, paths["releaseFolder"], "dms3_release")
		dms3build.RemoteRunCommand(ssh, "chmod +x dms3_release/"+dms3build.BuildEnv[server.Platform].DirName+"/go_dms3server")

		// copy server installer to remote device platform
		dms3build.RemoteCopyFile(ssh, paths["releaseFolder"]+"/"+dms3build.BuildEnv[server.Platform].DirName+"/dms3server_remote_installer", "dms3server_remote_installer")
		dms3build.RemoteRunCommand(ssh, "chmod +x dms3server_remote_installer")

		// run server installer, then remove on completion
		dms3build.RemoteRunCommand(ssh, "echo '"+server.RemoteAdminPassword+"' | sudo -S ./dms3server_remote_installer")
		dms3build.RemoteRunCommand(ssh, "rm dms3server_remote_installer")
		fmt.Println("")

	}

}
