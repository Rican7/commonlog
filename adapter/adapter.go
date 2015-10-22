/**
 * CommonLog
 *
 * Copyright © 2015 Trevor N. Suarez (Rican7)
 */

package adapter

import (
	"github.com/Rican7/commonlog"
)

/**
 * Types
 */

// LogAdapter is an interface defining a logger that is built from an underlying
// adapter that can be accessed directly
type LogAdapter interface {
	commonlog.Logger

	Adaptee() interface{}
}
