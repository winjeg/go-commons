package goconf

import (
	"strings"
	"testing"
)

const (
	testYmlFile = "test.yaml"
	testIniFile = "test.ini"

	host = "10.1.1.1"
	port = 3306
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
	if err != nil {
		t.FailNow()
	}
	if !strings.EqualFold(x.DbAddr, host) || x.Port != port {
		t.FailNow()
	}
}

func TestIni2Object(t *testing.T) {
	var x TestMyConf
	err := Ini2Object(testIniFile, &x)
	if err != nil {
		t.FailNow()
	}
	if !strings.EqualFold(x.Mysql.Host, host) || x.Mysql.Port != port || !strings.EqualFold(testName, x.Name) {
		t.FailNow()
	}
}
