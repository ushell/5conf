package log

import (
	"github.com/sirupsen/logrus"
	"path"
)

var logLevel = map[string]logrus.Level{
	"error":   logrus.ErrorLevel,
	"warning": logrus.WarnLevel,
	"info":    logrus.InfoLevel,
	"debug":   logrus.DebugLevel,
}

type Logger struct {
	Writer *logrus.Logger
}

type Config struct {
	Dir string
	Level string
	Filename string
}

var (
	l Logger
	c *Config
)

func Init(conf *Config) Logger {
	c = conf
	l.Writer = GetLogger()
	return l
}

func GetLogger() *logrus.Logger {
	if l.Writer != nil {
		return l.Writer
	}

	if c == nil {
		c = &Config{
			Dir: "",
			Level: "error",
			Filename: "error.log",
		}
	}

	logger := logrus.New()

	plugin := logPlugin(path.Join(c.Dir, c.Filename))

	// 设置输出
	logger.SetOutput(&plugin)

	// 设置日志级别
	level, _ := logrus.ParseLevel(c.Level)

	logger.SetLevel(level)

	// 日志格式
	logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006.01.02 15:04:05"})

	logger.AddHook(newHookLog())

	l.Writer = logger

	return logger
}

func Error(message string) {
	l.Writer.Error(message)
}

func Warning(message string) {
	l.Writer.Warning(message)
}

func Debug(message string) {
	l.Writer.Debug(message)
}

func Info(message string) {
	l.Writer.Info(message)
}