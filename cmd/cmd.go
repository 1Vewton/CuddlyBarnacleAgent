package cmd

import (
	"os"

	"github.com/1Vewton/CuddlyBarnacleAgent/utils/logger"
)

var cmdLogger *logger.Logger = logger.NewLogger(
	"cmd",
	nil,
)

func init() {
	mainCMD.AddCommand(showConfigCMD)
}

// Execute executes the command
func Execute() {
	if err := mainCMD.Execute(); err != nil {
		cmdLogger.Error(err.Error())
		os.Exit(1)
	}
}
