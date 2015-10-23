/**
 * CommonLog
 *
 * Copyright © 2015 Trevor N. Suarez (Rican7)
 */

// Package iowriter defines an adapter that uses the io.Writer interface
package iowriter

import (
	"fmt"
	"io"

	"github.com/Rican7/commonlog"
	"github.com/Rican7/commonlog/adapter"
	"github.com/Rican7/commonlog/level"
)

/**
 * Types
 */

type delegate struct {
	io.Writer
}

type logAdapter struct {
	commonlog.Logger
	adaptee *delegate
}

/**
 * Functions
 */

// New constructs a new instance by injecting an adaptee
func New(adaptee io.Writer) adapter.LogAdapter {
	adapted := &delegate{adaptee}

	return &logAdapter{
		commonlog.NewLogger(adapted),
		adapted,
	}
}

// Adaptee gets the underyling adaptee
func (a *logAdapter) Adaptee() interface{} {
	return a.adaptee.Writer
}

// Log an error based on a specified level, a format, and a splat of arguments
func (a *delegate) Log(lvl level.LogLevel, format string, args ...interface{}) {
	// Validate the passed in level
	if ok, err := lvl.IsValid(); !ok {
		panic(err)
	}

	buf := []byte(fmt.Sprintf(format, args...))

	if len(buf) == 0 || buf[len(buf)-1] != '\n' {
		buf = append(buf, '\n')
	}

	a.Write(buf)
}
