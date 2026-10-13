package tools

import (
	"errors"
	"reflect"
	"testing"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/monitor"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/prompts"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/textresult"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

// TestReflectTypeToEinoType converts reflect type to eino type
func TestReflectTypeToEinoType(
	t *testing.T,
) {
	t.Parallel()
	stringType := reflect.String
	einoStringType := ReflectTypeToEinoType(stringType)
	if einoStringType != schema.String {
		t.Errorf(
			"got %s, expected %s",
			einoStringType,
			schema.String,
		)
	}
}

// TestCreateToolParamList tests the create tool params
func TestCreateToolParamList(
	t *testing.T,
) {
	t.Parallel()
	resultData := textresult.TextError{}
	paramPrompts := prompts.TextResultParam
	data, errCreate := CreateToolParams(
		paramPrompts,
		resultData,
	)
	if errCreate != nil {
		t.Error(errCreate)
	}
	_, ok := data["Line"]
	if !ok {
		t.Error("Expected to have Line field")
	}
	_, ok = data["Level"]
	if !ok {
		t.Error("Expected to have Level field")
	}
	_, ok = data["Type"]
	if !ok {
		t.Error("Expected to have Type field")
	}
	_, ok = data["Reason"]
	if !ok {
		t.Error("Expected to have Reason field")
	}
}

// TestInterface tests whether it implements tool interfacce
func TestInterface(t *testing.T) {
	t.Parallel()
	var resultToolTest any = &ProposeProblemTool{}
	_, ok := resultToolTest.(tool.InvokableTool)
	if !ok {
		t.Error("ProposeProblemTool does not implement invokable tool")
	}
}

// TestProcessMonitorError tests the process monitor error
func TestProcessMonitorError(t *testing.T) {
	t.Parallel()
	const errorText string = "This is a test error"
	monitor1 := monitor.NewAgentsRecord()
	tmpDir := t.TempDir()
	monitor1.Load(
		tmpDir,
		"test.json",
	)
	err := ProcessMonitorError(
		monitor1,
		agenttypes.FactualCheck,
		"test_tool",
		errors.New(
			errorText,
		),
		"{}",
	)
	if err == nil {
		t.Error(
			"expected to throw an error here",
		)
	}
	if err.Error() != errorText {
		t.Errorf(
			"expected returned error to be %s, got %s",
			errorText,
			err.Error(),
		)
	}
}
