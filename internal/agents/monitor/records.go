package monitor

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
)

// FailureRecord records failure
type FailureRecord struct {
	Reason     string
	AgentInput string
	ToolName   string
}

// NewFailureRecord creates new failure record
func NewFailureRecord(
	reason string,
	agentInput string,
	toolName string,
) *FailureRecord {
	return &FailureRecord{
		Reason:     reason,
		AgentInput: agentInput,
		ToolName:   toolName,
	}
}

// NewFailureRecordFromError creates new failure record from error
func NewFailureRecordFromError(
	errInfo error,
	agentInput string,
	toolName string,
) (*FailureRecord, error) {
	if errInfo == nil {
		return nil, errors.New(
			"The error passed cannot be nil",
		)
	}
	record := NewFailureRecord(
		errInfo.Error(),
		agentInput,
		toolName,
	)
	return record, nil
}

// SuccessRecord defines the record for success
type SuccessRecord struct {
	ToolName string
}

// SingleAgentRecord defines the recorded data for single agent
type SingleAgentRecord struct {
	sync.RWMutex
	TotalCalls int
	Failures   []*FailureRecord
	Successes  []*SuccessRecord
}

// NewSingleAgentRecord creates new Record
func NewSingleAgentRecord() *SingleAgentRecord {
	return &SingleAgentRecord{
		TotalCalls: 0,
		Failures:   []*FailureRecord{},
		Successes:  []*SuccessRecord{},
	}
}

// NewCall adds total calls
func (record *SingleAgentRecord) NewCall() {
	record.Lock()
	defer record.Unlock()
	record.TotalCalls++
}

// CallSuccess records the circumstance that the call for the tool is successed
func (record *SingleAgentRecord) CallSuccess(
	toolName string,
) {
	record.Lock()
	defer record.Unlock()
	newSuccess := &SuccessRecord{
		ToolName: toolName,
	}
	record.Successes = append(record.Successes, newSuccess)
}

// DetailedCallFail records the circumstance that the call for the tool failed in detailed information
func (record *SingleAgentRecord) DetailedCallFail(
	errInfo error,
	agentInput string,
	toolName string,
) {
	record.Lock()
	defer record.Unlock()
	newFailure := NewFailureRecord(
		errInfo.Error(),
		agentInput,
		toolName,
	)
	record.Failures = append(record.Failures, newFailure)
}

// CallFail records the circumastance that the call for the tool failed
func (record *SingleAgentRecord) CallFail(
	toolName string,
) {
	record.Lock()
	defer record.Unlock()
	newFailure := &FailureRecord{
		ToolName: toolName,
	}
	record.Failures = append(record.Failures, newFailure)
}

// GetTotalCalls gets total calls
func (record *SingleAgentRecord) GetTotalCalls() int {
	record.RLock()
	defer record.RUnlock()
	result := record.TotalCalls
	return result
}

// GetSuccessCalls gets success calls
func (record *SingleAgentRecord) GetSuccessCalls() int {
	record.RLock()
	defer record.RUnlock()
	result := len(record.Successes)
	return result
}

// AgentsRecord records performance of all agents
type AgentsRecord struct {
	sync.RWMutex
	AgentsData map[agenttypes.AgentType]*SingleAgentRecord
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

// NewCall calls when a new tool call is executed by agent
func (record *AgentsRecord) NewCall(
	agentType agenttypes.AgentType,
) {
	record.Lock()
	defer record.Unlock()
	data, exists := record.AgentsData[agentType]
	if !exists || data == nil {
		record.AgentsData[agentType] = NewSingleAgentRecord()
	}
	record.AgentsData[agentType].NewCall()
}

// AddSuccessCall adds new success call record
func (record *AgentsRecord) AddSuccessCall(
	agentType agenttypes.AgentType,
	toolName string,
) {
	record.Lock()
	defer record.Unlock()
	data, exists := record.AgentsData[agentType]
	if !exists || data == nil {
		record.AgentsData[agentType] = NewSingleAgentRecord()
		record.AgentsData[agentType].NewCall()
	}
	record.AgentsData[agentType].CallSuccess(
		toolName,
	)
}

// AddFailureRecord adds new failure record
func (record *AgentsRecord) AddFailureRecord(
	agentType agenttypes.AgentType,
	toolName string,
) {
	record.Lock()
	defer record.Unlock()
	data, exists := record.AgentsData[agentType]
	if !exists || data == nil {
		record.AgentsData[agentType] = NewSingleAgentRecord()
		record.AgentsData[agentType].NewCall()
	}
	record.AgentsData[agentType].CallFail(
		toolName,
	)
}

// GetTotalCalls gets total calls for an agent
func (record *AgentsRecord) GetTotalCalls(
	agentType agenttypes.AgentType,
) int {
	record.RLock()
	defer record.RUnlock()
	result := record.AgentsData[agentType].GetTotalCalls()
	return result
}
