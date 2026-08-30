package monitor

import (
	"sync"
)

// FailureRecord records failure
type FailureRecord struct {
	Reason     string
	AgentInput string
	ToolName   string
}

// SingleAgentRecord defines the recorded data for single agent
type SingleAgentRecord struct {
	sync.RWMutex
	TotalCalls   int
	TotalSuccess int
	Failures     []FailureRecord
}

// NewSingleAgentRecord creates new Record
func NewSingleAgentRecord() *SingleAgentRecord {
	return &SingleAgentRecord{
		TotalCalls:   0,
		TotalSuccess: 0,
		Failures:     []FailureRecord{},
	}
}

// NewCall adds total calls
func (record *SingleAgentRecord) NewCall() {
	record.Lock()
	defer record.Unlock()
	record.TotalCalls += 1
}

// CallSuccess records that the call for the tool is successed
func (record *SingleAgentRecord) CallSuccess() {
	record.Lock()
	defer record.Unlock()
	record.TotalSuccess += 1
}

// GetTotalCalls gets total calls
func (record *SingleAgentRecord) GetTotalCalls() int {
	record.Lock()
	defer record.Unlock()
	result := record.TotalCalls
	return result
}
