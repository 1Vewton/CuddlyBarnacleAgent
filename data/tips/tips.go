package tips

import (
	"encoding/json"
	"os"
	"slices"
	"sync"
)

// Tips stores the tips for the agent
type Tips struct {
	sync.RWMutex
	path string
	Data []string
}

// NewTips creates new tips
func NewTips(
	path string,
) *Tips {
	return &Tips{
		path: path,
		Data: []string{},
	}
}

// FromFile set the Tips from file
func (tips *Tips) FromFile() error {
	tips.Lock()
	defer tips.Unlock()
	// File checking and creating
	_, errExist := os.Stat(tips.path)
	if errExist != nil {
		if os.IsExist(errExist) {
			file, errCreate := os.Create(tips.path)
			if errCreate != nil {
				return errCreate
			}
			defer file.Close()
		} else {
			return errExist
		}
	}
	// Read file and decode
	data, errRead := os.ReadFile(tips.path)
	if errRead != nil {
		return errRead
	}
	errDecode := json.Unmarshal(data, tips)
	return errDecode
}

// NewTip creates new tip to tips
func (tips *Tips) NewTip(
	tip string,
) {
	tips.Lock()
	defer tips.Unlock()
	tips.Data = append(tips.Data, tip)
}

// GetAllTips gets all the tips
func (tips *Tips) GetAllTips() []string {
	tips.RLock()
	defer tips.RUnlock()
	return tips.Data
}

// ReturnTipsThroughIdx gets the tips through a idx list
func (tips *Tips) ReturnTipsThroughIdx(
	idx []int,
) []string {
	tips.RLock()
	defer tips.RUnlock()
	// Process input idx
	idxClone := idx
	slices.Sort(idxClone)
	// Get result
	result := []string{}
	currentIdx := 0
	for i, data := range tips.Data {
		if idxClone[currentIdx] == i {
			result = append(result, data)
			currentIdx += 1
			if currentIdx >= len(idxClone) {
				break
			}
		}
	}
	return result
}

// SaveFile saves the data in file
func (tips *Tips) SaveFile() error {
	tips.RLock()
	defer tips.RUnlock()
	tipsByte, errMarshal := json.Marshal(tips)
	if errMarshal != nil {
		return errMarshal
	}
	errWrite := os.WriteFile(
		tips.path,
		tipsByte,
		0644,
	)
	return errWrite
}
