// copyleft anyone can use this code freely, and this repo will remain active
// welcome to come up with any kind of issues relating to this project

package conf

import (
	"errors"
	"io/ioutil"
	"os"
	"runtime"

	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
)

// Yaml2Object  yaml config file to an object
func Yaml2Object(fileName string, object interface{}) error {
	f, openErr := os.Open(fileName)
	var data []byte
	if openErr != nil {
		var srcDir = "."
		if _, fileNameWithPath, _, ok := runtime.Caller(1); ok {
			srcDir = fileNameWithPath
		}
		data = ReadConfigFile(fileName, srcDir)
		if data == nil {
			return errors.New("the specified file cannot be found")
		}
	} else {
		d, err := ioutil.ReadAll(f)
		if err != nil {
			return errors.New("read file err: " + err.Error())
		}
		data = d
	}

	return yaml.Unmarshal(data, object)
}

// Ini2Object ini config to object
func Ini2Object(fileName string, object interface{}) error {
	var srcDir = "."
	if _, fileNameWithPath, _, ok := runtime.Caller(1); ok {
		srcDir = fileNameWithPath
	}
	d := ReadConfigFile(fileName, srcDir)
	if d == nil {
		return errors.New("the specified file cannot be found")
	}
	return ini.MapTo(object, d)
}
