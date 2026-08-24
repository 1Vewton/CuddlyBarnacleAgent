package datastruct

import (
	"reflect"
	"testing"
)

// TestStruct defines the data structure for test
type TestStruct struct {
	Field1 string
}

// TestGetField tests the get field function
func TestGetField(
	t *testing.T,
) {
	newTestStruct := TestStruct{
		Field1: "abc",
	}
	typeField1, typeValid := GetField(
		newTestStruct,
		"Field1",
	)
	if !typeValid {
		t.Error("this data struct is expected to be valid")
	}
	if typeField1 != reflect.String {
		t.Errorf(
			"expected type %s, got %s",
			"string",
			typeField1.String(),
		)
	}
	_, typeValid = GetField(
		newTestStruct,
		"Field2",
	)
	if typeValid {
		t.Error("Expected field2 to be invalid")
	}
}
