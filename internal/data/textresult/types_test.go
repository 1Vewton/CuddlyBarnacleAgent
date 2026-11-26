package textresult

import (
	"testing"
)

// TestDefine tests the defining of new error type
func TestDefine(t *testing.T) {
	t.Parallel()
	errorTypes := &AllErrorTypes{
		ErrorTypes: map[ErrorField][]*ErrorType{},
	}
	errorType0 := errorTypes.Define(
		"test",
		"error for testing",
		NotCategorized,
	)
	errorType1 := errorTypes.Define(
		"testLogic",
		"error for testing logic",
		Logic,
	)
	typeFetched0, err := errorTypes.GetTypeByErrorCode(
		0,
	)
	if err != nil {
		t.Fatal(err)
	}
	if !typeFetched0.Equals(errorType0) {
		t.Fatal(
			"the fetched error type 0 is not equals to expected type 0",
		)
	}
	typeFetched1, err := errorTypes.GetTypeByErrorCode(
		1,
	)
	if err != nil {
		t.Fatal(err)
	}
	if !typeFetched1.Equals(errorType1) {
		t.Fatal(
			"the fetched error type 1 is not equals to expected type 1",
		)
	}
}
