/**
 * CommonLog
 *
 * Copyright © 2015 Trevor N. Suarez (Rican7)
 */

package commonlog

import (
	"github.com/Rican7/commonlog/level"
)

/**
 * Types
 */

type common struct {
	LevelLogger
}

/**
 * Functions
 */

// NewLogger creates a new Logger instance from a LevelLogger by using a common
// pass-through pattern for each separate level-specific logging method to the
// passed LevelLogger.Log() method
func NewLogger(l LevelLogger) Logger {
	return &common{l}
}

// Emergency calls the LevelLogger's Log() method with the level.Emergency level
func (l *common) Emergency(format string, args ...interface{}) {
	l.Log(level.EMERGENCY, format, args...)
}

// Alert calls the LevelLogger's Log() method with the level.Alert level
func (l *common) Alert(format string, args ...interface{}) {
	l.Log(level.ALERT, format, args...)
}

// Critical calls the LevelLogger's Log() method with the level.Critical level
func (l *common) Critical(format string, args ...interface{}) {
	l.Log(level.CRITICAL, format, args...)
}

// Error calls the LevelLogger's Log() method with the level.Error level
func (l *common) Error(format string, args ...interface{}) {
	l.Log(level.ERROR, format, args...)
}

// Warning calls the LevelLogger's Log() method with the level.Warning level
func (l *common) Warning(format string, args ...interface{}) {
	l.Log(level.WARNING, format, args...)
}

// Notice calls the LevelLogger's Log() method with the level.Notice level
func (l *common) Notice(format string, args ...interface{}) {
	l.Log(level.NOTICE, format, args...)
}

// Info calls the LevelLogger's Log() method with the level.Info level
func (l *common) Info(format string, args ...interface{}) {
	l.Log(level.INFO, format, args...)
}

// Debug calls the LevelLogger's Log() method with the level.Debug level
func (l *common) Debug(format string, args ...interface{}) {
	l.Log(level.DEBUG, format, args...)
}
