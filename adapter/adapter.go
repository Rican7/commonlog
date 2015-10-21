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

type LogAdapter interface {
	commonlog.Logger

	Adaptee() interface{}
}
