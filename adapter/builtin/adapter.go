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

type delegate struct {
	*log.Logger
}

type logAdapter struct {
	commonlog.Logger
	adaptee *delegate
}

/**
 * Functions
 */

// New constructs a new instance by injecting an adaptee
func New(adaptee *log.Logger) adapter.LogAdapter {
	adapted := &delegate{adaptee}

	return &logAdapter{
		commonlog.NewLogger(adapted),
		adapted,
	}
}

// NewWithSetup constructs a new instance by building the adaptee
func NewWithSetup(out io.Writer, prefix string, flag int) adapter.LogAdapter {
	builtinLogger := log.New(out, prefix, flag)

	return New(builtinLogger)
}

// Adaptee gets the underyling adaptee
func (a *logAdapter) Adaptee() interface{} {
	return a.adaptee.Logger
}

// Log an error based on a specified level, a format, and a splat of arguments
func (a *delegate) Log(lvl level.LogLevel, format string, args ...interface{}) {
	// Validate the passed in level
	if ok, err := lvl.IsValid(); !ok {
		panic(err)
	}

	switch lvl {
	case level.EMERGENCY:
		fallthrough
	case level.ALERT:
		fallthrough
	case level.CRITICAL:
		a.Fatalf(format, args...)
	case level.ERROR:
		a.Panicf(format, args...)
	case level.WARNING:
		fallthrough
	case level.NOTICE:
		fallthrough
	case level.INFO:
		fallthrough
	case level.DEBUG:
		a.Printf(format, args...)
	}
}
