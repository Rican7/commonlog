/**
 * CommonLog
 *
 * Copyright © 2015 Trevor N. Suarez (Rican7)
 */

package iowriter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/Rican7/commonlog/adapter"
	"github.com/Rican7/commonlog/level"
)

func TestNew(t *testing.T) {
	l := New(ioutil.Discard)

	if _, ok := l.(adapter.LogAdapter); !ok {
		t.Errorf("%v doesn't satisfy the adapter.LogAdapter interface", l)
	}
}

func TestAdaptee(t *testing.T) {
	writer := ioutil.Discard

	adaptee := New(writer).Adaptee()

	if adaptee != writer {
		t.Errorf("Returned adaptee %T doesn't match expected %T", adaptee, writer)
	}
}

func TestLog(t *testing.T) {
	const testFormat = "My name is %s\n"
	const testArg = "Trevor"

	writer := &bytes.Buffer{}
	logger := New(writer)

	logger.Log(level.INFO, testFormat, testArg)

	loggedMessage := writer.String()
	expectedMessage := fmt.Sprintf(testFormat, testArg)

	if loggedMessage != expectedMessage {
		t.Errorf("Logged message %q doesn't match expected %q", loggedMessage, expectedMessage)
	}
}

func TestLogAppendsNewLineWhenMissing(t *testing.T) {
	const testFormat = "My name is %s"
	const testArg = "Trevor"

	writer := &bytes.Buffer{}
	logger := New(writer)

	logger.Log(level.INFO, testFormat, testArg)

	loggedMessage := writer.String()
	expectedMessage := fmt.Sprintf(testFormat, testArg) + "\n"

	if loggedMessage != expectedMessage {
		t.Errorf("Logged message %q doesn't match expected %q", loggedMessage, expectedMessage)
	}
}

func TestLogPanicsWithInvalidLogLevel(t *testing.T) {
	writer := &bytes.Buffer{}
	logger := New(writer)

	defer func() {
		if r := recover(); nil == r {
			t.Error("Expected to panic, but continued as normal")
		}
	}()

	logger.Log(level.LogLevel(255), "", "")
}
