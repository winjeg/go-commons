package log

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// the full example usage of log package
func TestLog(t *testing.T) {
	l := GetLogger(nil)
	l.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	contextLogger := l.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})
	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
	l.Trace("Something very low level.")
	l.Debug("Useful debugging information.")
	l.Info("Something noteworthy happened!")
	l.Warn("You should probably take a look at this.")
	l.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	// log.Fatal("Bye.")
	// Calls panic() after logging
	// log.Panic("I'm bailing.")
}

type abc struct {
	A string
	B LogSettings
}

type bcd struct {
	B string
	A abc
}

type cde struct {
	A string
}

func TestStruct(t *testing.T) {
	var x bcd
	x.A.B.Level = "DEBUG"
	v := getConf(x)
	assert.NotNil(t, v)
	var a cde
	m := getConf(a)
	assert.NotNil(t, m)
}

func TestGetLogger(t *testing.T) {
	lg := logrus.New()
	lg.Out = &ConfigWriter{FileName: "E:/Desktop/a.log", Std: true}
	lg.Info("aaaaaaaaaaaaaa")
}
