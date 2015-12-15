/**
 * CommonLog
 *
 * Copyright © 2015 Trevor N. Suarez (Rican7)
 */

package builtin

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/Rican7/commonlog/adapter"
	"github.com/Rican7/commonlog/level"
)

/**
 * Variables
 */

var discardLogger = log.New(ioutil.Discard, "", 0)

func TestNew(t *testing.T) {
	l := New(discardLogger)

	if _, ok := l.(adapter.LogAdapter); !ok {
		t.Errorf("%v doesn't satisfy the adapter.LogAdapter interface", l)
	}
}

func TestAdaptee(t *testing.T) {
	logger := discardLogger

	adaptee := New(logger).Adaptee()

	if adaptee != logger {
		t.Errorf("Returned adaptee %T doesn't match expected %T", adaptee, logger)
	}
}

func TestLog(t *testing.T) {
	const testFormat = "My name is %s\n"
	const testArg = "Trevor"

	writer := &bytes.Buffer{}
	logger := New(log.New(writer, "", 0))

	logger.Log(level.INFO, testFormat, testArg)

	loggedMessage := writer.String()
	expectedMessage := fmt.Sprintf(testFormat, testArg)

	if loggedMessage != expectedMessage {
		t.Errorf("Logged message %q doesn't match expected %q", loggedMessage, expectedMessage)
	}
}

func TestLogPanicsWithInvalidLogLevel(t *testing.T) {
	writer := &bytes.Buffer{}
	logger := New(log.New(writer, "", 0))

	defer func() {
		if r := recover(); nil == r {
			t.Error("Expected to panic, but continued as normal")
		}
	}()

	logger.Log(level.LogLevel(255), "", "")
}
