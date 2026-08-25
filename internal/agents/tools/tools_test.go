package tools

import (
	"reflect"
	"testing"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/prompts"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/textresult"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

// TestReflectTypeToEinoType converts reflect type to eino type
func TestReflectTypeToEinoType(
	t *testing.T,
) {
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
	var resultToolTest any = NewProvideResultTool("114514")
	_, ok := resultToolTest.(tool.InvokableTool)
	if !ok {
		t.Error("ProvideResultTool does not implement invokable tool")
	}
}
