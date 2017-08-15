package main

import (
	"fmt"
	"go-distributed-motion-s3/dms3libs"
	"path/filepath"
)

// BuildReleaseFolder comment
// func BuildReleaseFolder() {

// 	dms3libs.RmDir(releaseDir)
// 	dms3libs.MkDir(filepath.Join(releaseDir, linuxArm7))
// 	dms3libs.MkDir(filepath.Join(releaseDir, linuxAMD64))
// 	dms3libs.MkDir(filepath.Join(releaseDir, "dms3client"))
// 	dms3libs.MkDir(filepath.Join(releaseDir, "dms3server/media"))
// 	dms3libs.MkDir(filepath.Join(releaseDir, "dms3libs"))
// 	dms3libs.MkDir(filepath.Join(releaseDir, "dms3mail"))

// }

// BuildComponents comment
func BuildComponents(releaseDir string, platform string) {

	fmt.Print("building dms3 components for Linux/ARM7 platform...")
	_, err := dms3libs.RunCommand("env GOOS=linux GOARCH=arm GOARM=7 go build -o " + filepath.Join(releaseDir, platform) + "/go_dms3client go_dms3client.go")
	dms3libs.CheckErr(err)
	_, err = dms3libs.RunCommand("env GOOS=linux GOARCH=arm GOARM=7 go build -o " + filepath.Join(releaseDir, platform) + "/go_dms3server go_dms3server.go")
	dms3libs.CheckErr(err)
	_, err = dms3libs.RunCommand("env GOOS=linux GOARCH=arm GOARM=7 go build -o " + filepath.Join(releaseDir, platform) + "/go_dms3mail go_dms3mail.go")
	dms3libs.CheckErr(err)
	fmt.Println("success")

}
