package agents

import (
	"testing"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/monitor"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/tools"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/task"
)

// TestToolsConfig tests the configuration for tools
func TestToolsConfig(t *testing.T) {
	t.Parallel()
	const taskName string = "taskA"
	tmpDir := t.TempDir()
	testMonitor := monitor.NewAgentsRecord()
	testMonitor.Load(
		tmpDir,
		"monitor.json",
	)
	testManager := task.NewRawManager()
	testManager.Load(
		tmpDir,
		"test.json",
	)
	err := testManager.AddTask(
		taskName,
		"test",
	)
	if err != nil {
		t.Fatal(err)
	}
	testToolList := NewToolsList()
	testToolList.AddTool(
		tools.NewProposeProblemTool(
			taskName,
			agenttypes.FactualCheck,
			testMonitor,
			testManager,
		),
	)
	if len(testToolList.Tools) != 1 {
		t.Errorf(
			"expected length of the tool list to be %d, got %d",
			1,
			len(testToolList.Tools),
		)
	}
}
