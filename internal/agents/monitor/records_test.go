package monitor

import (
	"testing"
)

// TestSingleAgentRecord tests operations related to single agent record
func TestSingleAgentRecord(t *testing.T) {
	record := NewSingleAgentRecord()
	record.NewCall()
	if record.GetTotalCalls() != 1 {
		t.Errorf(
			"expected total calls to be %d, got %d",
			1,
			record.GetTotalCalls(),
		)
	}
}
