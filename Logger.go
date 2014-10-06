package commonlog

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

type Logger interface {
	Emergency(string, ...interface{})
	Alert(string, ...interface{})
	Critical(string, ...interface{})
	Error(string, ...interface{})
	Warning(string, ...interface{})
	Notice(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})

	Log(LogLevel, string, ...interface{})
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

// Get the name of a given log level constant
func (l LogLevel) String() string {
	if int(l) >= len(logLevels) {
		return InvalidLogLevelError{InvalidValue: &l}.Error()
	}

	return logLevels[l]
}

// Get the name of a given log level constant
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

// Define our Error interface handler for our InvalidLogLevelError
func (e InvalidLogLevelError) Error() string {
	if nil != e.InvalidValue {
		return "Invalid log level constant. Must be out of range."

	} else if nil != e.InvalidName {
		return "Invalid log level name. No log level exists with that given name."
	}

	return ""
}
