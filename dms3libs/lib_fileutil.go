package dms3libs

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// IsFile returns true/false on existence of file passed in
func IsFile(filename string) bool {

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true

}

// MkDir creates a new folder with permissions passed in
func MkDir(newPath string) {

	error := os.MkdirAll(newPath, os.ModePerm)
	CheckErr(error)

}

// RmDir creates a new folder with permissions passed in
func RmDir(dir string) {

	if IsFile(dir) {
		err := os.RemoveAll(dir)
		CheckErr(err)
	}

}

// walkDir generates a map of directories (0) and files (1)
func walkDir(dirname string) map[string]int {

	fileList := map[string]int{}

	error := filepath.Walk(dirname, func(path string, f os.FileInfo, err error) error {

		// exclude root directory
		if f.IsDir() && f.Name() == dirname {
			return nil
		}

		if f.IsDir() {
			fileList[path] = 0 // directory
		} else {
			fileList[path] = 1 // file
		}

		return nil
	})
	CheckErr(error)

	return fileList

}

// CopyFile copies a file from src to dest
func CopyFile(src string, dest string) {

	srcFile, err := os.Open(src)
	CheckErr(err)
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	CheckErr(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	CheckErr(err)

	err = destFile.Sync()
	CheckErr(err)

}

// CopyDir copies a directory from srcDir to destDir
func CopyDir(srcDir string, destDir string) {

	dirTree := walkDir(srcDir)

	// create directory tree...
	for dirName, dirType := range dirTree {

		if dirType == 0 {
			MkDir(destDir + "/" + strings.TrimLeft(dirName, srcDir))
		}

	}

	// ...then copy files into directory tree
	for dirName, dirType := range dirTree {

		if dirType == 1 {
			CopyFile(dirName, destDir+"/"+strings.TrimLeft(dirName, srcDir))
		}

	}

}
