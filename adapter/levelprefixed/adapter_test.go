/**
 * CommonLog
 *
 * Copyright © 2015 Trevor N. Suarez (Rican7)
 */

package levelprefixed

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/Rican7/commonlog/adapter"
	"github.com/Rican7/commonlog/level"
)

/**
 * Mocks
 */

type levelLogger struct {
	buffer bytes.Buffer
}

func (l *levelLogger) Log(lvl level.LogLevel, msg string, args ...interface{}) {
	l.buffer.WriteString(fmt.Sprintf(msg, args...))
}

func TestNew(t *testing.T) {
	l := New(&levelLogger{})

	if _, ok := l.(adapter.LogAdapter); !ok {
		t.Errorf("%v doesn't satisfy the adapter.LogAdapter interface", l)
	}
}

func TestAdaptee(t *testing.T) {
	logger := &levelLogger{}

	adaptee := New(logger).Adaptee()

	if adaptee != logger {
		t.Errorf("Returned adaptee %T doesn't match expected %T", adaptee, logger)
	}
}

func TestLog(t *testing.T) {
	const testFormat = "My name is %s"
	const testArg = "Trevor"
	const testLevel = level.INFO

	adaptee := &levelLogger{}
	logger := New(adaptee)

	logger.Log(testLevel, testFormat, testArg)

	loggedMessage := adaptee.buffer.String()
	expectedMessage := fmt.Sprintf(testFormat, testArg)

	if !strings.Contains(loggedMessage, expectedMessage) {
		t.Errorf("Logged message %q doesn't contain expected %q", loggedMessage, expectedMessage)
	}

	if !strings.HasPrefix(loggedMessage, testLevel.String()) {
		t.Errorf("Logged message %q doesn't start with expected %q", loggedMessage, testLevel.String())
	}
}

func TestLogPanicsWithInvalidLogLevel(t *testing.T) {
	logger := New(&levelLogger{})

	defer func() {
		if r := recover(); nil == r {
			t.Error("Expected to panic, but continued as normal")
		}
	}()

	logger.Log(level.LogLevel(255), "", "")
}
