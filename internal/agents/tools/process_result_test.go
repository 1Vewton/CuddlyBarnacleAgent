package tools

import (
	"testing"
)

// TestStruct defines the data struct for testing
type TestStruct struct {
	Field1 int
}

// TestGiveSuccessfulResult tests giving successful result
func TestGiveSuccessfulResult(t *testing.T) {
	t.Parallel()
	definition := "test"
	testData := &TestStruct{
		Field1: 114514,
	}
	res := GiveSuccessResult(
		&definition,
		testData,
	)
	isResultError := IsResultError(res)
	if isResultError {
		t.Error(
			res,
		)
		t.Fatal(
			"expected not to get an error",
		)
	}
}
