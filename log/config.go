package log

import (
	"io"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type RotateFileConfig struct {
	Enable     bool   `json:"enable" yaml:"enable" ini:"enable"`
	Filename   string `json:"filename" yaml:"filename" ini:"filename"`
	MaxSize    int    `json:"maxSize" yaml:"max-size" ini:"max-size"`
	MaxBackups int    `json:"maxBackups" yaml:"max-backups" ini:"max-backups"`
	MaxAge     int    `json:"maxAge" yaml:"max-age" ini:"max-age"`
	Level      string `json:"level" yaml:"level" ini:"level"`
	Formatter  string `json:"formatter" yaml:"formatter" ini:"formatter"`
}

type RotateFileHook struct {
	Config    RotateFileConfig
	logWriter io.Writer
}

func NewRotateFileHook(config RotateFileConfig) (logrus.Hook, error) {
	hook := RotateFileHook{
		Config: config,
	}
	hook.logWriter = &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
	}
	return &hook, nil
}

func (hook *RotateFileHook) Levels() []logrus.Level {
	return logrus.AllLevels[:convertLogLevel(hook.Config.Level)+1]
}

func (hook *RotateFileHook) Fire(entry *logrus.Entry) (err error) {
	formatter := convertFormatter(hook.Config.Formatter)
	b, err := formatter.Format(entry)
	if err != nil {
		return err
	}
	if len(hook.Config.Filename) > 0 {
		_, err := hook.logWriter.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}

func convertLogLevel(level string) logrus.Level {
	switch strings.ToLower(level) {
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

func convertFormatter(format string) logrus.Formatter {
	switch format {
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
