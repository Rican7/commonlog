/**
 * CommonLog
 *
 * Copyright © 2014 Trevor N. Suarez (Rican7)
 */

package level

import (
	"strings"
)

/**
 * Types
 */

type LogLevel uint8

type InvalidLogLevelError struct {
	InvalidValue *LogLevel
	InvalidName  *string
}

/**
 * Constants
 */

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

// Map a string name of our log levels to our constants
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

// Get a log level value by a string name
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

// Check if a log level is valid based on the standard defined log levels
func (l *LogLevel) IsValid() (bool, error) {
	if int(*l) >= len(logLevels) {
		return false, &InvalidLogLevelError{InvalidValue: l}
	}

	return true, nil
}

// Get the name of a given log level constant
func (l *LogLevel) String() string {
	if valid, err := l.IsValid(); !valid {
		return err.Error()
	}

	return logLevels[*l]
}

// Define our Error interface handler for our InvalidLogLevelError
func (e InvalidLogLevelError) Error() string {
	if nil != e.InvalidValue {
		return "Invalid log level constant. Must be out of range."

	} else if nil != e.InvalidName {
		return "Invalid log level name. No log level exists with that given name."
	}

	return ""
}
