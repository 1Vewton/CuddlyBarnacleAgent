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
	fmt.Println("LLMAPIKey: The API key of the LLM model you want to use")
	fmt.Println("LLMBaseURL: The base url of the provider of the LLM you want to use")
	fmt.Println("EmbeddingModelName: The name of the Embedding moddel you want to use")
	fmt.Println("EmbeddingModelAPIKey: The API Key of the Embedding moddel you want to use")
	fmt.Println("EmbeddingModelBaseURL: The base url of the provider of the Embedding moddel you want to use")
}
