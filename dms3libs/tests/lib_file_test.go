package dms3libs_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
)

func TestIsFile(t *testing.T) {

	testFile := "lib_file_test.go"

	if !dms3libs.IsFile(testFile) {
		t.Error(testFile + " file not found, but should have been")
	}

}

func TestMkDir(t *testing.T) {

	newDir := filepath.Join(dms3libs.GetPackageDir(), "tmpDir")
	dms3libs.MkDir(newDir)

	if dms3libs.IsFile(newDir) {
		dms3libs.CheckErr(os.RemoveAll(newDir))
	} else {
		t.Error("directory " + newDir + " not found, but should have been")
	}

}

func TestRmDir(t *testing.T) {

	newDir := filepath.Join(dms3libs.GetPackageDir(), "tmpDir")

	dms3libs.MkDir(newDir)
	dms3libs.RmDir(newDir)

	if dms3libs.IsFile(newDir) {
		t.Error("directory not removed, but should have been")
	}

}

func TestWalkDir(t *testing.T) {

	dirCount := 0
	fileCount := 0

	newDir := filepath.Join(dms3libs.GetPackageDir(), "tmpDir")
	newFile := "tmpFile"

	dms3libs.MkDir(newDir)
	dms3libs.CopyFile(filepath.Join(dms3libs.GetPackageDir(), "lib_audio_test.wav"), newDir+"/"+newFile)

	for _, dirType := range dms3libs.WalkDir(newDir) {

		if dirType == 0 {
			dirCount += 1
		} else {
			fileCount += 1
		}

	}

	if dirCount != 1 {
		t.Error("wrong directory count in", newDir)
	}

	if fileCount != 1 {
		t.Error("wrong file count in", newDir)
	}

	// avoid using os.RemoveAll() in the event of an error with newDir creation
	os.Remove(newDir + "/" + newFile)
	os.Remove(newDir)

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
