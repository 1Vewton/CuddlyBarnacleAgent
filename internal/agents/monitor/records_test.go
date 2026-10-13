package monitor

import (
	"errors"
	"testing"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
)

// TestSingleAgentRecord tests operations related to single agent record
func TestSingleAgentRecord(t *testing.T) {
	t.Parallel()
	record := NewSingleAgentRecord()
	record.NewCall(
		"test_tool",
	)
	if record.GetTotalCalls() != 1 {
		t.Errorf(
			"expected total calls to be %d, got %d",
			1,
			record.GetTotalCalls(),
		)
	}
}

// TestAgentsRecord tests the monitor for all agents
func TestAgentsRecord(t *testing.T) {
	t.Parallel()
	const toolName string = "test_tool"
	const fileName string = "test.json"
	tmpDir := t.TempDir()
	newRecord := NewAgentsRecord()
	err := newRecord.Load(
		tmpDir,
		fileName,
	)
	if err != nil {
		t.Error(err)
	}
	// Initialize calls
	newRecord.NewCall(
		agenttypes.FactualCheck,
		toolName,
	)
	newRecord.AddSuccessCall(
		agenttypes.FactualCheck,
		toolName,
	)
	newRecord.NewCall(
		agenttypes.FactualCheck,
		toolName,
	)
	newRecord.AddSuccessCall(
		agenttypes.FactualCheck,
		toolName,
	)
	newRecord.NewCall(
		agenttypes.FactualCheck,
		toolName,
	)
	newRecord.AddFailureCall(
		agenttypes.FactualCheck,
		toolName,
	)
	newRecord.NewCall(
		agenttypes.GrammaticalCheck,
		toolName,
	)
	newRecord.AddFailureCallByDetail(
		agenttypes.GrammaticalCheck,
		toolName,
		errors.New(
			"this is a test error",
		),
		"Hello, World!",
	)
	err = newRecord.SaveFile()
	if err != nil {
		t.Error(err)
	}
	// Check data
	totalCallsFactualCheck := newRecord.GetTotalCalls(
		agenttypes.FactualCheck,
	)
	if totalCallsFactualCheck != 3 {
		t.Errorf(
			"expected total calls to be %d, got %d",
			3,
			totalCallsFactualCheck,
		)
	}
	totalSuccessCallsFactualCheck := newRecord.GetSuccessfulCalls(
		agenttypes.FactualCheck,
	)
	if totalSuccessCallsFactualCheck != 2 {
		t.Errorf(
			"expected total successful calls to be %d, got %d",
			2,
			totalSuccessCallsFactualCheck,
		)
	}
	totalFailureCallsFactualCheck := newRecord.GetFailedCalls(
		agenttypes.FactualCheck,
	)
	if totalFailureCallsFactualCheck != 1 {
		t.Errorf(
			"expected total failed calls to be %d, got %d",
			1,
			totalSuccessCallsFactualCheck,
		)
	}
	newRecord2 := NewAgentsRecord()
	err = newRecord2.Load(
		tmpDir,
		fileName,
	)
	if err != nil {
		t.Error(err)
	}
	totalCallsFactualCheck = newRecord2.GetTotalCalls(
		agenttypes.FactualCheck,
	)
	if totalCallsFactualCheck != 3 {
		t.Errorf(
			"expected the total number of total calls to be %d, got %d",
			3,
			totalCallsFactualCheck,
		)
	}
	totalSuccessCallsFactualCheck = newRecord2.GetSuccessfulCalls(
		agenttypes.FactualCheck,
	)
	if totalSuccessCallsFactualCheck != 2 {
		t.Errorf(
			"expected the total number of successes to be %d, got %d",
			2,
			totalSuccessCallsFactualCheck,
		)
	}
	totalFailureCallsFactualCheck = newRecord2.GetFailedCalls(
		agenttypes.FactualCheck,
	)
	if totalFailureCallsFactualCheck != 1 {
		t.Errorf(
			"expected the total number of successes to be %d, got %d",
			1,
			totalFailureCallsFactualCheck,
		)
	}
	err = newRecord2.NewCallWithSave(
		agenttypes.FactualCheck,
		toolName,
	)
	if err != nil {
		t.Error(err)
	}
	err = newRecord2.AddSuccessCallWithSave(
		agenttypes.FactualCheck,
		toolName,
	)
	if err != nil {
		t.Error(err)
	}
	err = newRecord2.NewCallWithSave(
		agenttypes.FactualCheck,
		toolName,
	)
	if err != nil {
		t.Error(err)
	}
	err = newRecord2.AddFailureCallWithSave(
		agenttypes.FactualCheck,
		toolName,
	)
	if err != nil {
		t.Error(err)
	}
	newRecord3 := NewAgentsRecord()
	err = newRecord3.Load(
		tmpDir,
		fileName,
	)
	if err != nil {
		t.Error(err)
	}
	totalCalls := newRecord3.GetTotalCalls(
		agenttypes.FactualCheck,
	)
	if totalCalls != 5 {
		t.Errorf(
			"expected number of calls to be %d, got %d",
			5,
			totalCalls,
		)
	}
}
