package tools

import (
	"context"
	"encoding/json"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/prompts"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/textresult"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

// ProvideResultTool provides result
type ProvideResultTool struct {
	taskName string
}

// NewProvideResultTool creates new ProvideResultTool
func NewProvideResultTool(
	taskName string,
) *ProvideResultTool {
	return &ProvideResultTool{
		taskName: taskName,
	}
}

// Info gives information for this tool
func (t *ProvideResultTool) Info(
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
		Name: prompts.ProvideResultToolName,
		Desc: "Provides result to the user",
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

// InvokableRun runs the provides result tool
func (t *ProvideResultTool) InvokableRun(
	ctx context.Context,
	argumentsInJSON string,
	opts ...tool.Option,
) (string, error) {
	var inputs []textresult.TextError
	errUnmarshal := json.Unmarshal(
		[]byte(argumentsInJSON),
		&inputs,
	)
	if errUnmarshal != nil {
		return "", nil
	}
	return "", nil
}
