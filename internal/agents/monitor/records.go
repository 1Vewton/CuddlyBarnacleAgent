package monitor

import (
	"errors"
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

// CallSuccess records that the call for the tool is successed
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
	AgentsData map[string]agenttypes.AgentType
}
