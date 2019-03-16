package goconf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testYmlFile = "test.yaml"
	testIniFile = "test.ini"

	host     = "10.1.1.1"
	port     = 3306
	testName = "tom"
)

type TestYmlConf struct {
	DbAddr string `yaml:"dbAddr"`
	Port   int    `yaml:"dbPort"`
}

type TestMyConf struct {
	Mysql TestIniConf `ini:"mysql"`
	Name  string      `ini:"name"`
}

type TestIniConf struct {
	Host string `ini:"host"`
	Port int    `ini:"port"`
}

func TestYaml2Object(t *testing.T) {
	var x TestYmlConf
	err := Yaml2Object(testYmlFile, &x)
	assert.Nil(t, err)
	assert.Equal(t, host, x.DbAddr)
	assert.Equal(t, port, x.Port)
}

func TestIni2Object(t *testing.T) {
	var x TestMyConf
	err := Ini2Object(testIniFile, &x)
	assert.Nil(t, err)
	assert.Equal(t, host, x.Mysql.Host)
	assert.Equal(t, port, x.Mysql.Port)
	assert.Equal(t, testName, x.Name)
}
