package prompts

import (
	"fmt"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/textresult"
)

const (
	lineDescription  string = "The position of the error"
	levelDescription string = `
The level of this problem. 
- 0: Warning
- 1: Error
	`
	reasonDescription string = "The reason you believe this part is an error"
)

// GetTypeDescription gets the description for error types
func GetTypeDescription(
	errorTypes *textresult.AllErrorTypes,
	fields []textresult.ErrorField,
) (string, error) {
	description := `
The error types available for you to add: 
Type in the error code for error if you want add new error to text result: 
	`
	allErrorTypes, err := errorTypes.GetAllErrorTypesForField(
		fields,
	)
	if err != nil {
		return description, err
	}
	for _, errorType := range allErrorTypes {
		description += fmt.Sprintf(
			"\nerror type %s\nexplanation: %s\nErrorCode: %d\n",
			errorType.Name,
			errorType.Explanation,
			errorType.ErrorCode,
		)
	}
	return description, nil
}

// GetTextResultParam gets text result param
func GetTextResultParam(
	typeDescription string,
) map[string]ParameterDescription {
	return map[string]ParameterDescription{
		"Line": {
			Description: lineDescription,
			Required:    true,
		},
		"Level": {
			Description: levelDescription,
			Required:    true,
		},
		"Type": {
			Description: typeDescription,
			Required:    true,
		},
		"Reason": {
			Description: reasonDescription,
			Required:    true,
		},
	}
}
