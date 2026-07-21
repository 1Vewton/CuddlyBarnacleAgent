package cmd

import (
	"github.com/spf13/cobra"
)

// mainCMD runs the cmd program
var mainCMD *cobra.Command = &cobra.Command{
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
