//go:build !test

package settings

import (
	"fmt"

	"github.com/joho/godotenv"
)

// Read the .env file
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		configLogger.Error(
			fmt.Sprintf(
				"Env reading failed due to %s",
				err,
			),
		)
	}
}
