package main

import (
	"fmt"
	"go-distributed-motion-s3/dms3build"
	"go-distributed-motion-s3/dms3libs"
	"strconv"
)

// this script will be copied to the dms3 device component platform, executed, and
// then deleted automatically
//
// NOTE: must be run with admin privileges

func main() {

	// get platform type passed in as arg
	args := dms3libs.GetArgs()
	val, err := strconv.Atoi(args[1])
	dms3libs.CheckErr(err)

	fmt.Print("Stopping systemd dms3client.service...")
	dms3libs.RunCommand("systemctl stop dms3client.service")
	fmt.Println("Success")
	fmt.Println()

	fmt.Print("Moving files into /usr/local/bin... ")
	dms3libs.CopyFile("dms3_release/"+dms3build.BuildEnv[val].DirName+"/go_dms3client", "/usr/local/bin/go_dms3client")
	_, err = dms3libs.RunCommand("chmod +x " + "/usr/local/bin/go_dms3client")
	dms3libs.CheckErr(err)
	dms3libs.CopyFile("dms3_release/"+dms3build.BuildEnv[val].DirName+"/go_dms3mail", "/usr/local/bin/go_dms3mail")
	_, err = dms3libs.RunCommand("chmod +x " + "/usr/local/bin/go_dms3mail")
	dms3libs.CheckErr(err)
	fmt.Println("Success")
	fmt.Println()

	fmt.Print("Copying files into /etc/distributed-motion-s3... ")
	dms3libs.MkDir("/etc/distributed-motion-s3")
	dms3libs.CopyDir("dms3_release/dms3client", "/etc/distributed-motion-s3")
	dms3libs.CopyDir("dms3_release/dms3libs", "/etc/distributed-motion-s3")
	dms3libs.CopyDir("dms3_release/dms3mail", "/etc/distributed-motion-s3")
	dms3libs.RmDir("dms3_release")
	fmt.Println("Success")
	fmt.Println()

	fmt.Print("Restarting systemd dms3client.service...")
	dms3libs.RunCommand("systemctl start dms3client.service")
	fmt.Println("Success")
	fmt.Println()

}
