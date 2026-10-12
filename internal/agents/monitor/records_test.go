package monitor

import (
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
	// Check data
	totalCalls := newRecord.GetTotalCalls(
		agenttypes.FactualCheck,
	)
	if totalCalls != 3 {
		t.Errorf(
			"expected total calls to be %d, got %d",
			3,
			totalCalls,
		)
	}
}
