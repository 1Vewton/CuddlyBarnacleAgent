package textresult

import (
	"fmt"
)

// Levels defines the level of error
type Levels int

const (
	// NotSupported -1
	NotSupported Levels = -1
	// Warning defines the warning level.
	// 0
	Warning Levels = iota
	// Error defines the error level.
	// 1
	Error
)

// ToInt converts Levels to int
func (level Levels) ToInt() (int, error) {
	switch level {
	case Warning:
		return 0, nil
	case Error:
		return 1, nil
	case NotSupported:
		return -1, nil
	default:
		return -1, fmt.Errorf(
			"%d does not supported",
			level,
		)
	}
}

// ToLevels converts int to Levels
func ToLevels(
	integer int,
) (Levels, error) {
	switch integer {
	case 0:
		return Warning, nil
	case 1:
		return Error, nil
	case -1:
		return NotSupported, nil
	default:
		return NotSupported, fmt.Errorf(
			"Level %d not invalid, 0: Warning, 1: Error",
			integer,
		)
	}
}
