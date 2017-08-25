package main

import (
	"fmt"
	"go-distributed-motion-s3/dms3build"

	"github.com/hypersleep/easyssh"
)

func main() {

	var ssh *easyssh.MakeConfig

	for itr := range dms3build.Clients {

		ssh = &easyssh.MakeConfig{
			User:     dms3build.Clients[itr].User,
			Server:   dms3build.Clients[itr].Server,
			Password: dms3build.Clients[itr].Password,
			Port:     dms3build.Clients[itr].Port,
		}

		// copy dms3 release folder to remote device platform
		dms3build.RemoteCopyDir(ssh, "dms3_release", "dms3_release")
		dms3build.RemoteRunCommand(ssh, "chmod +x dms3_release/linux_arm7/go_dms3client")
		dms3build.RemoteRunCommand(ssh, "chmod +x dms3_release/linux_arm7/go_dms3mail")

		// copy client installer to remote device platform
		dms3build.RemoteCopyFile(ssh, "dms3build/dms3client_remote_installer.sh", "dms3client_remote_installer.sh")
		dms3build.RemoteRunCommand(ssh, "chmod +x dms3client_remote_installer.sh")

		// run client installer, then remove on completion
		dms3build.RemoteRunCommand(ssh, "sudo ./dms3client_remote_installer.sh")
		dms3build.RemoteRunCommand(ssh, "rm dms3client_remote_installer.sh")
		fmt.Println("")

	}

	for itr := range dms3build.Servers {

		ssh = &easyssh.MakeConfig{
			User:     dms3build.Servers[itr].User,
			Server:   dms3build.Servers[itr].Server,
			Password: dms3build.Servers[itr].Password,
			Port:     dms3build.Servers[itr].Port,
		}

		// copy dms3 release folder to remote device platform
		dms3build.RemoteCopyDir(ssh, "dms3_release", "dms3_release")
		dms3build.RemoteRunCommand(ssh, "chmod +x dms3_release/linux_amd64/go_dms3server")

		// copy server installer to remote device platform
		dms3build.RemoteCopyFile(ssh, "dms3build/dms3server_remote_installer.sh", "dms3server_remote_installer.sh")
		dms3build.RemoteRunCommand(ssh, "chmod +x dms3server_remote_installer.sh")

		// run server installer, then remove on completion
		dms3build.RemoteRunCommand(ssh, "sudo ./dms3server_remote_installer.sh")
		dms3build.RemoteRunCommand(ssh, "rm dms3server_remote_installer.sh")
		fmt.Println("")

	}

}
