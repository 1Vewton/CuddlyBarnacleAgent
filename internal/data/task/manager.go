package task

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// Manager defines the manager for tasks
type Manager struct {
	sync.RWMutex
	Data     map[string]string
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
		err = os.WriteFile(
			path,
			data,
			0644,
		)
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
