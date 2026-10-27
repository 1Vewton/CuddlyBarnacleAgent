package tools

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/monitor"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/prompts"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/task"
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

// TestShowError tests the error showing
func TestShowError(t *testing.T) {
	t.Parallel()
	const expectedString string = `{"crucial_error": "ttsc"}`
	resultString := ShowError(
		errors.New("ttsc"),
	)
	if resultString != expectedString {
		t.Errorf(
			"expected %s, got %s",
			expectedString,
			resultString,
		)
	}
	const expectedString2 string = `{"crucial_error": "unknown error occurs"}`
	resultString2 := ShowError(
		nil,
	)
	if resultString2 != expectedString2 {
		t.Errorf(
			"expected %s, got %s",
			expectedString2,
			resultString2,
		)
	}
}

// TestGiveResult tests the result giving
func TestGiveResult(t *testing.T) {
	t.Parallel()
	// Give success result
	defintion := "test definition"
	const value int = 0
	expectedResult := &BaseProcessResult{
		Success: true,
		Message: "tool execution succeed",
		Data: &DataWithDefinition{
			Value:      value,
			Definition: &defintion,
		},
	}
	expectedInfo, err := json.Marshal(expectedResult)
	if err != nil {
		t.Fatal(err)
	}
	resultInfo := GiveSuccessResult(
		&defintion,
		value,
	)
	if string(expectedInfo) != resultInfo {
		t.Errorf(
			"expectec %s, got %s",
			string(expectedInfo),
			resultInfo,
		)
	}
	// Give failed result
	testErr := errors.New(
		"this is just a test",
	)
	expectedErrResult := &BaseProcessResult{
		Success: false,
		Message: testErr.Error(),
		Data:    nil,
	}
	byteData, err := json.Marshal(expectedErrResult)
	if err != nil {
		t.Fatal(err)
	}
	failedResult := GiveFailResult(testErr)
	if string(byteData) != failedResult {
		t.Fatalf(
			"expected %s, got %s",
			string(byteData),
			failedResult,
		)
	}
}

// TestProposeProblemTool tests the tool to propose problem
func TestProposeProblemTool(t *testing.T) {
	t.Parallel()
	// Initialize data needed
	ctx := t.Context()
	const taskName string = "taskA"
	tmpDir := t.TempDir()
	testMonitor := monitor.NewAgentsRecord()
	testMonitor.Load(
		tmpDir,
		"monitor.json",
	)
	testManager := task.NewRawManager()
	err := testManager.Load(
		tmpDir,
		"manager.json",
	)
	if err != nil {
		t.Fatal(err)
	}
	err = testManager.AddTask(
		taskName,
		"test",
	)
	if err != nil {
		t.Fatal(err)
	}
	// Create tool
	proposeProblemTool := NewProposeProblemTool(
		taskName,
		agenttypes.FactualCheck,
		testMonitor,
		testManager,
	)
	// Checks if this tool is a invokable tool
	var testTool interface{} = proposeProblemTool
	_, ok := testTool.(tool.InvokableTool)
	if !ok {
		t.Errorf(
			"it is not an invokable tool",
		)
	}
	// inputs
	level, err := textresult.Error.ToInt()
	if err != nil {
		t.Fatal(err)
	}
	testInputs := []textresult.TextError{
		{
			Level:  level,
			Type:   textresult.UncategorizedError.ToInt(),
			Line:   100,
			Reason: "test1",
		},
		{
			Level:  level,
			Type:   textresult.UncategorizedError.ToInt(),
			Line:   102,
			Reason: "test2",
		},
	}
	byteData, err := json.Marshal(
		testInputs,
	)
	if err != nil {
		t.Fatal(err)
	}
	result, err := proposeProblemTool.InvokableRun(
		ctx,
		string(byteData),
	)
	if err != nil {
		t.Fatal(err)
	}
	var proposeProblemResult *BaseProcessResult
	err = json.Unmarshal(
		[]byte(result),
		&proposeProblemResult,
	)
	if err != nil {
		t.Fatalf(
			"error occurred: %s, data: %s",
			err.Error(),
			result,
		)
	}
	if !proposeProblemResult.Success {
		t.Fatal(
			"tool execution failed",
		)
	}
	totalCalls := testMonitor.GetTotalCalls(
		agenttypes.FactualCheck,
	)
	if totalCalls != 1 {
		t.Fatalf(
			"expected total calls to be %d, got %d",
			1,
			totalCalls,
		)
	}
	successCalls := testMonitor.GetSuccessfulCalls(
		agenttypes.FactualCheck,
	)
	if successCalls != 1 {
		t.Fatalf(
			"expected the successful calls count to be %d, got %d",
			1,
			successCalls,
		)
	}
	failedCalls := testMonitor.GetFailedCalls(
		agenttypes.FactualCheck,
	)
	if failedCalls != 0 {
		t.Errorf(
			"expected the failed calls count to be %d, got %d",
			0,
			failedCalls,
		)
	}
}
