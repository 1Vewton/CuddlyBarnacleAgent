package cmd

import (
	"github.com/1Vewton/CuddlyBarnacleAgent/utils/config/ini"
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
