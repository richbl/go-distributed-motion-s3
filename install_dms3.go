package main

import (
	"fmt"
	"go-distributed-motion-s3/dms3_build"

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

		fmt.Print("Copying files to remote client ", dms3build.Clients[itr].User, "@", dms3build.Clients[itr].Server, "... ")
		dms3build.RemoteCopyDir(ssh, "dms3_release", "dms3_release")
		ssh.Run("chmod +x dms3_release/linux_arm7/go_dms3client", 5)
		ssh.Run("chmod +x dms3_release/linux_arm7/go_dms3mail", 5)
		fmt.Println("Success")

		fmt.Print("Copying installer script to remote client ", dms3build.Clients[itr].User, "@", dms3build.Clients[itr].Server, "... ")
		ssh.Scp("dms3_build/dms3client_remote_installer.sh", "dms3client_remote_installer.sh")
		ssh.Run("chmod +x dms3client_remote_installer.sh", 5)
		fmt.Println("Success")

		fmt.Print("Running installer script on remote client ", dms3build.Clients[itr].User, "@", dms3build.Clients[itr].Server, "... ")
		ssh.Run("sudo ./dms3client_remote_installer.sh", 5)
		fmt.Println("Success")

		fmt.Print("Removing installer script on remote client ", dms3build.Clients[itr].User, "@", dms3build.Clients[itr].Server, "... ")
		ssh.Run("rm dms3client_remote_installer.sh", 5)
		fmt.Println("Success")
		fmt.Println("")

	}

	for itr := range dms3build.Servers {

		ssh = &easyssh.MakeConfig{
			User:     dms3build.Servers[itr].User,
			Server:   dms3build.Servers[itr].Server,
			Password: dms3build.Servers[itr].Password,
			Port:     dms3build.Servers[itr].Port,
		}

		fmt.Print("Copying files to remote server ", dms3build.Servers[itr].User, "@", dms3build.Servers[itr].Server, "... ")
		dms3build.RemoteCopyDir(ssh, "dms3_release", "dms3_release")
		ssh.Run("chmod +x dms3_release/linux_amd64/go_dms3server", 5)
		fmt.Println("Success")

		fmt.Print("Copying installer script to remote server ", dms3build.Servers[itr].User, "@", dms3build.Servers[itr].Server, "... ")
		ssh.Scp("dms3_build/dms3server_remote_installer.sh", "dms3server_remote_installer.sh")
		ssh.Run("chmod +x dms3server_remote_installer.sh", 5)
		fmt.Println("Success")

		fmt.Print("Running installer script on remote server ", dms3build.Servers[itr].User, "@", dms3build.Servers[itr].Server, "... ")
		ssh.Run("./dms3server_remote_installer.sh", 5)
		fmt.Println("Success")

		fmt.Print("Removing installer script on remote server ", dms3build.Servers[itr].User, "@", dms3build.Servers[itr].Server, "... ")
		ssh.Run("rm dms3server_remote_installer.sh", 5)
		fmt.Println("Success")
		fmt.Println("")

	}

}
