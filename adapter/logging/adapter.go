/**
 * CommonLog
 *
 * Copyright © 2014 Trevor N. Suarez (Rican7)
 */

package logging

import (
	"github.com/Rican7/commonlog/adapter"
	"github.com/Rican7/commonlog/level"
	logging "github.com/op/go-logging"
)

/**
 * Types
 */

type loggingAdapter struct {
	adaptee *logging.Logger
}

/**
 * Functions
 */

// Construct a new instance by injecting the adaptee
func New(adaptee *logging.Logger) adapter.LogAdapter {
	return &loggingAdapter{adaptee}
}

// Get the underyling adaptee
func (a *loggingAdapter) Adaptee() interface{} {
	return a.adaptee
}

// Convenient alias for loggingAdapter.Log()
func (a *loggingAdapter) Emergency(format string, args ...interface{}) {
	a.Log(level.EMERGENCY, format, args)
}

// Convenient alias for loggingAdapter.Log()
func (a *loggingAdapter) Alert(format string, args ...interface{}) {
	a.Log(level.ALERT, format, args)
}

// Convenient alias for loggingAdapter.Log()
func (a *loggingAdapter) Critical(format string, args ...interface{}) {
	a.Log(level.CRITICAL, format, args)
}

// Convenient alias for loggingAdapter.Log()
func (a *loggingAdapter) Error(format string, args ...interface{}) {
	a.Log(level.ERROR, format, args)
}

// Convenient alias for loggingAdapter.Log()
func (a *loggingAdapter) Warning(format string, args ...interface{}) {
	a.Log(level.WARNING, format, args)
}

// Convenient alias for loggingAdapter.Log()
func (a *loggingAdapter) Notice(format string, args ...interface{}) {
	a.Log(level.NOTICE, format, args)
}

// Convenient alias for loggingAdapter.Log()
func (a *loggingAdapter) Info(format string, args ...interface{}) {
	a.Log(level.INFO, format, args)
}

// Convenient alias for loggingAdapter.Log()
func (a *loggingAdapter) Debug(format string, args ...interface{}) {
	a.Log(level.DEBUG, format, args)
}

/**
 * Log an error based on a specified level, a format, and a splat of arguments
 */
func (a *loggingAdapter) Log(lvl level.LogLevel, format string, args ...interface{}) {
	// Validate the passed in level
	if ok, err := lvl.IsValid(); !ok {
		panic(err)
	}

	switch lvl {
	case level.EMERGENCY:
		// TODO: Handle this level once a generic log method is in the logging.Logger
		fallthrough
	case level.ALERT:
		// TODO: Handle this level once a generic log method is in the logging.Logger
		fallthrough
	case level.CRITICAL:
		a.adaptee.Critical(format, args)
	case level.ERROR:
		a.adaptee.Error(format, args)
	case level.WARNING:
		a.adaptee.Warning(format, args)
	case level.NOTICE:
		a.adaptee.Notice(format, args)
	case level.INFO:
		a.adaptee.Info(format, args)
	case level.DEBUG:
		a.adaptee.Debug(format, args)
	}
}
