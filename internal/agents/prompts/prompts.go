package prompts

import (
	"fmt"
)

const (
	tipsShowingCommand string = "The user wants you to follow the following tips:\n"
)

// GetTipsInfoCommand gets the command
func GetTipsInfoCommand(
	tips []string,
) string {
	rawString := ""
	for i, tip := range tips {
		rawString += fmt.Sprintf(
			"%d. %s\n",
			i+1,
			tip,
		)
	}
	return fmt.Sprintf(
		"%s%s",
		tipsShowingCommand,
		rawString,
	)
}
