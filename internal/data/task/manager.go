package task

import (
	"sync"
)

// Manager defines the manager for tasks
type Manager struct {
	sync.RWMutex
	Data    map[string]string
	taskDir string
}
