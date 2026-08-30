package textresult

import (
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents"
)

// TextError defines an error occurs in the text, including logic and other errors
type TextError struct {
	Level  int
	Type   int
	Line   int
	Reason string
}

// ListResult defines the list of errors
type ListResult struct {
	Array                 []TextError
	OverWriteOriginalData bool
}

// StoredTextError defines the error stored in the program
type StoredTextError struct {
	Level    Levels
	Type     ErrorType
	Line     int
	Reason   string
	Proposer agents.AgentType
}

// FromRawToProcessed process the basic text error
func FromRawToProcessed(
	rawData TextError,
	proposer agents.AgentType,
) (*StoredTextError, error) {
	level, err := ToLevels(rawData.Level)
	if err != nil {
		return nil, err
	}
	return &StoredTextError{
		Level:    level,
		Type:     ToErrorType(rawData.Type),
		Line:     rawData.Line,
		Reason:   rawData.Reason,
		Proposer: proposer,
	}, nil
}

// MultipleRawToProcessed process mutilple text errors
func MultipleRawToProcessed(
	rawDatas []TextError,
	proposer agents.AgentType,
) ([]*StoredTextError, error) {
	result := []*StoredTextError{}
	for _, data := range rawDatas {
		processedData, err := FromRawToProcessed(data, proposer)
		if err != nil {
			return nil, err
		}
		result = append(result, processedData)
	}
	return result, nil
}
