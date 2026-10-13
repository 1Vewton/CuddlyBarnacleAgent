package monitor

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
)

// AgentsRecord records performance of all agents
type AgentsRecord struct {
	sync.RWMutex
	AgentsData map[agenttypes.AgentType]*SingleAgentRecord
	filePath   string
}

// NewAgentsRecord creates new agent record
func NewAgentsRecord() *AgentsRecord {
	return &AgentsRecord{
		AgentsData: map[agenttypes.AgentType]*SingleAgentRecord{},
	}
}

// Load loads record from file
func (record *AgentsRecord) Load(
	fileDir string,
	fileName string,
) error {
	record.Lock()
	defer record.Unlock()
	filePath := fmt.Sprintf(
		"%s/%s",
		fileDir,
		fileName,
	)
	record.filePath = filePath
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Stat(fileDir)
			if err != nil {
				if os.IsNotExist(err) {
					err = os.Mkdir(fileDir, 0644)
					if err != nil {
						return err
					}
				} else {
					return err
				}
			}
			data, err := json.Marshal(record)
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
	}
	data, err := os.ReadFile(
		filePath,
	)
	if err != nil {
		return err
	}
	err = json.Unmarshal(
		data,
		&record,
	)
	return err
}

// fillEmpty fills empty in map in found
func (record *AgentsRecord) fillEmpty(
	agentType agenttypes.AgentType,
	toolName string,
	needsAddRecord bool,
) {
	data, exists := record.AgentsData[agentType]
	if !exists || data == nil {
		record.AgentsData[agentType] = NewSingleAgentRecord()
		if needsAddRecord {
			record.AgentsData[agentType].NewCall(
				toolName,
			)
		}
	}
}

// NewCall calls when a new tool call is executed by agent
func (record *AgentsRecord) NewCall(
	agentType agenttypes.AgentType,
	toolName string,
) {
	record.Lock()
	defer record.Unlock()
	record.fillEmpty(
		agentType,
		toolName,
		false,
	)
	record.AgentsData[agentType].NewCall(
		toolName,
	)
}

// AddSuccessCall adds new success call record
func (record *AgentsRecord) AddSuccessCall(
	agentType agenttypes.AgentType,
	toolName string,
) {
	record.Lock()
	defer record.Unlock()
	record.fillEmpty(
		agentType,
		toolName,
		true,
	)
	record.AgentsData[agentType].CallSuccess(
		toolName,
	)
}

// AddFailureCall adds new failure call
func (record *AgentsRecord) AddFailureCall(
	agentType agenttypes.AgentType,
	toolName string,
) {
	record.Lock()
	defer record.Unlock()
	record.fillEmpty(
		agentType,
		toolName,
		true,
	)
	record.AgentsData[agentType].CallFail(
		toolName,
	)
}

// AddFailureCallByDetail adds new failure call in detail
func (record *AgentsRecord) AddFailureCallByDetail(
	agentType agenttypes.AgentType,
	toolName string,
	errInfo error,
	agentInput string,
) {
	record.Lock()
	defer record.Unlock()
	record.fillEmpty(
		agentType,
		toolName,
		true,
	)
	record.AgentsData[agentType].DetailedCallFail(
		errInfo,
		agentInput,
		toolName,
	)
}

// GetTotalCalls gets total calls for an agent
func (record *AgentsRecord) GetTotalCalls(
	agentType agenttypes.AgentType,
) int {
	record.RLock()
	defer record.RUnlock()
	record.fillEmpty(
		agentType,
		"N/A",
		false,
	)
	result := record.AgentsData[agentType].GetTotalCalls()
	return result
}

// GetSuccessfulCalls gets number of successful calls for an agent
func (record *AgentsRecord) GetSuccessfulCalls(
	agentType agenttypes.AgentType,
) int {
	record.RLock()
	defer record.RUnlock()
	record.fillEmpty(
		agentType,
		"N/A",
		false,
	)
	result := record.AgentsData[agentType].GetSuccessCalls()
	return result
}

// GetFailedCalls gets number of failed calls for an agent
func (record *AgentsRecord) GetFailedCalls(
	agentType agenttypes.AgentType,
) int {
	record.RLock()
	defer record.RUnlock()
	record.fillEmpty(
		agentType,
		"N/A",
		false,
	)
	result := record.AgentsData[agentType].GetFailedCalls()
	return result
}

// SaveFile saves the file
func (record *AgentsRecord) SaveFile() error {
	record.Lock()
	defer record.Unlock()
	data, err := json.Marshal(record)
	if err != nil {
		return err
	}
	err = os.WriteFile(
		record.filePath,
		data,
		0644,
	)
	return err
}

// NewCallWithSave adds new call and saves the file
func (record *AgentsRecord) NewCallWithSave(
	agentType agenttypes.AgentType,
	toolName string,
) error {
	record.NewCall(
		agentType,
		toolName,
	)
	err := record.SaveFile()
	return err
}

// AddSuccessCallWithSave adds successful call and saves the file
func (record *AgentsRecord) AddSuccessCallWithSave(
	agentType agenttypes.AgentType,
	toolName string,
) error {
	record.AddSuccessCall(
		agentType,
		toolName,
	)
	err := record.SaveFile()
	return err
}

// AddFailureCallWithSave adds calls that will fail and saves the file
func (record *AgentsRecord) AddFailureCallWithSave(
	agentType agenttypes.AgentType,
	toolName string,
) error {
	record.AddFailureCall(
		agentType,
		toolName,
	)
	err := record.SaveFile()
	return err
}
