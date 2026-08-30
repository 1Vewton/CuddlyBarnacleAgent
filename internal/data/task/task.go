package task

import (
	"sync"
	"time"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/textresult"
	"github.com/google/uuid"
)

// Task defines task for checking errors
type Task struct {
	sync.RWMutex
	TaskID     string
	ArticleID  string
	CreateTime time.Time
	EditTime   time.Time
	Problems   []*textresult.StoredTextError
}

// NewTask creates new task
func NewTask(
	articleID string,
) *Task {
	taskID := uuid.NewString()
	return &Task{
		TaskID:     taskID,
		ArticleID:  articleID,
		CreateTime: time.Now(),
		EditTime:   time.Now(),
		Problems:   []*textresult.StoredTextError{},
	}
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

// GetAllProblemsFor gets the all tasks proposed by certain agent
func (task *Task) GetAllProblemsFor(
	agentType agenttypes.AgentType,
) []*textresult.StoredTextError {
	task.RLock()
	defer task.RUnlock()
	result := []*textresult.StoredTextError{}
	for _, task := range task.Problems {
		if agentType == task.Proposer {
			result = append(result, task)
		}
	}
	return result
}
