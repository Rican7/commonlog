/**
 * CommonLog
 *
 * Copyright © 2015 Trevor N. Suarez (Rican7)
 */

// Package level defines the standard levels and their validations
package level

import (
	"strings"
)

/**
 * Types
 */

// A LogLevel defines a logging severity level, akin to the severity levels used
// in the Syslog standard. (https://tools.ietf.org/html/rfc5424)
type LogLevel uint8

// InvalidLogLevelError is an error representing an invalid log level occurrence
type InvalidLogLevelError struct {
	InvalidValue *LogLevel
	InvalidName  *string
}

/**
 * Constants
 */

// The log severity levels, as defined in RFC 5424 (Section 6.2.1):
// https://tools.ietf.org/html/rfc5424#section-6.2.1
const (
	EMERGENCY LogLevel = iota
	ALERT
	CRITICAL
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

// errorValue defines the value returned for failed log level creations
const errorValue = ^LogLevel(0)

/**
 * Variables
 */

// logLevelNames is a map of log level constants to their string names
var logLevelNames = map[LogLevel]string{
	EMERGENCY: "EMERGENCY",
	ALERT:     "ALERT",
	CRITICAL:  "CRITICAL",
	ERROR:     "ERROR",
	WARNING:   "WARNING",
	NOTICE:    "NOTICE",
	INFO:      "INFO",
	DEBUG:     "DEBUG",
}

// logLevelNamesInverse is an inverted map of logLevelNames
var logLevelNamesInverse map[string]LogLevel

/**
 * Functions
 */

// init initializes the package
func init() {
	// Initialize our inverted log level name map
	logLevelNamesInverse = make(map[string]LogLevel, len(logLevelNames))
	for level, name := range logLevelNames {
		logLevelNamesInverse[name] = level
	}
}

// All returns an array of the standard defined log levels
func All() []LogLevel {
	all := make([]LogLevel, len(logLevelNames))
	for k := range logLevelNames {
		all[k] = k
	}

	return all
}

// NewLogLevel gets a log level value by a string name
func NewLogLevel(name string) (LogLevel, error) {
	// Cleanup the input
	name = strings.ToUpper(strings.TrimSpace(name))

	level, ok := logLevelNamesInverse[name]

	if !ok {
		return errorValue, &InvalidLogLevelError{InvalidName: &name}
	}

	return level, nil
}

// IsValid checks if a log level is valid based on the standard defined levels
func (l LogLevel) IsValid() (bool, error) {
	if _, ok := logLevelNames[l]; !ok {
		return false, &InvalidLogLevelError{InvalidValue: &l}
	}

	return true, nil
}

// String gets the name of a given log level constant
func (l LogLevel) String() string {
	if valid, err := l.IsValid(); !valid {
		return err.Error()
	}

	return logLevelNames[l]
}

// Error satisfies the error interface by returning a string message
func (e InvalidLogLevelError) Error() string {
	msg := "Unknown error"

	if nil != e.InvalidValue {
		msg = "Invalid log level constant. Must be out of range."

	} else if nil != e.InvalidName {
		msg = "Invalid log level name. No log level exists with that given name."
	}

	return msg
}
