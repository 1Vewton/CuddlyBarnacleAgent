package monitor

import (
	"errors"
	"sync"
)

// NormalRecord records normal call
type NormalRecord struct {
	ToolName string
}

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
	TotalCalls []*NormalRecord
	Failures   []*FailureRecord
	Successes  []*SuccessRecord
}

// NewSingleAgentRecord creates new Record
func NewSingleAgentRecord() *SingleAgentRecord {
	return &SingleAgentRecord{
		TotalCalls: []*NormalRecord{},
		Failures:   []*FailureRecord{},
		Successes:  []*SuccessRecord{},
	}
}

// NewCall adds total calls
func (record *SingleAgentRecord) NewCall(
	toolName string,
) {
	record.Lock()
	defer record.Unlock()
	newCall := &NormalRecord{
		ToolName: toolName,
	}
	record.TotalCalls = append(record.TotalCalls, newCall)
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
	result := len(record.TotalCalls)
	return result
}

// GetSuccessCalls gets success calls
func (record *SingleAgentRecord) GetSuccessCalls() int {
	record.RLock()
	defer record.RUnlock()
	result := len(record.Successes)
	return result
}

// GetFailedCalls gets failed calls
func (record *SingleAgentRecord) GetFailedCalls() int {
	record.RLock()
	defer record.RUnlock()
	result := len(record.Failures)
	return result
}
