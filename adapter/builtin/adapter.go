/**
 * CommonLog
 *
 * Copyright © 2014 Trevor N. Suarez (Rican7)
 */

package builtin

import (
	"github.com/Rican7/commonlog/adapter"
	"github.com/Rican7/commonlog/level"
	"io"
	"log"
)

/**
 * Types
 */

type logAdapter struct {
	adaptee *log.Logger
}

/**
 * Functions
 */

// Construct a new instance by injecting the adaptee
func New(adaptee *log.Logger) adapter.LogAdapter {
	return &logAdapter{adaptee}
}

// Construct a new instance by building the adaptee
func NewWithSetup(out io.Writer, prefix string, flag int) adapter.LogAdapter {
	builtinLogger := log.New(out, prefix, flag)

	return &logAdapter{builtinLogger}
}

// Get the underyling adaptee
func (a *logAdapter) Adaptee() interface{} {
	return a.adaptee
}

// Convenient alias for logAdapter.Log()
func (a *logAdapter) Emergency(format string, args ...interface{}) {
	a.Log(level.EMERGENCY, format, args)
}

// Convenient alias for logAdapter.Log()
func (a *logAdapter) Alert(format string, args ...interface{}) {
	a.Log(level.ALERT, format, args)
}

// Convenient alias for logAdapter.Log()
func (a *logAdapter) Critical(format string, args ...interface{}) {
	a.Log(level.CRITICAL, format, args)
}

// Convenient alias for logAdapter.Log()
func (a *logAdapter) Error(format string, args ...interface{}) {
	a.Log(level.ERROR, format, args)
}

// Convenient alias for logAdapter.Log()
func (a *logAdapter) Warning(format string, args ...interface{}) {
	a.Log(level.WARNING, format, args)
}

// Convenient alias for logAdapter.Log()
func (a *logAdapter) Notice(format string, args ...interface{}) {
	a.Log(level.NOTICE, format, args)
}

// Convenient alias for logAdapter.Log()
func (a *logAdapter) Info(format string, args ...interface{}) {
	a.Log(level.INFO, format, args)
}

// Convenient alias for logAdapter.Log()
func (a *logAdapter) Debug(format string, args ...interface{}) {
	a.Log(level.DEBUG, format, args)
}

/**
 * Log an error based on a specified level, a format, and a splat of arguments
 */
func (a *logAdapter) Log(lvl level.LogLevel, format string, args ...interface{}) {
	// Validate the passed in level
	if ok, err := lvl.IsValid(); !ok {
		panic(err)
	}

	switch lvl {
	case level.EMERGENCY:
		// TODO: Handle this level once a generic log method is in the log.Logger
		fallthrough
	case level.ALERT:
		// TODO: Handle this level once a generic log method is in the log.Logger
		fallthrough
	case level.CRITICAL:
		// TODO: Handle this level once a generic log method is in the log.Logger
		a.adaptee.Fatalf(format, args)
	case level.ERROR:
		// TODO: Handle this level once a generic log method is in the log.Logger
		a.adaptee.Panicf(format, args)
	case level.WARNING:
		// TODO: Handle this level once a generic log method is in the log.Logger
		fallthrough
	case level.NOTICE:
		// TODO: Handle this level once a generic log method is in the log.Logger
		fallthrough
	case level.INFO:
		// TODO: Handle this level once a generic log method is in the log.Logger
		fallthrough
	case level.DEBUG:
		a.adaptee.Printf(format, args)
	}
}
