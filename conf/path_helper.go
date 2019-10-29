package conf

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	logger "github.com/winjeg/go-commons/log"
)

func getPathSep() string {
	pathSep := "/"
	if strings.EqualFold(runtime.GOOS, "windows") {
		pathSep = `\`
	}
	return pathSep
}

// get parent directory of the current directory
func getParentDirectory(directory string) string {
	if strings.LastIndex(directory, getPathSep()) == len(directory)-1 {
		directory = directory[:len(directory)-1]
	}
	runes := []rune(directory)
	return string(runes[0:strings.LastIndex(directory, getPathSep())])
}

// get config file from where the source code lies
func getFileDirFromSrc(fileName, srcDir string) string {
	dir, _ := filepath.Abs(srcDir)
	tmpFile, _ := os.Stat(dir)
	if tmpFile != nil && !tmpFile.IsDir() {
		lastIdx := strings.LastIndex(dir, getPathSep())
		if lastIdx >= 0 {
			dir = dir[:lastIdx]
		}
	}

	f, err := os.Open(path.Join(dir, fileName))
	if err == nil {
		logger.Errors(f.Close())
		return dir
	}
	// change to the dir that contains the source code
	cdErr := os.Chdir(dir)
	logError(cdErr)

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
				return ""
			}
			return dir
		}
	}
	defer func() {
		if f != nil {
			err := f.Close()
			logger.Errors(err)
		}
	}()
	// return found file
	return dir
}

func logError(err error) {
	if err != nil {
		log.Println("error changing dir...", err)
	}
}

// get config file from where the executables lies
func getFileDirFromExecutable(fileName string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	f, err := os.Open(path.Join(dir, fileName))
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		return ""
	}
	return dir
}

// ReadConfigFile  read config file first from where the executable file lies
// then where the source code lies, or it's parent directory recursively
func FindDirOfFile(fileName, srcFile string) string {
	dir := ""
	dir = getFileDirFromExecutable(fileName)
	if len(dir) == 0 {
		dir = getFileDirFromSrc(fileName, srcFile)
	}
	return dir
}
