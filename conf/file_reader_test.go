package goconf

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

const testFileName  = "readme.md"

// read config file
func TestReadConfigFile(t *testing.T) {
	if _, fileNameWithPath, _, ok := runtime.Caller(1); ok {
		d := ReadConfigFile(testFileName, fileNameWithPath)
		assert.NotNil(t, d)
	}
}
