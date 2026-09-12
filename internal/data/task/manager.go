package task

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/textresult"
)

// Manager defines the manager for tasks
type Manager struct {
	sync.RWMutex
	Data     map[string]string
	taskDir  string
	filePath string
}

// NewManager creates new manager.
// For fileName field, not .json is needed.
func NewManager(
	fileDir string,
	fileName string,
) *Manager {
	return &Manager{
		Data: make(map[string]string),
		filePath: fmt.Sprintf(
			"%s/%s.json",
			fileDir,
			fileName,
		),
		taskDir: fileDir,
	}
}

// NewRawManager just creats a new manager
func NewRawManager() *Manager {
	return &Manager{
		Data: make(map[string]string),
	}
}

// Load loads json file to the manager
func (manager *Manager) Load(
	fileDir string,
	fileName string,
) error {
	manager.Lock()
	defer manager.Unlock()
	path := fmt.Sprintf(
		"%s/%s.json",
		fileDir,
		fileName,
	)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		manager.Data = make(map[string]string)
		data, err := json.Marshal(manager)
		if err != nil {
			return err
		}
		err = os.MkdirAll(
			fileDir,
			0644,
		)
		if err != nil {
			return err
		}
		err = os.WriteFile(
			path,
			data,
			0644,
		)
		if err != nil {
			return err
		}
	} else if err == nil {
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, manager)
		if err != nil {
			return err
		}
	} else {
		return err
	}
	manager.filePath = path
	manager.taskDir = fileDir
	return nil
}

// SaveFile saves file to the target file
func (manager *Manager) SaveFile() error {
	manager.RLock()
	defer manager.RUnlock()
	data, err := json.Marshal(manager)
	if err != nil {
		return err
	}
	err = os.WriteFile(
		manager.filePath,
		data,
		0644,
	)
	return err
}

// AddTask adds new task to the target file
func (manager *Manager) AddTask(
	taskName string,
	articleID string,
) error {
	manager.Lock()
	defer manager.Unlock()
	_, ok := manager.Data[taskName]
	if ok {
		return fmt.Errorf(
			"task %s already exists",
			taskName,
		)
	}
	newTask := NewTask(taskName, articleID)
	fileName, err := newTask.SaveFile(
		manager.taskDir,
	)
	if err != nil {
		return err
	}
	manager.Data[taskName] = fileName
	return nil
}

// GetTask gets task from the data
func (manager *Manager) GetTask(
	taskName string,
) (*Task, error) {
	manager.RLock()
	defer manager.RUnlock()
	fileName, exists := manager.Data[taskName]
	if !exists {
		return nil, fmt.Errorf(
			"%s not exists",
			taskName,
		)
	}
	return NewTaskFromFile(
		fileName,
	)
}

// AddQuestion adds question to it
func (manager *Manager) AddQuestion(
	taskName string,
	problems []*textresult.StoredTextError,
) error {
	manager.RLock()
	defer manager.RUnlock()
	fileName, exists := manager.Data[taskName]
	if !exists {
		return fmt.Errorf(
			"%s not exists",
			taskName,
		)
	}
	task, err := NewTaskFromFile(
		fileName,
	)
	if err != nil {
		return err
	}
	task.AddMultipleProblems(
		problems,
	)
	storedFileName, err := task.SaveFile(
		manager.taskDir,
	)
	if err != nil {
		return err
	}
	manager.Data[taskName] = storedFileName
	return nil
}

// GetAllProblemsForTask gets all problems for certain task
func (manager *Manager) GetAllProblemsFor(
	taskName string,
	proposer agenttypes.AgentType,
) ([]*textresult.StoredTextError, error) {
	manager.RLock()
	defer manager.RUnlock()
	fileName, exists := manager.Data[taskName]
	if !exists {
		return nil, fmt.Errorf(
			"%s not exists",
			taskName,
		)
	}
	task, err := NewTaskFromFile(
		fileName,
	)
	if err != nil {
		return nil, err
	}

	return task.GetAllProblemsFor(
		proposer,
	), nil
}
