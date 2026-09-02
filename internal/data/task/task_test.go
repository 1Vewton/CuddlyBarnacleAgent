package task

import (
	"testing"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/textresult"
)

// TestTask tests the functions related to
func TestTask(t *testing.T) {
	t.Parallel()
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
	newTask := NewTask("TestTask", "114514")
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

// TestFileOperation tests the load and saving of task
func TestFileOperation(t *testing.T) {
	t.Parallel()
	tmpDir := t.TempDir()
	newTask := NewTask(
		"test",
		"1919810",
	)
	fileName, err := newTask.SaveFile(tmpDir)
	if err != nil {
		t.Error(err)
	}
	reloadedTask, err := NewTaskFromFile(
		fileName,
	)
	if err != nil {
		t.Error(err)
	}
	if reloadedTask.GetTaskID() != newTask.GetTaskID() {
		t.Errorf(
			"expected %s, got %s",
			newTask.GetTaskID(),
			reloadedTask.GetTaskID(),
		)
	}
}
