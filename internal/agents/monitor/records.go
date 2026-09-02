package monitor

import (
	"errors"
	"sync"
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

// SingleAgentRecord defines the recorded data for single agent
type SingleAgentRecord struct {
	sync.RWMutex
	TotalCalls   int
	TotalSuccess int
	Failures     []*FailureRecord
}

// NewSingleAgentRecord creates new Record
func NewSingleAgentRecord() *SingleAgentRecord {
	return &SingleAgentRecord{
		TotalCalls:   0,
		TotalSuccess: 0,
		Failures:     []*FailureRecord{},
	}
}

// NewCall adds total calls
func (record *SingleAgentRecord) NewCall() {
	record.Lock()
	defer record.Unlock()
	record.TotalCalls++
}

// CallSuccess records that the call for the tool is successed
func (record *SingleAgentRecord) CallSuccess() {
	record.Lock()
	defer record.Unlock()
	record.TotalSuccess++
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
	result := record.TotalSuccess
	return result
}
