package main

import (
	"fmt"
	"go-distributed-motion-s3/dms3libs"
	"path/filepath"
)

func main() {

	const (
		execDir    = "/usr/local/bin"             // set to anywhere seen by $PATH
		confDir    = "/etc/distributed-motion-s3" // default location to store config files (*.toml)
		releaseDir = "dms3_release"               // release folder containing all platform builds
		linuxArm7  = "linux_arm7"                 // Linux Arm7 platform
		linuxAMD64 = "linux_amd64"                // Linux AMD64 platform
	)

	// build platform-specific components into release folder
	//
	dms3libs.RmDir(releaseDir)
	dms3libs.MkDir(filepath.Join(releaseDir, linuxArm7))
	dms3libs.MkDir(filepath.Join(releaseDir, linuxAMD64))
	dms3libs.MkDir(filepath.Join(releaseDir, "dms3client"))
	dms3libs.MkDir(filepath.Join(releaseDir, "dms3server/media"))
	dms3libs.MkDir(filepath.Join(releaseDir, "dms3libs"))
	dms3libs.MkDir(filepath.Join(releaseDir, "dms3mail"))

	BuildComponents(releaseDir, linuxAMD64)
	BuildComponents(releaseDir, linuxArm7)

	// fmt.Print("building dms3 components for Linux/ARM7 platform...")
	// _, err := dms3libs.RunCommand("env GOOS=linux GOARCH=arm GOARM=7 go build -o " + filepath.Join(releaseDir, linuxArm7) + "/go_dms3client go_dms3client.go")
	// dms3libs.CheckErr(err)
	// _, err = dms3libs.RunCommand("env GOOS=linux GOARCH=arm GOARM=7 go build -o " + filepath.Join(releaseDir, linuxArm7) + "/go_dms3server go_dms3server.go")
	// dms3libs.CheckErr(err)
	// _, err = dms3libs.RunCommand("env GOOS=linux GOARCH=arm GOARM=7 go build -o " + filepath.Join(releaseDir, linuxArm7) + "/go_dms3mail go_dms3mail.go")
	// dms3libs.CheckErr(err)
	// fmt.Println("success")

	// fmt.Print("building dms3 components for Linux/AMD64 platform...")
	// _, err = dms3libs.RunCommand("env GOOS=linux GOARCH=amd64 go build -o " + filepath.Join(releaseDir, linuxAMD64) + "/go_dms3client go_dms3client.go")
	// dms3libs.CheckErr(err)
	// _, err = dms3libs.RunCommand("env GOOS=linux GOARCH=amd64 go build -o " + filepath.Join(releaseDir, linuxAMD64) + "/go_dms3server go_dms3server.go")
	// dms3libs.CheckErr(err)
	// _, err = dms3libs.RunCommand("env GOOS=linux GOARCH=amd64 go build -o " + filepath.Join(releaseDir, linuxAMD64) + "/go_dms3mail go_dms3mail.go")
	// dms3libs.CheckErr(err)
	// fmt.Println("success")

	// copy service daemons into release folder
	//
	fmt.Print("copying service daemons into " + releaseDir + " folder...")
	dms3libs.CopyFile("dms3client/daemons/systemd/dms3client.service", filepath.Join(releaseDir, "dms3client/dms3client.service"))
	dms3libs.CopyFile("dms3server/daemons/systemd/dms3server.service", filepath.Join(releaseDir, "dms3server/dms3server.service"))
	fmt.Println("success")

	// copy dms3server media files into release folder
	//
	fmt.Print("copying dms3server media files (WAV) into " + releaseDir + " folder...")
	dms3libs.CopyFile("dms3server/media/motion_start.wav", filepath.Join(releaseDir, "dms3server/media/motion_start.wav"))
	dms3libs.CopyFile("dms3server/media/motion_stop.wav", filepath.Join(releaseDir, "dms3server/media/motion_stop.wav"))
	fmt.Println("success")

	// copy TOML files into release folder
	//
	fmt.Print("copying dms3 component config files (TOML) into " + releaseDir + " folder...")
	dms3libs.CopyFile("dms3client.toml", filepath.Join(releaseDir, "dms3client/dms3client.toml"))
	dms3libs.CopyFile("dms3server.toml", filepath.Join(releaseDir, "dms3server/dms3server.toml"))
	dms3libs.CopyFile("dms3libs.toml", filepath.Join(releaseDir, "dms3libs/dms3libs.toml"))
	dms3libs.CopyFile("dms3mail.toml", filepath.Join(releaseDir, "dms3mail/dms3mail.toml"))
	fmt.Println("success")

	// copy release folder into /etc/distributed-motion-s3
	//
	fmt.Print("copying " + releaseDir + " folder into " + confDir + " folder...")
	dms3libs.RmDir(confDir)
	dms3libs.CopyDir(releaseDir, confDir)
	dms3libs.RmDir(filepath.Join(confDir, linuxArm7))
	dms3libs.RmDir(filepath.Join(confDir, linuxAMD64))
	fmt.Println("success")

	// copy dms3 components into /usr/local/bin
	//
	fmt.Print("copying Linux/AMD64 dms3 components into " + execDir + " folder (root permissions expected)...")
	dms3libs.CopyDir(filepath.Join(releaseDir, linuxAMD64), execDir)
	fmt.Println("success")

	// fmt.Println("copying Linux/Arm7 dms3 components into " + execDir + " folder (root permissions expected)...")
	// dms3libs.CopyDir(filepath.Join(releaseDir, linuxArm7), execDir)
	// fmt.Println("success")
}
