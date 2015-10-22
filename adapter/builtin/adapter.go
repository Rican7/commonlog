/**
 * CommonLog
 *
 * Copyright © 2015 Trevor N. Suarez (Rican7)
 */

// Package builtin defines an adapter using the built-in "log" package
package builtin

import (
	"io"
	"log"

	"github.com/Rican7/commonlog"
	"github.com/Rican7/commonlog/adapter"
	"github.com/Rican7/commonlog/level"
)

/**
 * Types
 */

type adapted log.Logger

type logAdapter struct {
	commonlog.Logger
	adaptee *adapted
}

/**
 * Functions
 */

// New constructs a new instance by injecting an adaptee
func New(adaptee *log.Logger) adapter.LogAdapter {
	adapted := adapted(*adaptee)

	return &logAdapter{
		commonlog.NewLogger(&adapted),
		&adapted,
	}
}

// NewWithSetup constructs a new instance by building the adaptee
func NewWithSetup(out io.Writer, prefix string, flag int) adapter.LogAdapter {
	builtinLogger := log.New(out, prefix, flag)

	return New(builtinLogger)
}

// Adaptee gets the underyling adaptee
func (a *logAdapter) Adaptee() interface{} {
	return a.adaptee
}

// Log an error based on a specified level, a format, and a splat of arguments
func (a *adapted) Log(lvl level.LogLevel, format string, args ...interface{}) {
	// Validate the passed in level
	if ok, err := lvl.IsValid(); !ok {
		panic(err)
	}

	adaptee := (*log.Logger)(a)

	switch lvl {
	case level.EMERGENCY:
		// TODO: Handle this level once a generic log method is in the log.Logger
		fallthrough
	case level.ALERT:
		// TODO: Handle this level once a generic log method is in the log.Logger
		fallthrough
	case level.CRITICAL:
		// TODO: Handle this level once a generic log method is in the log.Logger
		adaptee.Fatalf(format, args...)
	case level.ERROR:
		// TODO: Handle this level once a generic log method is in the log.Logger
		adaptee.Panicf(format, args...)
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
		adaptee.Printf(format, args...)
	}
}
