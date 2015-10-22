/**
 * CommonLog
 *
 * Copyright © 2015 Trevor N. Suarez (Rican7)
 */

// Package glog defines an adapter using the google "glog" package
package glog

import (
	"github.com/Rican7/commonlog"
	"github.com/Rican7/commonlog/adapter"
	"github.com/Rican7/commonlog/level"
	"github.com/golang/glog"
)

/**
 * Types
 */

type adapted struct{}

type glogAdapter struct {
	commonlog.Logger
}

/**
 * Functions
 */

// New constructs a new instance
func New() adapter.LogAdapter {
	return &glogAdapter{
		commonlog.NewLogger(&adapted{}),
	}
}

// Adaptee gets the underyling adaptee
func (a *glogAdapter) Adaptee() interface{} {
	return nil
}

// Log an error based on a specified level, a format, and a splat of arguments
func (a *adapted) Log(lvl level.LogLevel, format string, args ...interface{}) {
	// Validate the passed in level
	if ok, err := lvl.IsValid(); !ok {
		panic(err)
	}

	switch lvl {
	case level.EMERGENCY:
		// TODO: Handle this level once a generic log method is in the glog.Logger
		fallthrough
	case level.ALERT:
		// TODO: Handle this level once a generic log method is in the glog.Logger
		fallthrough
	case level.CRITICAL:
		// TODO: Handle this level once a generic log method is in the glog.Logger
		glog.Fatalf(format, args...)
	case level.ERROR:
		glog.Errorf(format, args...)
	case level.WARNING:
		glog.Warningf(format, args...)
	case level.NOTICE:
		// TODO: Handle this level once a generic log method is in the glog.Logger
		fallthrough
	case level.INFO, level.DEBUG:
		glog.Infof(format, args...)
	}
}
