package log

import (
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"

	"io"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
)

var (
	//default settings
	settings = LogSettings{
		Output:       "std",
		Format:       "text",
		Level:        "info",
		ReportCaller: true,
	}
	lock   sync.Mutex
	logger *logrus.Logger
)

// supporting ini/yaml/json
type LogSettings struct {
	Output       string `json:"output" yaml:"output" ini:"output"`
	Format       string `json:"format" yaml:"format" ini:"format"`
	Level        string `json:"level" yaml:"level" ini:"level"`
	ReportCaller bool   `json:"reportCaller" yaml:"report-caller" ini:"report-caller"`
}

func initLogger(c interface{}) {
	var conf = settings
	if c != nil {
		conf = getConf(c)
	}
	l := logrus.New()
	// for windows no color output
	if windows() && strings.EqualFold(conf.Format, "colored") {
		l.SetOutput(ansicolor.NewAnsiColorWriter(getOutput(conf)))
	} else {
		l.SetOutput(getOutput(conf))
	}
	l.SetFormatter(getFormatter(conf))
	l.SetLevel(getLogLevel(conf))
	l.SetReportCaller(conf.ReportCaller)
	logger = l
}

func GetLogger(c interface{}) *logrus.Logger {
	if logger != nil {
		return logger
	} else {
		lock.Lock()
		initLogger(c)
		lock.Unlock()
	}
	return logger
}

var conf *LogSettings

// check all fields of a struct
func getConfig(raw interface{}) {
	if v, ok := raw.(LogSettings); ok {
		conf = &v
	}
	if v, ok := raw.(*LogSettings); ok && v != nil {
		conf = v
	}
	getType := reflect.TypeOf(raw)
	getValue := reflect.ValueOf(raw)
	if getType.Kind() == reflect.Struct {
		for i := 0; i < getType.NumField(); i++ {
			value := getValue.Field(i).Interface()
			if reflect.TypeOf(value).Kind() != reflect.Struct {
				continue
			}
			getConfig(value)
		}
	}
}

func getConf(raw interface{}) LogSettings {
	getConfig(raw)
	if conf == nil {
		return settings
	}
	return *conf
}

func windows() bool {
	return strings.EqualFold(runtime.GOOS, "windows")
}

// get log level, default level info
func getLogLevel(settings LogSettings) logrus.Level {
	switch strings.ToLower(settings.Level) {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}

func getFormatter(c LogSettings) logrus.Formatter {
	switch c.Format {
	case "colored":
		return &logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		}
	case "text":
		return &logrus.TextFormatter{}
	case "json":
		return &logrus.JSONFormatter{}
	default:
		return &logrus.TextFormatter{
			FullTimestamp: true,
			ForceColors:   true,
		}
	}
}

func getOutput(c LogSettings) io.Writer {
	switch c.Output {
	case "std":
		return os.Stdout

	default:
		return os.Stdout
	}
}
