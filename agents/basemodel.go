package agents

import (
	"context"

	"github.com/cloudwego/eino-ext/components/model/openai"
)

// NewOpenAIBaseModel creats new base model using OpenAI-api compatible llm
func NewOpenAIBaseModel(
	ctx context.Context,
	APIKey string,
	baseURL string,
	model string,
) (*openai.ChatModel, error) {
	return openai.NewChatModel(
		ctx,
		&openai.ChatModelConfig{
			APIKey:  APIKey,
			BaseURL: baseURL,
			Model:   model,
		},
	)
}
