package main

import (
	"fmt"
	"go-distributed-motion-s3/dms3build"
	"os"

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

	// dms3client component installation
	//
	for itr := range dms3build.Clients {

		ssh = &easyssh.MakeConfig{
			User:     dms3build.Clients[itr].User,
			Server:   dms3build.Clients[itr].Server,
			Password: dms3build.Clients[itr].SSHPassword,
			Port:     dms3build.Clients[itr].Port,
		}

		// copy dms3 release folder to remote device platform
		dms3build.RemoteCopyDir(ssh, "dms3_release", "dms3_release")
		dms3build.RemoteRunCommand(ssh, "chmod +x dms3_release/"+dms3build.BuildEnv[dms3build.Clients[itr].Platform].DirName+"/go_dms3client")
		dms3build.RemoteRunCommand(ssh, "chmod +x dms3_release/"+dms3build.BuildEnv[dms3build.Clients[itr].Platform].DirName+"/go_dms3mail")

		// copy client installer to remote device platform
		installerFile := dms3build.ReplaceFileContents("dms3build/dms3client_remote_installer.sh", map[string]string{
			"!PASSWORD": dms3build.Clients[itr].RemoteAdminPassword,
		})
		dms3build.RemoteCopyFile(ssh, installerFile.Name(), "dms3client_remote_installer.sh")
		dms3build.RemoteRunCommand(ssh, "chmod +x dms3client_remote_installer.sh")
		os.Remove(installerFile.Name())

		// run client installer, then remove on completion
		dms3build.RemoteRunCommand(ssh, "echo '"+dms3build.Clients[itr].RemoteAdminPassword+"' | sudo ./dms3client_remote_installer.sh")
		dms3build.RemoteRunCommand(ssh, "rm dms3client_remote_installer.sh")
		fmt.Println("")

	}

	// dms3server component installation
	//
	for itr := range dms3build.Servers {

		ssh = &easyssh.MakeConfig{
			User:     dms3build.Servers[itr].User,
			Server:   dms3build.Servers[itr].Server,
			Password: dms3build.Servers[itr].SSHPassword,
			Port:     dms3build.Servers[itr].Port,
		}

		// copy dms3 release folder to remote device platform
		dms3build.RemoteCopyDir(ssh, "dms3_release", "dms3_release")
		dms3build.RemoteRunCommand(ssh, "chmod +x dms3_release/"+dms3build.BuildEnv[dms3build.Servers[itr].Platform].DirName+"/go_dms3server")

		// copy server installer to remote device platform
		installerFile := dms3build.ReplaceFileContents("dms3build/dms3server_remote_installer.sh", map[string]string{
			"!PASSWORD": dms3build.Servers[itr].RemoteAdminPassword,
		})
		dms3build.RemoteCopyFile(ssh, installerFile.Name(), "dms3server_remote_installer.sh")
		dms3build.RemoteRunCommand(ssh, "chmod +x dms3server_remote_installer.sh")
		os.Remove(installerFile.Name())

		// run server installer, then remove on completion
		dms3build.RemoteRunCommand(ssh, "echo '"+dms3build.Servers[itr].RemoteAdminPassword+"' | sudo -S ./dms3server_remote_installer.sh")
		dms3build.RemoteRunCommand(ssh, "rm dms3server_remote_installer.sh")
		fmt.Println("")

	}

}
