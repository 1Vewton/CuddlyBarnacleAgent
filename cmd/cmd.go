package cmd

import (
	"os"

	"github.com/1Vewton/CuddlyBarnacleAgent/utils/logger"
	"github.com/spf13/cobra"
)

var cmdLogger *logger.Logger = logger.NewLogger(
	"cmd",
	nil,
)

// CBACommand runs the cmd program
var CBACommand *cobra.Command = &cobra.Command{
	Use:   "cba",
	Short: "CBA is a program for checking of document logic and other problems",
	Long: `
CBA is a program that can carry out checks for a document to see if there is
- Logic Error
- Spelling Error
- Wrong Fact
- Conflicts between previous documents
This program intended to act as the compiler of articles
	`,
	Run: func(cmd *cobra.Command, args []string) {
		Welcome()
	},
}

// Execute executes the command
func Execute() {
	if err := CBACommand.Execute(); err != nil {
		cmdLogger.Error(err.Error())
		os.Exit(1)
	}
}
