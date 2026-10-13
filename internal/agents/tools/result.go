package tools

import (
	"context"
	"encoding/json"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/monitor"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/prompts"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/task"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/textresult"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

// ProposeProblemTool provides result
type ProposeProblemTool struct {
	taskName     string
	agentType    agenttypes.AgentType
	agentMonitor *monitor.AgentsRecord
	taskManager  *task.Manager
	toolName     string
}

// NewProposeProblemTool creates new ProposeProblemTool
func NewProposeProblemTool(
	taskName string,
	agentType agenttypes.AgentType,
	agentMonitor *monitor.AgentsRecord,
	taskManager *task.Manager,
) *ProposeProblemTool {
	return &ProposeProblemTool{
		taskName:     taskName,
		agentType:    agentType,
		agentMonitor: agentMonitor,
		taskManager:  taskManager,
		toolName:     prompts.ProposeProblemToolName,
	}
}

// Info gives information for this tool
func (t *ProposeProblemTool) Info(
	ctx context.Context,
) (*schema.ToolInfo, error) {
	expSchema := textresult.TextError{}
	singleParamSchema, err := CreateToolParams(
		prompts.TextResultParam,
		expSchema,
	)
	if err != nil {
		return nil, err
	}
	paramSchema := CreateListParam(
		singleParamSchema,
		prompts.TextErrorDescription,
		true,
		prompts.TextErrorArrayDescription,
	)
	return &schema.ToolInfo{
		Name: t.toolName,
		Desc: "Propose problems you find in the article",
		ParamsOneOf: schema.NewParamsOneOfByParams(
			map[string]*schema.ParameterInfo{
				"Array": paramSchema,
				"OverWriteOriginalData": NewParam(
					schema.Boolean,
					prompts.OverwriteDataDescription,
					true,
				),
			},
		),
	}, nil
}

// InvokableRun runs the propose problem tool
func (t *ProposeProblemTool) InvokableRun(
	ctx context.Context,
	argumentsInJSON string,
	opts ...tool.Option,
) (string, error) {
	t.agentMonitor.NewCall(
		t.agentType,
		t.toolName,
	)
	var inputs []textresult.TextError
	err := json.Unmarshal(
		[]byte(argumentsInJSON),
		&inputs,
	)
	if err != nil {
		return "", err
	}
	storedProblems, err := textresult.MultipleRawToProcessed(
		inputs,
		t.agentType,
	)
	if err != nil {
		return "", err
	}
	t.taskManager.AddQuestion(
		t.taskName,
		storedProblems,
	)
	return "", nil
}
