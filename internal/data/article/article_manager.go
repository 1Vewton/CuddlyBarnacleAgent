package article

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// Manager manages articles
type Manager struct {
	sync.RWMutex
	Articles map[string]SingleArticle `json:"article"`
	filePath string
}

// NewManager creates new manager
func NewManager() *Manager {
	return &Manager{
		Articles: make(map[string]SingleArticle),
	}
}

// Load loads file as manager.
// No need to add suffix for fileName.
func (manager *Manager) Load(
	fileDir string,
	fileName string,
) error {
	manager.Lock()
	defer manager.Unlock()
	filePath := fmt.Sprintf(
		"%s/%s",
		fileDir,
		fileName,
	)
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			data, err := json.Marshal(
				manager,
			)
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
				filePath,
				data,
				0644,
			)
		}
		return err
	}
	data, err := os.ReadFile(
		filePath,
	)
	if err != nil {
		return err
	}
	err = json.Unmarshal(
		data,
		manager,
	)
	if err != nil {
		return err
	}
	manager.filePath = filePath
	return nil
}
