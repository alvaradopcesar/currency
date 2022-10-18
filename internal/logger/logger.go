package logger

import "go.uber.org/zap"

// Logger to print information to standard output.
type Logger interface {
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	SetLevel(level logLevel)
	With(args ...interface{}) *zap.SugaredLogger
}

// logLevel represents a level of log.
type logLevel int

const (
	ERROR logLevel = iota
	WARNING
	INFO
	DEBUG
)

// logger is an adapted zap logger.
type logger struct {
	*zap.SugaredLogger
	config *zap.Config
}

// New creates a new zap logger.
func New(serviceName string, production bool) Logger {
	if production {
		config := zap.NewProductionConfig()
		l, _ := config.Build()
		return &logger{l.Sugar().With(zap.String("service", serviceName)), &config}
	}

	config := zap.NewDevelopmentConfig()
	l, _ := config.Build()
	return &logger{l.Sugar().With(zap.String("service", serviceName)), &config}
}

// SetLevel sets the current log level.
func (l *logger) SetLevel(level logLevel) {
	switch level {
	case DEBUG:
		l.config.Level.SetLevel(zap.DebugLevel)
	case INFO:
		l.config.Level.SetLevel(zap.InfoLevel)
	case WARNING:
		l.config.Level.SetLevel(zap.WarnLevel)
	case ERROR:
		l.config.Level.SetLevel(zap.ErrorLevel)
	}
}
