package tools

import (
	"reflect"
	"testing"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/prompts"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/textresult"
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
	_, errCreate := CreateToolParams(
		paramPrompts,
		resultData,
	)
	if errCreate != nil {
		t.Error(errCreate)
	}
}
