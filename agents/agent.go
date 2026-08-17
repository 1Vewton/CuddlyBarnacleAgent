package agents

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/adk"
)

// Agent defines the data structure for storing agent data
type Agent struct {
	einoAgent    *adk.ChatModelAgent
	systemPrompt string
}

// NewAgent creates new agent
func NewAgent(
	ctx context.Context,
	baseModel *openai.ChatModel,
	name string,
	job string,
	instruction string,
	systemPrompt string,
) (*Agent, error) {
	chatModelAgent, errCreateAgent := adk.NewChatModelAgent(
		ctx,
		&adk.ChatModelAgentConfig{
			Name: name,
			Description: fmt.Sprintf(
				"The agent that is responsible for %s",
				job,
			),
			Instruction: instruction,
			Model:       baseModel,
		},
	)
	if errCreateAgent != nil {
		return nil, errCreateAgent
	}
	return &Agent{
		einoAgent:    chatModelAgent,
		systemPrompt: systemPrompt,
	}, nil
}
