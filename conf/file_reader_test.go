package goconf

import (
	"runtime"
	"testing"
)

const testFileName  = "readme.md"

// read config file
func TestReadConfigFile(t *testing.T) {
	if _, fileNameWithPath, _, ok := runtime.Caller(1); ok {
		d := ReadConfigFile(testFileName, fileNameWithPath)
		if d == nil {
			t.FailNow()
		}
	}
}
