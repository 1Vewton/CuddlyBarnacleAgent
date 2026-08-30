package task

import (
	"testing"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/textresult"
)

// TestTask tests the functions related to
func TestTask(t *testing.T) {
	testData := []*textresult.StoredTextError{
		{
			Level:    textresult.Error,
			Type:     textresult.UncategorizedError,
			Line:     114514,
			Reason:   "1919810",
			Proposer: agenttypes.FactualCheck,
		},
		{
			Level:    textresult.Error,
			Type:     textresult.UncategorizedError,
			Line:     1919810,
			Reason:   "1919810",
			Proposer: agenttypes.GrammaticalCheck,
		},
	}
	newTask := NewTask("TestTask")
	newTask.AddMultipleProblems(testData)
	result := newTask.GetAllProblemsFor(agenttypes.GrammaticalCheck)
	length := len(result)
	if length != 1 {
		t.Errorf(
			"Expected length of return list to be %d, got %d",
			1,
			length,
		)
	}
	if result[0].Line != 1919810 {
		t.Errorf(
			"Expect the element on position 0 has line field %d, got %d",
			1919810,
			result[0].Line,
		)
	}
}
