package dms3libs_test

import (
	"go-distributed-motion-s3/dms3libs"
	"os"
	"path/filepath"
	"testing"
)

func init() {
	dms3libs.LoadLibConfig("../../config/dms3libs.toml")
}

func TestIsFile(t *testing.T) {

	testFile := "lib_util_test.go"

	if !dms3libs.IsFile(testFile) {
		t.Error(testFile + " file not found, but should have been")
	}

}

func TestMkDir(t *testing.T) {

	dms3libs.MkDir(filepath.Join(dms3libs.GetPackageDir(), "tmpDir"))

	if dms3libs.IsFile("tmpDir") {
		os.Remove("tmpDir")
	} else {
		t.Error("directory not found, but should have been")
	}

}

func TestRmDir(t *testing.T) {

	dms3libs.MkDir(filepath.Join(dms3libs.GetPackageDir(), "tmpDir"))
	dms3libs.RmDir(filepath.Join(dms3libs.GetPackageDir(), "tmpDir"))

	if dms3libs.IsFile("tmpDir") {
		t.Error("directory not removed, but should have been")
	}

}

func TestWalkDir(t *testing.T) {

	dirCount := 0
	fileCount := 0

	dms3libs.MkDir(filepath.Join(dms3libs.GetPackageDir(), "tmpDir"))
	dms3libs.CopyFile(filepath.Join(dms3libs.GetPackageDir(), "lib_audio_test.wav"), filepath.Join(dms3libs.GetPackageDir(), "tmpDir/tmpFile"))
	currentDir := filepath.Join(dms3libs.GetPackageDir(), "tmpDir")

	for _, dirType := range dms3libs.WalkDir(currentDir) {

		switch dirType {
		case 0:
			dirCount += 1
		case 1:
			fileCount += 1
		}
	}

	if dirCount != 1 {
		t.Error("wrong directory count in", currentDir)
	}

	if fileCount != 1 {
		t.Error("wrong file count in", currentDir)
	}

	os.Remove(currentDir)

}

func TestCopyFile(t *testing.T) {

	dms3libs.CopyFile(filepath.Join(dms3libs.GetPackageDir(), "lib_audio_test.wav"), "tmpfile")

	if dms3libs.IsFile("tmpfile") {
		os.Remove("tmpfile")
	} else {
		t.Error("file not found, but should have been")
	}

}

func TestCopyDir(t *testing.T) {

	dms3libs.CopyDir(dms3libs.GetPackageDir(), "tmpDir")

	if dms3libs.IsFile("tmpDir") {
		dms3libs.RmDir(filepath.Join(dms3libs.GetPackageDir(), "tmpDir"))
	} else {
		t.Error("directory not found, but should have been")
	}

}

func TestCountFilesInDir(t *testing.T) {

	dms3libs.MkDir(filepath.Join(dms3libs.GetPackageDir(), "tmpDir"))
	dms3libs.CopyFile(filepath.Join(dms3libs.GetPackageDir(), "lib_audio_test.wav"), filepath.Join(dms3libs.GetPackageDir(), "tmpDir/tmpFile"))
	currentDir := filepath.Join(dms3libs.GetPackageDir(), "tmpDir")

	if dms3libs.CountFilesInDir(currentDir) != 1 {
		t.Error("incorrect file count")
	}

	dms3libs.RmDir(currentDir)

}
