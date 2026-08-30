package task

import (
	"sync"
	"time"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/textresult"
)

// Task defines task for checking errors
type Task struct {
	sync.RWMutex
	ArticleID   string
	CreatedTime time.Time
	EditTime    time.Time
	Problems    []*textresult.StoredTextError
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
