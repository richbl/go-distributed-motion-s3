// Package dms3libs file provides file services for dms3 device components
//
package dms3libs

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// IsFile returns true/false on existence of file/folder passed in
//
func IsFile(filename string) bool {

	_, error := os.Stat(filename)
	return (!errors.Is(error, fs.ErrNotExist))

}

// MkDir creates a new folder with permissions passed in
//
func MkDir(newPath string) {

	error := os.MkdirAll(newPath, os.ModePerm)
	CheckErr(error)

}

// RmDir removes the folder passed in
//
func RmDir(dir string) {

	if IsFile(dir) {
		error := os.RemoveAll(dir)
		CheckErr(error)
	}

}

// WalkDir generates a map of directories (0) and files (1)
//
func WalkDir(dirname string) map[string]int {

	fileList := map[string]int{}
	error := filepath.WalkDir(dirname, func(path string, f os.DirEntry, err error) error {

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
//
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

	srcAttrib, err := os.Stat(src)
	CheckErr(err)

	err = os.Chmod(dest, srcAttrib.Mode())
	CheckErr(err)

}

// CopyDir copies a directory from srcDir to destDir
//
func CopyDir(srcDir string, destDir string) {

	pathRoot := filepath.Dir(srcDir)

	if pathRoot == "." {
		pathRoot = ""
	}

	dirTree := WalkDir(srcDir)

	// create directory tree...
	for dirName, dirType := range dirTree {

		if dirType == 0 {
			MkDir(filepath.Join(destDir, dirName[len(pathRoot):]))
		}

	}

	// ...then copy files into directory tree
	for dirName, dirType := range dirTree {

		if dirType == 1 {
			CopyFile(dirName, filepath.Join(destDir, dirName[len(pathRoot):]))
		}

	}

}

// CountFilesInDir recursively counts the files in the dir passed in
//
func CountFilesInDir(srcDir string) int {

	fileCount := 0
	dirTree := WalkDir(srcDir)

	for _, dirType := range dirTree {

		if dirType == 1 {
			fileCount++
		}

	}

	return fileCount

}

// CheckFileLocation checks/sets the location of file and pathname passed in as defined in various
// TOML config files, returning a fully qualified path
//
func CheckFileLocation(configPath string, fileDir string, fileLocation *string, filename string) {

	// set default template location
	if *fileLocation == "" {
		*fileLocation = filepath.Join(configPath, fileDir)
	}

	if filename == "" || !IsFile(filepath.Join(*fileLocation, filename)) {
		LogFatal("unable to set file location... check TOML configuration file")
	}

}
