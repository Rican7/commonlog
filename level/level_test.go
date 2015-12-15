/**
 * CommonLog
 *
 * Copyright © 2015 Trevor N. Suarez (Rican7)
 */

package level

import "testing"

func TestAll(t *testing.T) {
	all := All()

	if len(logLevelNames) != len(all) {
		t.Errorf("Length of all (%v) does not match expected %v", len(all), len(logLevelNames))
	}

	for _, lvl := range all {
		if _, ok := logLevelNames[lvl]; !ok {
			t.Errorf("%d is not a defined level", lvl)
		}
	}
}

func TestNewLogLevel(t *testing.T) {
	valid, err := NewLogLevel("emeRGEncy")

	if nil != err {
		t.Errorf("Failed to build a valid log level with error %q", err)
	}

	if valid != EMERGENCY {
		t.Errorf("Log level %v doesn't match expected %v", valid, EMERGENCY)
	}

	invalid, err := NewLogLevel("wat?")

	if nil == err {
		t.Error("Expected an error but got nil")
	}

	if invalid != errorValue {
		t.Errorf("Invalid value %d doesn't match expected %d", invalid, errorValue)
	}
}

func TestIsValid(t *testing.T) {
	const valid = EMERGENCY
	const invalid = errorValue

	isValidValid, err := valid.IsValid()

	if nil != err {
		t.Errorf("IsValid check failed with error %q", err)
	}

	if !isValidValid {
		t.Errorf("Valid level %d failed to validate", valid)
	}

	isInvalidValid, err := invalid.IsValid()

	if nil == err {
		t.Error("Expected an error but got nil")
	}

	if isInvalidValid {
		t.Errorf("Invalid level %d validated unexpectedly", invalid)
	}
}

func TestString(t *testing.T) {
	const valid = EMERGENCY
	const invalid = errorValue

	validString := valid.String()

	if validString != logLevelNames[valid] {
		t.Errorf("Valid string %q doesn't match expected %q", validString, logLevelNames[valid])
	}

	invalidString := invalid.String()

	if invalidString == logLevelNames[invalid] {
		t.Errorf("Invalid string %q unexpectedly matches valid %q", invalidString, logLevelNames[invalid])
	}
}

func TestError(t *testing.T) {
	invalidValue := errorValue
	invalidName := "wat?!"

	normalMessage := InvalidLogLevelError{}.Error()
	invalidValueMessage := InvalidLogLevelError{InvalidValue: &invalidValue}.Error()
	invalidNameMessage := InvalidLogLevelError{InvalidName: &invalidName}.Error()

	for _, msg := range []string{normalMessage, invalidValueMessage, invalidNameMessage} {
		if "" == msg {
			t.Error("Error message unexpectedly empty")
		}
	}

	if normalMessage == invalidValueMessage ||
		normalMessage == invalidNameMessage ||
		invalidValueMessage == invalidNameMessage {
		t.Error("Error message failed distinction")
	}
}
