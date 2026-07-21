package tips

import (
	"fmt"

	"github.com/fatih/color"
)

// ShowConfigTips shows the tips for altering configuration
func ShowConfigTips() {
	top := color.New(color.FgBlue, color.Bold)
	top.Println(
		`
The following contents will show the name of fields of configuration and the definitions of the fields
		`,
	)
	fmt.Println("LLMModelName: The name of the LLM model you want to use")
}
