package datastruct

import (
	"reflect"
)

// GetField gets the field of certain data struct
func GetField(
	data any,
	fieldName string,
) (reflect.Kind, bool) {
	val := reflect.ValueOf(data)
	fieldVal := val.FieldByName(fieldName)
	if !fieldVal.IsValid() {
		return reflect.Invalid, false
	}
	return fieldVal.Kind(), true
}
