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
	record.NewCall()
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
	tmpDir := t.TempDir()
	newRecord := NewAgentsRecord()
	err := newRecord.Load(
		tmpDir,
		"test.json",
	)
	if err != nil {
		t.Error(err)
	}
	// Initialize calls
	newRecord.NewCall(agenttypes.FactualCheck)
	newRecord.AddSuccessCall(
		agenttypes.FactualCheck,
		"test_tool",
	)
	newRecord.NewCall(agenttypes.FactualCheck)
	newRecord.AddSuccessCall(
		agenttypes.FactualCheck,
		"test_tool",
	)
	newRecord.NewCall(agenttypes.FactualCheck)
	newRecord.AddFailureRecord(
		agenttypes.FactualCheck,
		"test_tool",
	)
	newRecord.NewCall(agenttypes.GrammaticalCheck)
	newRecord.AddFailureRecordByDetail(
		agenttypes.GrammaticalCheck,
		"test_tool",
		errors.New(
			"this is a test error",
		),
		"Hello, World!",
	)
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
}
