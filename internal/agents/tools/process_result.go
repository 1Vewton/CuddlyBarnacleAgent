package tools

import (
	"encoding/json"
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

// DataWithDefinition defines data with definition
type DataWithDefinition struct {
	Value      any     `json:"value"`
	Definition *string `json:"definition"`
}

// BaseProcessResult defines the basic process result
type BaseProcessResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
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
