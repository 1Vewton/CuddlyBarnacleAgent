package tools

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ShowError shows the crucial error to agent directly
func ShowError(
	errInfo error,
) string {
	format := `{"crucial_error": "%s"}`
	var errData string
	if errInfo == nil {
		errData = "unknown error occurs"
	} else {
		errData = errInfo.Error()
	}
	return fmt.Sprintf(
		format,
		errData,
	)
}

// IsResultError shows whether result is a crucial error
func IsResultError(
	data string,
) bool {
	if len(data) <= 19 {
		return false
	}
	firstPart := data[:19]
	if firstPart != `{"crucial_error": ` {
		return false
	}
	return true
}

// DataWithDefinition defines data with definition
type DataWithDefinition struct {
	Value      any     `json:"value"`
	Definition *string `json:"definition"`
}

// BaseProcessResult defines the basic process result
type BaseProcessResult struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    *DataWithDefinition `json:"data"`
}

// GiveSuccessResult gives the success result directly
func GiveSuccessResult(
	definitionForData *string,
	data any,
) string {
	var resultData *DataWithDefinition
	resultData = &DataWithDefinition{
		Value:      data,
		Definition: definitionForData,
	}
	result := &BaseProcessResult{
		Success: true,
		Message: "tool execution succeed",
		Data:    resultData,
	}
	bytesData, err := json.Marshal(result)
	if err != nil {
		return ShowError(
			err,
		)
	}
	return string(bytesData)
}

// GiveFailResult gives a failed result
func GiveFailResult(
	errInfo error,
) string {
	result := &BaseProcessResult{
		Success: false,
		Data:    nil,
	}
	if errInfo != nil {
		result.Message = errInfo.Error()
	} else {
		result.Message = "unknown error occurs"
	}
	data, err := json.Marshal(result)
	if err != nil {
		return ShowError(
			err,
		)
	}
	return string(data)
}

// GetResultInStruct gets the processed result in data struct
func GetResultInStruct(
	rawResult string,
) (*BaseProcessResult, error) {
	if IsResultError(rawResult) {
		return nil, errors.New(
			rawResult,
		)
	}
	var result *BaseProcessResult
	err := json.Unmarshal(
		[]byte(rawResult),
		&result,
	)
	return result, err
}
