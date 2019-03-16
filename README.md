# go-commons
[![Build Status](https://travis-ci.org/winjeg/go-commons.svg?branch=master)](https://travis-ci.org/winjeg/go-commons)
[![Go Report Card](https://goreportcard.com/badge/github.com/winjeg/go-commons)](https://goreportcard.com/report/github.com/winjeg/go-commons)

golang commonly used  utils and many thing else.
## conf
conf is a package for reading config file for golang projects.
in many cases when you need a config file to do some config, but your program or your unit tests 
can't read the config file correctly, because where to put the config file really sucks.

This program help your program find the real location of your config file if you don't mess it up.
The recommended position of the config file is under the root of your source code.
When you deploy executables the config file should be along with the executables.

supports both yaml and ini config files
### examples
```go
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

```
## log
a simple tool to use `logrus` as a logger and with simple configurations

You can use the struct from our package, in your configuration object, and pass the configuration object to 
`log.GetLogger(conf)`, it will work.
```go
// supporting ini/yaml/json
type LogSettings struct {
	Output       string `json:"output" yaml:"output" ini:"output"`
	Format       string `json:"format" yaml:"format" ini:"format"`
	Level        string `json:"level" yaml:"level" ini:"level"`
	ReportCaller bool   `json:"reportCaller" yaml:"report-caller" ini:"report-caller"`
}

```

```go
package any
import (
 "fmt"
 "github.com/winjeg/go-commons/log"
)

var logger = log.GetLogger(nil)
func loggerTest() {
        // Debug, Warn, Trace, Error
    	logger.Info("Something noteworthy happened!")
}

```

## http client
a simple client to make remote http request  and return string response body

currently supported method:
1. Get
2. Put   (content type: json)
3. Post (content type: json)
4. Delete (content type: json)

example
```go
package any
import (
	"fmt"
	"github.com/winjeg/go-commons/httpclient"
)

func test() {
    fmt.Println(httpclient.Get("https://www.google.com"))
}
```

## str
an unique id generator to generate unique id from

```go

func TestUUID(t *testing.T) {
	id1 := UUID()
	id2 := UUID()
	assert.NotEqual(t, id1, id2)
}

```

## properties
read properties from string or from file,and convert a map to a property file
```properties
ip=127.0.0.1
name=tom
```


## and others
... to be planned


