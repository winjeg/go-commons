// copyleft anyone can use this code freely, and this repo will remain active
// welcome to come up with any kind of issues relating to this project

package goconf

import (
	"github.com/winjeg/go-commons/log"

	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// get parent directory of the current directory
func getParentDirectory(directory string) string {
	if strings.LastIndex(directory, "/") == len(directory)-1 {
		directory = directory[:len(directory)-1]
	}
	runes := []rune(directory)
	return string(runes[0:strings.LastIndex(directory, "/")])
}

// get config file from where the source code lies
func getConfigFileFromSrc(fileName, srcDir string) *os.File {
	dir := path.Dir(srcDir)
	f, err := os.Open(path.Join(dir, fileName))
	// change to the dir that contains the source code
	ercd := os.Chdir(dir)
	logError(ercd)

	reachRoot := false
	for err != nil {
		er := os.Chdir("..")
		logError(er)
		dir = getParentDirectory(dir)
		if strings.EqualFold(runtime.GOOS, "windows") {
			// windows
			if len(dir) < 3 {
				// like C:  or D:
				reachRoot = true
			}
		} else {
			// osx or linux
			if len(dir) < 2 {
				// like /
				reachRoot = true
			}
		}
		f, err = os.Open(path.Join(dir, fileName))
		// find to the root of all dirs
		if reachRoot {
			if err != nil {
				return nil
			}
			return f
		}
	}
	// return found file
	return f
}

func logError(err error) {
	logger := log.GetLogger(nil)
	if err != nil {
		logger.Error("error changing dir...", err)
	}
}

// get config file from where the executables lies
func getConfigFileFromExecutable(fileName string) *os.File {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil
	}
	f, err := os.Open(path.Join(dir, fileName))
	if err != nil {
		return nil
	}
	return f
}

// read config file first from where the executable file lies
// then where the source code lies, or it's parent directory recursively
func ReadConfigFile(fileName, srcFile string) []byte {
	f := getConfigFileFromExecutable(fileName)
	if f == nil {
		f = getConfigFileFromSrc(fileName, srcFile)
	}
	if f == nil {
		return nil
	}
	d, err := ioutil.ReadAll(f)
	if err != nil {
		return nil
	}
	return d
}
