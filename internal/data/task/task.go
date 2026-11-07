package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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
	task.Problems = append(
		task.Problems[:idx],
		task.Problems[idx+1:]...,
	)
	return nil
}

// OverWrite overwrites one problem proposed by certain agent
func (task *Task) OverWrite(
	idx int,
	agentType agenttypes.AgentType,
	problem *textresult.StoredTextError,
) error {
	task.Lock()
	defer task.Unlock()
	if idx > len(task.Problems) {
		return fmt.Errorf(
			"%d is out of range (0-%d)",
			idx,
			len(task.Problems),
		)
	}
	for i := range task.Problems {
		if idx == i {
			if task.Problems[idx].Proposer != agentType {
				return fmt.Errorf(
					"the proposer fo the original question is %s, not you",
					task.Problems[idx].Proposer.ToString(),
				)
			}
		}
	}
	return nil
}

// DeleteAllProblemsFor deletes all problems for an agent type
func (task *Task) DeleteAllProblemsFor(
	agentType agenttypes.AgentType,
) error {
	task.Lock()
	defer task.Unlock()
	if task.Problems == nil {
		return errors.New(
			"the problems list is nil",
		)
	}
	for i, problem := range task.Problems {
		if problem == nil {
			return fmt.Errorf(
				"problem list %d is nil",
				i,
			)
		}
		if problem.Proposer == agentType {
			task.Problems = append(
				task.Problems[:i],
				task.Problems[i+1:]...,
			)
		}
	}
	return nil
}

// OverWriteAll defines the function for overwriting all problems
func (task *Task) OverWriteAll(
	problems []*textresult.StoredTextError,
	agentType agenttypes.AgentType,
) error {
	err := task.DeleteAllProblemsFor(agentType)
	if err != nil {
		return err
	}
	task.AddMultipleProblems(problems)
	return nil
}
