package article

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"sync"

	"github.com/google/uuid"
)

// Manager manages articles
type Manager struct {
	sync.RWMutex
	Articles map[string]*SingleArticle `json:"article"`
	filePath string
}

// NewManager creates new manager
func NewManager() *Manager {
	return &Manager{
		Articles: make(map[string]*SingleArticle),
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
	manager.filePath = filePath
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
			return err
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
	return nil
}

// GetAllArticles gets all articles
func (manager *Manager) GetAllArticles() map[string]*SingleArticle {
	manager.RLock()
	defer manager.RUnlock()
	return maps.Clone(
		manager.Articles,
	)
}

// AddNewArticle adds new article.
// Returns error and id of article
func (manager *Manager) AddNewArticle(
	originalFilePath string,
	targetDirPath string,
) (string, error) {
	manager.Lock()
	defer manager.Unlock()
	id := uuid.NewString()
	article, err := CreateNewSingleArticle(
		originalFilePath,
		targetDirPath,
	)
	if err != nil {
		return id, err
	}
	manager.Articles[id] = article
	return id, nil
}

// save saves the data to file
func (manager *Manager) save() error {
	data, err := json.Marshal(
		manager,
	)
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

// Save exportes save.
// save saves the data to file.
func (manager *Manager) Save() error {
	manager.Lock()
	defer manager.Unlock()
	err := manager.save()
	return err
}
