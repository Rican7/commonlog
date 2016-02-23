/**
 * CommonLog
 *
 * Copyright © 2015 Trevor N. Suarez (Rican7)
 */

package commonlog

import (
	"reflect"
	"testing"

	"github.com/Rican7/commonlog/level"
)

/**
 * Mocks
 */

type levelLoggerFunc func(level.LogLevel, string, ...interface{})

func (f levelLoggerFunc) Log(l level.LogLevel, msg string, args ...interface{}) {
	f(l, msg, args)
}

type lastStateLogger struct {
	lastLevel   *level.LogLevel
	lastMessage *string
	lastArgs    *[]interface{}
}

func (s *lastStateLogger) Clear() {
	s = &lastStateLogger{}
}

func (s *lastStateLogger) Log(l level.LogLevel, msg string, args ...interface{}) {
	s.lastLevel = &l
	s.lastMessage = &msg
	s.lastArgs = &args
}

/**
 * Tests
 */

func TestNewLogger(t *testing.T) {
	ll := func(level.LogLevel, string, ...interface{}) {}

	logger := NewLogger(levelLoggerFunc(ll))

	if _, ok := logger.(Logger); !ok {
		t.Errorf("%v doesn't satisfy the Logger interface", logger)
	}
}

func TestLevelMethods(t *testing.T) {
	const testMessage = "What"
	testArgs := []interface{}{1, 2, 3}

	// Setup a map of our expected levels and matching method calls
	calls := map[level.LogLevel]func(Logger, string, ...interface{}){
		level.EMERGENCY: Logger.Emergency,
		level.ALERT:     Logger.Alert,
		level.CRITICAL:  Logger.Critical,
		level.ERROR:     Logger.Error,
		level.WARNING:   Logger.Warning,
		level.NOTICE:    Logger.Notice,
		level.INFO:      Logger.Info,
		level.DEBUG:     Logger.Debug,
	}

	stateLogger := &lastStateLogger{}
	logger := &common{stateLogger}

	for lvl, call := range calls {
		call(logger, testMessage, testArgs...)

		if *stateLogger.lastLevel != lvl {
			t.Errorf("Logged level %q does not match expected %q", stateLogger.lastLevel, lvl)
		}

		if *stateLogger.lastMessage != testMessage {
			t.Errorf("Logged message %q does not match expected %q", stateLogger.lastMessage, testMessage)
		}

		if !reflect.DeepEqual(*stateLogger.lastArgs, testArgs) {
			t.Errorf("Logged args %q does not match expected %q", stateLogger.lastArgs, testArgs)
		}
	}
}
