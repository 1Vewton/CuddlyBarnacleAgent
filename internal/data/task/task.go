package task

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sync"
	"time"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/textresult"
	"github.com/google/uuid"
)

// Task defines task for checking errors
type Task struct {
	sync.RWMutex
	TaskName   string
	TaskID     string
	ArticleID  string
	CreateTime time.Time
	EditTime   time.Time
	Problems   []*textresult.StoredTextError
}

// NewTask creates new task
func NewTask(
	taskName string,
	articleID string,
) *Task {
	taskID := uuid.NewString()
	return &Task{
		TaskName:   taskName,
		TaskID:     taskID,
		ArticleID:  articleID,
		CreateTime: time.Now(),
		EditTime:   time.Now(),
		Problems:   []*textresult.StoredTextError{},
	}
}

// NewTaskFromFile creates new task from json file
func NewTaskFromFile(
	fileName string,
) (*Task, error) {
	newTask := &Task{}
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(
		data,
		&newTask,
	)
	return newTask, err
}

// GetFileName generates file name
func (task *Task) GetFileName(
	taskDir string,
) string {
	return fmt.Sprintf(
		"%s/%s_%d_%s.json",
		taskDir,
		task.TaskID,
		task.CreateTime.Nanosecond(),
		task.ArticleID,
	)
}

// GetTaskID gets the task id
func (task *Task) GetTaskID() string {
	task.RLock()
	defer task.RUnlock()
	result := task.TaskID
	return result
}

// SaveFile saves the task to json file
func (task *Task) SaveFile(
	taskDir string,
) (string, error) {
	task.RLock()
	defer task.RUnlock()
	data, err := json.Marshal(task)
	if err != nil {
		return "", err
	}
	fileName := task.GetFileName(taskDir)
	err = os.WriteFile(
		fileName,
		data,
		0644,
	)
	return fileName, err
}

// AddSingleProblem adds single problem to the Task
func (task *Task) AddSingleProblem(
	problem *textresult.StoredTextError,
) {
	task.Lock()
	defer task.Unlock()
	task.EditTime = time.Now()
	task.Problems = append(task.Problems, problem)
}

// AddMultipleProblems adds multiple problems to the Task
func (task *Task) AddMultipleProblems(
	problems []*textresult.StoredTextError,
) {
	for _, problem := range problems {
		task.AddSingleProblem(problem)
	}
}

// GetAllProblemsFor gets the all problems proposed by certain agent.
// Also returns the indexes of the returned problems.
func (task *Task) GetAllProblemsFor(
	agentType agenttypes.AgentType,
) ([]*textresult.StoredTextError, []int) {
	task.RLock()
	defer task.RUnlock()
	result := []*textresult.StoredTextError{}
	resultIdx := []int{}
	for i, task := range task.Problems {
		if agentType == task.Proposer {
			result = append(result, task)
			resultIdx = append(resultIdx, i)
		}
	}
	return result, resultIdx
}

// DeleteProblem deletes problem
func (task *Task) DeleteProblem(
	idx int,
	agentType agenttypes.AgentType,
) error {
	task.Lock()
	defer task.Unlock()
	if idx >= len(task.Problems) {
		return fmt.Errorf(
			"%d idx not exists",
			idx,
		)
	}
	if task.Problems[idx].Proposer != agentType {
		return fmt.Errorf(
			"the proposer of problem %d is not %s",
			idx,
			agentType.ToString(),
		)
	}
	result := slices.Delete(
		task.Problems,
		idx,
		idx+1,
	)
	task.Problems = result
	return nil
}

// OverWrite overwrites the problems proposed by certain agent
func (task *Task) OverWrite(
	agentType agenttypes.AgentType,
) {

}
