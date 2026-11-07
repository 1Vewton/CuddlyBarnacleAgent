package tools

import (
	"context"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/monitor"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/prompts"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/task"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

// GetProblemsTool defines the tool for fetching problem
type GetProblemsTool struct {
	agentType   agenttypes.AgentType
	monitor     *monitor.AgentsRecord
	taskManager *task.Manager
	taskName    string
	toolName    string
}

// NewGetProblemsTool creates new tool for getting problems
func NewGetProblemsTool(
	agentType agenttypes.AgentType,
	taskName string,
	taskManager *task.Manager,
	monitor *monitor.AgentsRecord,
) *GetProblemsTool {
	return &GetProblemsTool{
		taskName:    taskName,
		agentType:   agentType,
		monitor:     monitor,
		taskManager: taskManager,
		toolName:    prompts.GetProblemsTool,
	}
}

// Info gives the information about this tool
func (t *GetProblemsTool) Info(
	ctx context.Context,
) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: t.toolName,
		Desc: prompts.GetProblemsTool,
	}, nil
}

// InvokableRun runs the get problems tool
func (t *GetProblemsTool) InvokableRun(
	ctx context.Context,
	argumentsInJSON string,
	opts ...tool.Option,
) (string, error) {
	err := t.monitor.NewCallWithSave(
		t.agentType,
		t.toolName,
	)
	if err != nil {
		err = ProcessMonitorError(
			t.monitor,
			t.agentType,
			t.toolName,
			err,
			argumentsInJSON,
		)
		return GiveFailResult(
			err,
		), err
	}
	t.taskManager.GetAllProblemsFor(
		t.taskName,
		t.agentType,
	)
	return "", nil
}
