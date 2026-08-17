package tools

import (
	"reflect"

	"github.com/cloudwego/eino/schema"
)

// ReflectTypeToEinoType converts reflect type to schema type.
// Allows:
// - Int
// - String
func ReflectTypeToEinoType(
	reflectType reflect.Kind,
) schema.DataType {
	switch reflectType {
	case reflect.String:
		return schema.String
	case reflect.Int:
		return schema.Integer
	case reflect.Bool:
		return schema.Boolean
	default:
		return schema.Null
	}
}

// NewParam creates new parameter
func NewParam(
	fieldType schema.DataType,
	description string,
	required bool,
) *schema.ParameterInfo {
	return &schema.ParameterInfo{
		Type:     fieldType,
		Desc:     description,
		Required: required,
	}
}
