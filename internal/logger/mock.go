package logger

import (
	"bytes"
	"log"

	"go.uber.org/zap"
)

// Mock to logger.
type Mock struct {
	*log.Logger
}

func (m Mock) Error(args ...interface{}) {
	m.Println(args...)
}

func (m Mock) Errorf(template string, args ...interface{}) {
	m.Printf(template, args...)
}

func (m Mock) Debug(args ...interface{}) {
	m.Println(args...)
}

func (m Mock) Debugf(template string, args ...interface{}) {
	m.Printf(template, args...)
}

func (m Mock) Info(args ...interface{}) {
	m.Println(args...)
}

func (m Mock) Infof(template string, args ...interface{}) {
	m.Printf(template, args...)
}

func (m Mock) Warn(args ...interface{}) {
	m.Println(args...)
}

func (m Mock) Warnf(template string, args ...interface{}) {
	m.Printf(template, args...)
}

func (m Mock) SetLevel(level logLevel) {}

func (m Mock) With(args ...interface{}) *zap.SugaredLogger {
	m.Println(args...)
	var l *zap.Logger
	return l.Sugar()
}

// NewMock returns a mock logger.
func NewMock() Logger {
	bytes.NewBuffer([]byte{})
	return Mock{log.New(bytes.NewBuffer([]byte{}), "", 0)}
}
