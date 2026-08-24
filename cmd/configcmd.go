package cmd

import (
	"github.com/1Vewton/CuddlyBarnacleAgent/cmd/cmdinterface"
	"github.com/1Vewton/CuddlyBarnacleAgent/cmd/tips"
	"github.com/1Vewton/CuddlyBarnacleAgent/pkg/config/ini"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// showConfigCMD shows the configuration
var showConfigCMD *cobra.Command = &cobra.Command{
	Use:   "showConfig",
	Short: "showConfig shows the configurations of the program",
	Long: `
Shows the configurations of following processes: 
- LLM
- Embeddings
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ini.IniConfig.ShowAllConfig()
	},
}

// setConfigCMD sets the configuration
var setConfigCMD *cobra.Command = &cobra.Command{
	Use:   "setConfig",
	Short: "setConfig sets the fields of the configuration",
	Long: `
The setConfig command can set the config through inputting the field first and input the new value for the inputted field.
The field name and the new value should be separated by space when inputting. 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		errorPrinter := color.New(
			color.FgRed,
			color.Bold,
		)
		tips.ShowConfigTips()
		// Input the field name and new value
		var targetField string
		var newValue string
		err := cmdinterface.ReadLine(
			"Input the field you want to change and the new value:",
			&targetField,
			&newValue,
		)
		if err != nil {
			cmdLogger.Error(err.Error())
			errorPrinter.Println(err.Error())
			return
		}
		// Set config
		err = ini.IniConfig.SetAnyConfig(
			targetField,
			newValue,
		)
		if err != nil {
			cmdLogger.Error(err.Error())
			errorPrinter.Println(err.Error())
			return
		}
	},
}
