package libconfig

import (
	"os"
	"path"
)

// GetPath returns the current path (curPath) of the executable
func GetPath() string {
	res, err := os.Executable()
	if err != nil {
		panic(err)
	}
	curPath := path.Dir(res)
	return curPath
}
