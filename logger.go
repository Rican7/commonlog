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

type Logger interface {
	Emergency(string, ...interface{})
	Alert(string, ...interface{})
	Critical(string, ...interface{})
	Error(string, ...interface{})
	Warning(string, ...interface{})
	Notice(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})

	LevelLogger
}

type LevelLogger interface {
	Log(level.LogLevel, string, ...interface{})
}
