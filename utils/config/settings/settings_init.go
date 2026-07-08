//go:build !test

package settings

import (
	"fmt"
)

// Read the .env file
func init() {
	err := Settings.Initialize(".env")
	if err != nil {
		configLogger.Error(
			fmt.Sprintf(
				"Env reading failed due to %s",
				err,
			),
		)
	}
}
