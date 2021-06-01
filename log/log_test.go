package log

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"errors"
	"testing"
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
	B Settings
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

func TestNewLogger(t *testing.T) {
	settings.FileConfig = RotateFileConfig{
		Filename:   "E:/Desktop/a.log",
		MaxSize:    500,
		MaxBackups: 7,
		MaxAge:     7,
		Level:      "debug",
		Formatter:  "text",
	}
	l := NewLogger(settings)
	l.Info("hello, world!")
	IgnoreErrors()
	testErr()
}

func testErr() {
	someErr()
}

func someErr() {
	Errors(errors.New("abc"))
}

func BenchmarkLogErrors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Errors(errors.New("hello"))
	}
	b.ReportAllocs()
}
