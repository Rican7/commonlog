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

// DefaultDelimiter defines the default delimiter used for the level prefix
const DefaultDelimiter = ':'

// DefaultPadding defines the default padding used for the level prefix
const DefaultPadding = "\t"

/**
 * Types
 */

type delegate struct {
	commonlog.LevelLogger

	prefix    string
	delimiter rune
	padding   string
}

type logger struct {
	commonlog.Logger
	adaptee commonlog.LevelLogger
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
	levelFormat = fmt.Sprintf(
		preformattedLevelFormat,
		maxLevelNameLength+1, // Add one for the length of the single rune delimiter
	)
}

// New constructs a new instance by injecting an adaptee
func New(adaptee commonlog.LevelLogger) adapter.LogAdapter {
	return NewWithPrefix(adaptee, "")
}

// NewWithPrefix constructs a new instance by injecting an adaptee and setting a prefix
func NewWithPrefix(adaptee commonlog.LevelLogger, prefix string) adapter.LogAdapter {
	return NewWithCustomPresentation(adaptee, prefix, DefaultDelimiter, DefaultPadding)
}

// NewWithCustomPresentation constructs a new instance by injecting an adaptee
// and setting options to control the presentation of the level prefix
func NewWithCustomPresentation(adaptee commonlog.LevelLogger, prefix string, delimiter rune, padding string) adapter.LogAdapter {
	return &logger{
		Logger: commonlog.NewLogger(
			&delegate{
				LevelLogger: adaptee,
				prefix:      prefix,
				delimiter:   delimiter,
				padding:     padding,
			},
		),
		adaptee: adaptee,
	}
}

// Adaptee gets the underyling adaptee
func (l *logger) Adaptee() interface{} {
	return l.adaptee
}

// Log an error based on a specified level, a format, and a splat of arguments
func (l *delegate) Log(lvl level.LogLevel, format string, args ...interface{}) {
	// Validate the passed in level
	if ok, err := lvl.IsValid(); !ok {
		panic(err)
	}

	// Prepend our level format
	format = "%s" + levelFormat + "%s" + format

	// Define our level string with our delimiter
	lvlString := lvl.String() + string(l.delimiter)

	// Prepend the prefix and level string to our formatter args
	args = append([]interface{}{l.prefix, lvlString, l.padding}, args...)

	l.LevelLogger.Log(lvl, format, args...)
}
