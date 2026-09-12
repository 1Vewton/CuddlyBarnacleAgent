package task

import (
	"testing"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/textresult"
	"github.com/1Vewton/CuddlyBarnacleAgent/pkg/osoperation"
)

// TestManagerFileOperation tests the file operation of task manager
func TestManagerFileOperation(t *testing.T) {
	t.Parallel()
	rawTmpDir := t.TempDir()
	tmpDir := osoperation.PathToWindows(rawTmpDir)
	newManager := NewRawManager()
	err := newManager.Load(tmpDir, "test")
	if err != nil {
		t.Error(err)
	}
	newManager.Data["1"] = "abc"
	err = newManager.SaveFile()
	if err != nil {
		t.Error(err)
	}
	reloadedManager := NewRawManager()
	err = reloadedManager.Load(tmpDir, "test")
	if err != nil {
		t.Error(err)
	}
	if reloadedManager.Data["1"] != newManager.Data["1"] {
		t.Errorf(
			"expected %s, got %s",
			newManager.Data["1"],
			reloadedManager.Data["1"],
		)
	}
}

// TestAddQuestions tests the question adding
func TestAddQuestions(
	t *testing.T,
) {
	t.Parallel()
	// Test data
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
	rawTmpDir := t.TempDir()
	tmpDir := osoperation.PathToWindows(rawTmpDir)
	newManager := NewRawManager()
	// manager loading and operating
	err := newManager.Load(
		tmpDir,
		"test.json",
	)
	if err != nil {
		t.Error(err)
	}
	err = newManager.AddTask(
		"1",
		"114514",
	)
	if err != nil {
		t.Error(err)
	}
	err = newManager.AddQuestion(
		"1",
		testData,
	)
	if err != nil {
		t.Error(err)
	}
	task, err := newManager.GetTask(
		"1",
	)
	if err != nil {
		t.Error(err)
	}
	allProblems := task.GetAllProblemsFor(
		agenttypes.GrammaticalCheck,
	)
	if len(allProblems) != 1 {
		t.Errorf(
			"expected length of allProblems to be %d, got %d",
			1,
			len(allProblems),
		)
	}
	allProblems, err = newManager.GetAllProblemsFor(
		"1",
		agenttypes.GrammaticalCheck,
	)
	if err != nil {
		t.Error(err)
	}
	if len(allProblems) != 1 {
		t.Errorf(
			"expected length of allProblems to be %d, got %d",
			1,
			len(allProblems),
		)
	}
}
