/**
 * CommonLog
 *
 * Copyright © 2014 Trevor N. Suarez (Rican7)
 */

package adapter

import (
	"../"
)

/**
 * Types
 */

type LogAdapter interface {
	commonlog.Logger

	Adaptee() interface{}
}
