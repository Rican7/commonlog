/**
 * CommonLog
 *
 * Copyright © 2014 Trevor N. Suarez (Rican7)
 */

package glog

import (
	"../"
	"../../level"
	"github.com/golang/glog"
)

/**
 * Types
 */

type glogAdapter struct {
}

/**
 * Functions
 */

// Construct a new instance
func New() adapter.LogAdapter {
	return &glogAdapter{}
}

// Get the underyling adaptee
func (a *glogAdapter) Adaptee() interface{} {
	return nil
}

// Convenient alias for glogAdapter.Log()
func (a *glogAdapter) Emergency(format string, args ...interface{}) {
	a.Log(level.EMERGENCY, format, args)
}

// Convenient alias for glogAdapter.Log()
func (a *glogAdapter) Alert(format string, args ...interface{}) {
	a.Log(level.ALERT, format, args)
}

// Convenient alias for glogAdapter.Log()
func (a *glogAdapter) Critical(format string, args ...interface{}) {
	a.Log(level.CRITICAL, format, args)
}

// Convenient alias for glogAdapter.Log()
func (a *glogAdapter) Error(format string, args ...interface{}) {
	a.Log(level.ERROR, format, args)
}

// Convenient alias for glogAdapter.Log()
func (a *glogAdapter) Warning(format string, args ...interface{}) {
	a.Log(level.WARNING, format, args)
}

// Convenient alias for glogAdapter.Log()
func (a *glogAdapter) Notice(format string, args ...interface{}) {
	a.Log(level.NOTICE, format, args)
}

// Convenient alias for glogAdapter.Log()
func (a *glogAdapter) Info(format string, args ...interface{}) {
	a.Log(level.INFO, format, args)
}

// Convenient alias for glogAdapter.Log()
func (a *glogAdapter) Debug(format string, args ...interface{}) {
	a.Log(level.DEBUG, format, args)
}

/**
 * Log an error based on a specified level, a format, and a splat of arguments
 */
func (a *glogAdapter) Log(lvl level.LogLevel, format string, args ...interface{}) {
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
		glog.Fatalf(format, args)
	case level.ERROR:
		glog.Errorf(format, args)
	case level.WARNING:
		glog.Warningf(format, args)
	case level.NOTICE:
		// TODO: Handle this level once a generic log method is in the glog.Logger
		fallthrough
	case level.INFO, level.DEBUG:
		glog.Infof(format, args)
	}
}
