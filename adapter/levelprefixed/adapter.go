/**
 * CommonLog
 *
 * Copyright © 2015 Trevor N. Suarez (Rican7)
 */

// Package levelprefixed defines a delegate adapter that prefixes messages with
// the log's severity level
package levelprefixed

import (
	"fmt"

	"github.com/Rican7/commonlog"
	"github.com/Rican7/commonlog/adapter"
	"github.com/Rican7/commonlog/level"
)

/**
 * Constants
 */

// preformattedLevelFormat is the level format before final formatting
const preformattedLevelFormat = "%%-%ds"

// levelDelimiter defines the delimiter used for the level prefix
const levelDelimiter = ":"

/**
 * Types
 */

type logger struct {
	commonlog.Logger

	prefix string
}

/**
 * Variables
 */

// maxLevelNameLength defines the maximum length of any level name
var maxLevelNameLength uint8

// levelFormat is the format of the level added to the message
var levelFormat string

/**
 * Functions
 */

// init initializes the package
func init() {
	// Initialize our maxLevelNameLength
	for _, level := range level.All() {
		lvlLength := uint8(len(level.String()))

		if lvlLength > maxLevelNameLength {
			maxLevelNameLength = lvlLength
		}
	}

	// Initialize our levelFormat
	levelFormat = fmt.Sprintf(preformattedLevelFormat, maxLevelNameLength)
}

// New constructs a new instance by injecting an adaptee
func New(adaptee commonlog.Logger) adapter.LogAdapter {
	return &logger{
		Logger: adaptee,
	}
}

// NewWithPrefix constructs a new instance by injecting an adaptee and setting a prefix
func NewWithPrefix(adaptee commonlog.Logger, prefix string) adapter.LogAdapter {
	return &logger{
		Logger: adaptee,
		prefix: prefix,
	}
}

// Adaptee gets the underyling adaptee
func (l *logger) Adaptee() interface{} {
	return l.Logger
}

// Log an error based on a specified level, a format, and a splat of arguments
func (l *logger) Log(lvl level.LogLevel, format string, args ...interface{}) {
	// Validate the passed in level
	if ok, err := lvl.IsValid(); !ok {
		panic(err)
	}

	// Prepend our level format
	format = "%s" + levelFormat + format

	// Define our level string with our delimiter
	lvlString := lvl.String() + levelDelimiter

	// Prepend the prefix and level string to our formatter args
	args = append([]interface{}{l.prefix, lvlString}, args...)

	l.Logger.Log(lvl, format, args...)
}
