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

func NewLogger(l LevelLogger) Logger {
	return &common{l}
}

// Convenient alias for logAdapter.Log()
func (l *common) Emergency(format string, args ...interface{}) {
	l.Log(level.EMERGENCY, format, args...)
}

// Convenient alias for logAdapter.Log()
func (l *common) Alert(format string, args ...interface{}) {
	l.Log(level.ALERT, format, args...)
}

// Convenient alias for logAdapter.Log()
func (l *common) Critical(format string, args ...interface{}) {
	l.Log(level.CRITICAL, format, args...)
}

// Convenient alias for logAdapter.Log()
func (l *common) Error(format string, args ...interface{}) {
	l.Log(level.ERROR, format, args...)
}

// Convenient alias for logAdapter.Log()
func (l *common) Warning(format string, args ...interface{}) {
	l.Log(level.WARNING, format, args...)
}

// Convenient alias for logAdapter.Log()
func (l *common) Notice(format string, args ...interface{}) {
	l.Log(level.NOTICE, format, args...)
}

// Convenient alias for logAdapter.Log()
func (l *common) Info(format string, args ...interface{}) {
	l.Log(level.INFO, format, args...)
}

// Convenient alias for logAdapter.Log()
func (l *common) Debug(format string, args ...interface{}) {
	l.Log(level.DEBUG, format, args...)
}
