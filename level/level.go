/**
 * CommonLog
 *
 * Copyright © 2015 Trevor N. Suarez (Rican7)
 */

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

/**
 * Variables
 */

// logLevels is a map of log level constants to their string names
var logLevels = []string{
	EMERGENCY: "EMERGENCY",
	ALERT:     "ALERT",
	CRITICAL:  "CRITICAL",
	ERROR:     "ERROR",
	WARNING:   "WARNING",
	NOTICE:    "NOTICE",
	INFO:      "INFO",
	DEBUG:     "DEBUG",
}

/**
 * Functions
 */

// NewLogLevel gets a log level value by a string name
func NewLogLevel(name string) (LogLevel, error) {
	// Cleanup the input
	name = strings.TrimSpace(name)

	for level, levelName := range logLevels {
		if strings.EqualFold(name, levelName) {
			return LogLevel(level), nil
		}
	}

	return ^LogLevel(0), &InvalidLogLevelError{InvalidName: &name}
}

// IsValid checks if a log level is valid based on the standard defined levels
func (l *LogLevel) IsValid() (bool, error) {
	if int(*l) >= len(logLevels) {
		return false, &InvalidLogLevelError{InvalidValue: l}
	}

	return true, nil
}

// String gets the name of a given log level constant
func (l *LogLevel) String() string {
	if valid, err := l.IsValid(); !valid {
		return err.Error()
	}

	return logLevels[*l]
}

// Error satisfies the error interface by returning a string message
func (e InvalidLogLevelError) Error() string {
	if nil != e.InvalidValue {
		return "Invalid log level constant. Must be out of range."

	} else if nil != e.InvalidName {
		return "Invalid log level name. No log level exists with that given name."
	}

	return ""
}
