package logger

import (
	"encoding/json"
	"errors"
)

// A global variable so that log functions can be directly acessed
var log Logger

// Fields type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

const (
	// Debug has verbose message
	Debug = "debug"
	// Info is default log level
	Info = "info"
	// Warn is for logging message about possible issues
	Warn = "warn"
	// Error is logging errors
	Error = "error"
	// Fatal is for logging fatal message. The system shutsdown after logging the message.
	Fatal = "fatal"
)

const (
	InstanceZapLogger int = iota
)

var (
	errInvalidLoggerInstance = errors.New("Invalid logger instance")
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	WithFields(keyValues Fields) Logger
}

// Configuration stores the config for the logger
type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
}

func NewLogger(config Configuration, loggerInstance int) error {
	switch loggerInstance {
	case InstanceZapLogger:
		logger, err := NewZapLogger(config)
		if err != nil {
			return err
		}
		log = logger
		return nil
	default:
		return errInvalidLoggerInstance
	}
}

func WriteLog(level string, component, action string, payload interface{}, desc, method, traceId string) {
	data, _ := json.Marshal(payload)
	contextLogger := WithFields(Fields{
		"component": component,
		"action":    action,
		"payload":   string(data),
		"method":    method,
		"traceId":   traceId})
	switch level {
	case Info:
		contextLogger.Infof(desc)
	case Error:
		contextLogger.Errorf(desc)
	case Debug:
		contextLogger.Debugf(desc)
	default:
		contextLogger.Infof(desc)
	}
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func WithFields(keyValues Fields) Logger {
	cek := keyValues
	_ = cek
	return log.WithFields(keyValues)
}
