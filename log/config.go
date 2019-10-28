package log

import (
	"io"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type RotateFileConfig struct {
	Filename   string           `yaml:"file-name"`
	MaxSize    int              `yaml:"max-size"`
	MaxBackups int              `yaml:"max-backups"`
	MaxAge     int              `yaml:"max-age"`
	Level      logrus.Level     `yaml:"level"`
	Formatter  logrus.Formatter `yaml:"formatter"`
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
	return logrus.AllLevels[:hook.Config.Level+1]
}

func (hook *RotateFileHook) Fire(entry *logrus.Entry) (err error) {
	b, err := hook.Config.Formatter.Format(entry)
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
