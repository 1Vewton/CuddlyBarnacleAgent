package textresult

import (
	"fmt"
	"sync"
)

// ErrorField defines field for error
type ErrorField int

const (
	// NotCategorized defines the field that the no error belongs to
	NotCategorized ErrorField = iota
	// Logic defines the field for logic errors
	Logic
	// Grammar defines the field for grammatical errors
	Grammar
)

// ErrorType defines type of error
type ErrorType struct {
	ErrorCode   int
	Name        string
	Explanation string
}

// Equals checks if the types are the same
func (errorType *ErrorType) Equals(
	another *ErrorType,
) bool {
	return errorType.ErrorCode == another.ErrorCode &&
		errorType.Explanation == another.Explanation &&
		errorType.Name == another.Name
}

// AllErrorTypes defines the all the error types
type AllErrorTypes struct {
	lock       sync.RWMutex
	ErrorTypes map[ErrorField][]*ErrorType
}

// Define defines new error type
func (all *AllErrorTypes) Define(
	name string,
	explanation string,
	field ErrorField,
) *ErrorType {
	all.lock.Lock()
	defer all.lock.Unlock()
	errorCode := len(all.ErrorTypes)
	newErrorType := &ErrorType{
		ErrorCode:   errorCode,
		Name:        name,
		Explanation: explanation,
	}
	_, exists := all.ErrorTypes[field]
	if !exists {
		all.ErrorTypes[field] = []*ErrorType{}
	}
	all.ErrorTypes[field] = append(all.ErrorTypes[field], newErrorType)
	return newErrorType
}

// AllPulbicErrors defines all the errors for public use
var AllPublicErrors = &AllErrorTypes{
	ErrorTypes: map[ErrorField][]*ErrorType{},
}

var (
	UncategorizedError = AllPublicErrors.Define(
		"uncategorized_error",
		"error that is not categorized",
		NotCategorized,
	)
)

// findAllErrorTypesForField finds all error types for a field
func (all *AllErrorTypes) findAllErrorTypesForField(
	field ErrorField,
) ([]*ErrorType, bool) {
	data, exists := all.ErrorTypes[field]
	if !exists {
		return nil, exists
	}
	return data, true
}

// GetAllErrorTypesForField gets all error types for a field
func (all *AllErrorTypes) GetAllErrorTypesForField(
	fields []ErrorField,
) ([]*ErrorType, error) {
	all.lock.RLock()
	defer all.lock.RUnlock()
	result := []*ErrorType{}
	for _, field := range fields {
		foundTypes, found := all.findAllErrorTypesForField(
			field,
		)
		if found {
			for _, foundType := range foundTypes {
				result = append(result, foundType)
			}
		}
	}
	return result, nil
}

// findType finds one single error type
func (all *AllErrorTypes) findType(
	field ErrorField,
	errorCode int,
) (*ErrorType, bool) {
	allTypesForAField, exists := all.ErrorTypes[field]
	if !exists {
		return nil, exists
	}
	for _, errorType := range allTypesForAField {
		if errorType.ErrorCode == errorCode {
			return errorType, true
		}
	}
	return nil, false
}

// GetType gets error type
func (all *AllErrorTypes) GetType(
	fields []ErrorField,
	errorCode int,
) (*ErrorType, error) {
	all.lock.RLock()
	defer all.lock.RUnlock()
	for _, field := range fields {
		typeFound, found := all.findType(
			field,
			errorCode,
		)
		if found {
			return typeFound, nil
		}
	}
	return nil, fmt.Errorf(
		"%d type with field cannot be found",
		errorCode,
	)
}

// GetTypeByErrorCode gets type directly by error code
func (all *AllErrorTypes) GetTypeByErrorCode(
	errorCode int,
) (*ErrorType, error) {
	all.lock.RLock()
	defer all.lock.RUnlock()
	for field := range all.ErrorTypes {
		typeFound, exists := all.findType(
			field,
			errorCode,
		)
		if exists {
			return typeFound, nil
		}
	}
	return nil, fmt.Errorf(
		"cannot find error type with code %d",
		errorCode,
	)
}

// GetFieldByErrorCode gets the field by error code
func (all *AllErrorTypes) GetFieldByErrorCode(
	errorCode int,
) (ErrorField, error) {
	all.lock.RLock()
	defer all.lock.RUnlock()
	for errorField := range all.ErrorTypes {
		_, found := all.findType(
			errorField,
			errorCode,
		)
		if found {
			return errorField, nil
		}
	}
	return NotCategorized, fmt.Errorf(
		"error type with %d code is not found",
		errorCode,
	)
}
