// copyleft anyone can use this code freely, and this repo will remain active
// welcome to come up with any kind of issues relating to this project

package goconf

import (
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"

	"errors"
	"runtime"
)

// yaml config file to an object
func Yaml2Object(fileName string, object interface{}) error {
	var srcDir = "."
	if _, fileNameWithPath, _, ok := runtime.Caller(1); ok {
		srcDir = fileNameWithPath
	}
	d := ReadConfigFile(fileName, srcDir)
	if d == nil {
		return errors.New("the specified file cannot be found")
	}
	return yaml.Unmarshal(d, object)
}

// ini config to object
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
