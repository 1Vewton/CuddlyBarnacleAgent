package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/monitor"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/prompts"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/task"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/textresult"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/operations"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

// ProposeProblemToolInput defines the input for tool that provides result
type ProposeProblemToolInput struct {
	Array                 []textresult.TextError
	OverWriteOriginalData bool
}

// ProposeProblemTool provides result
type ProposeProblemTool struct {
	taskName     string
	agentType    agenttypes.AgentType
	agentMonitor *monitor.AgentsRecord
	taskManager  *task.Manager
	toolName     string
	errorTypes   *textresult.AllErrorTypes
	fields       []textresult.ErrorField
}

// NewProposeProblemTool creates new ProposeProblemTool
func NewProposeProblemTool(
	taskName string,
	agentType agenttypes.AgentType,
	agentMonitor *monitor.AgentsRecord,
	taskManager *task.Manager,
	errorTypes *textresult.AllErrorTypes,
	fields []textresult.ErrorField,
) *ProposeProblemTool {
	return &ProposeProblemTool{
		taskName:     taskName,
		agentType:    agentType,
		agentMonitor: agentMonitor,
		taskManager:  taskManager,
		toolName:     prompts.ProposeProblemToolName,
		errorTypes:   errorTypes,
		fields:       fields,
	}
}

// Info gives information for this tool
func (t *ProposeProblemTool) Info(
	ctx context.Context,
) (*schema.ToolInfo, error) {
	expSchema := textresult.TextError{}
	errorTypeDescription, err := prompts.GetTypeDescription(
		t.errorTypes,
		t.fields,
	)
	if err != nil {
		return nil, err
	}
	paramDescription := prompts.GetTextResultParam(
		errorTypeDescription,
	)
	singleParamSchema, err := CreateToolParams(
		paramDescription,
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
		Desc: prompts.ProposeProblemToolDescription,
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
	err := t.agentMonitor.NewCallWithSave(
		t.agentType,
		t.toolName,
	)
	if err != nil {
		err = ProcessMonitorError(
			t.agentMonitor,
			t.agentType,
			t.toolName,
			err,
			argumentsInJSON,
		)
		return GiveFailResult(
			err,
		), err
	}
	var inputs *ProposeProblemToolInput
	err = json.Unmarshal(
		[]byte(argumentsInJSON),
		&inputs,
	)
	if err != nil {
		err = ProcessMonitorError(
			t.agentMonitor,
			t.agentType,
			t.toolName,
			err,
			argumentsInJSON,
		)
		return GiveFailResult(
			err,
		), err
	}
	storedProblems, err := textresult.MultipleRawToProcessed(
		inputs.Array,
		t.agentType,
		t.errorTypes,
	)
	if err != nil {
		err = ProcessMonitorError(
			t.agentMonitor,
			t.agentType,
			t.toolName,
			err,
			argumentsInJSON,
		)
		return GiveFailResult(
			err,
		), err
	}
	for _, problem := range storedProblems {
		typeInfo := problem.Type
		if typeInfo != nil {
			typeField, err := t.errorTypes.GetFieldByErrorCode(
				typeInfo.ErrorCode,
			)
			if err != nil {
				return GiveFailResult(
					err,
				), err
			}
			isFieldIncluded := operations.IsIncluded(
				t.fields,
				typeField,
				func(f1 textresult.ErrorField, f2 textresult.ErrorField) bool {
					return f1 == f2
				},
			)
			if !isFieldIncluded {
				err := fmt.Errorf(
					"field %d does not supported",
					typeField,
				)
				return GiveFailResult(
					err,
				), err
			}
		}
	}
	err = t.taskManager.AddQuestion(
		t.agentType,
		t.taskName,
		storedProblems,
		inputs.OverWriteOriginalData,
	)
	if err != nil {
		err = ProcessMonitorError(
			t.agentMonitor,
			t.agentType,
			t.toolName,
			err,
			argumentsInJSON,
		)
		return GiveFailResult(
			err,
		), err
	}
	err = t.taskManager.SaveFile()
	if err != nil {
		err = ProcessMonitorError(
			t.agentMonitor,
			t.agentType,
			t.toolName,
			err,
			argumentsInJSON,
		)
		return GiveFailResult(
			err,
		), err
	}
	err = t.agentMonitor.AddSuccessCallWithSave(
		t.agentType,
		t.toolName,
	)
	if err != nil {
		return GiveFailResult(
			err,
		), err
	}
	return GiveSuccessResult(
		nil,
		nil,
	), nil
}
