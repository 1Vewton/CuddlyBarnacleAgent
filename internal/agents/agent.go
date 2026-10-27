package agents

import (
	"context"
	"fmt"
	"sync"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
)

// Agent defines the data structure for storing agent data
type Agent struct {
	einoAgent    *adk.ChatModelAgent
	systemPrompt string
	AgentType    agenttypes.AgentType
}

// PromptConfig defines configurations for prompt for instance system prompt and instruction
type PromptConfig struct {
	Instruction  string
	SystemPrompt string
}

// AgentConfig defines the basic configurations for the agent
type AgentConfig struct {
	ModelName string
	AgentType agenttypes.AgentType
	TaskName  string
}

// ToolsList defines list of tools
type ToolsList struct {
	sync.RWMutex
	tools []tool.BaseTool
}

// NewToolsList creates new tools list
func NewToolsList() *ToolsList {
	return &ToolsList{}
}

// AddTool adds new tool to it, the data must be invokable tool
func (toolsList *ToolsList) AddTool(
	newTools []tool.InvokableTool,
) {
	toolsList.Lock()
	defer toolsList.Unlock()
	for _, tool := range newTools {
		toolsList.tools = append(
			toolsList.tools,
			tool,
		)
	}
}

// NewAgent creates new agent
func NewAgent(
	ctx context.Context,
	baseModel *openai.ChatModel,
	promptConfig *PromptConfig,
	agentConfig *AgentConfig,
	toolsList *ToolsList,
) (*Agent, error) {
	chatModelAgent, errCreateAgent := adk.NewChatModelAgent(
		ctx,
		&adk.ChatModelAgentConfig{
			Name: agentConfig.AgentType.ToString(),
			Description: fmt.Sprintf(
				"The agent that is responsible for %s",
				agentConfig.TaskName,
			),
			Instruction: promptConfig.Instruction,
			Model:       baseModel,
			ToolsConfig: adk.ToolsConfig{
				ToolsNodeConfig: compose.ToolsNodeConfig{
					Tools: toolsList.tools,
				},
			},
		},
	)
	if errCreateAgent != nil {
		return nil, errCreateAgent
	}
	return &Agent{
		einoAgent:    chatModelAgent,
		systemPrompt: promptConfig.SystemPrompt,
		AgentType:    agentConfig.AgentType,
	}, nil
}
