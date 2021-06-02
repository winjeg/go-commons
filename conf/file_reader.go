// copyleft anyone can use this code freely, and this repo will remain active
// welcome to come up with any kind of issues relating to this project

package conf

import (
	"io/ioutil"
	"os"
	"path"
)

// ReadConfigFile  read config file first from where the executable file lies
// then where the source code lies, or it's parent directory recursively
func ReadConfigFile(fileName, srcFile string) []byte {
	dir := FindDirOfFile(fileName, srcFile)
	if len(dir) == 0 {
		return nil
	}
	f, openErr := os.Open(path.Join(dir, fileName))
	if openErr != nil {
		return nil
	}
	d, err := ioutil.ReadAll(f)
	defer func() {
		if f != nil {
			logError(f.Close())
		}
	}()
	if err != nil {
		return nil
	}
	return d
}
