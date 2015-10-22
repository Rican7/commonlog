/**
 * CommonLog
 *
 * Copyright © 2015 Trevor N. Suarez (Rican7)
 */

package logging

import (
	"github.com/Rican7/commonlog"
	"github.com/Rican7/commonlog/adapter"
	"github.com/Rican7/commonlog/level"
	logging "github.com/op/go-logging"
)

/**
 * Types
 */

type adapted logging.Logger

type loggingAdapter struct {
	commonlog.Logger
	adaptee *adapted
}

/**
 * Functions
 */

// New constructs a new instance by injecting the adaptee
func New(adaptee *logging.Logger) adapter.LogAdapter {
	adapted := adapted(*adaptee)

	return &loggingAdapter{
		commonlog.NewLogger(&adapted),
		&adapted,
	}
}

// Adaptee gets the underyling adaptee
func (a *loggingAdapter) Adaptee() interface{} {
	return a.adaptee
}

// Log an error based on a specified level, a format, and a splat of arguments
func (a *adapted) Log(lvl level.LogLevel, format string, args ...interface{}) {
	// Validate the passed in level
	if ok, err := lvl.IsValid(); !ok {
		panic(err)
	}

	adaptee := (*logging.Logger)(a)

	switch lvl {
	case level.EMERGENCY:
		// TODO: Handle this level once a generic log method is in the logging.Logger
		fallthrough
	case level.ALERT:
		// TODO: Handle this level once a generic log method is in the logging.Logger
		fallthrough
	case level.CRITICAL:
		adaptee.Critical(format, args...)
	case level.ERROR:
		adaptee.Error(format, args...)
	case level.WARNING:
		adaptee.Warning(format, args...)
	case level.NOTICE:
		adaptee.Notice(format, args...)
	case level.INFO:
		adaptee.Info(format, args...)
	case level.DEBUG:
		adaptee.Debug(format, args...)
	}
}
