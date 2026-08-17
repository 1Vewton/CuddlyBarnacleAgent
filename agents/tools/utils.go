package tools

import (
	"fmt"
	"reflect"

	"github.com/1Vewton/CuddlyBarnacleAgent/agents"
	"github.com/1Vewton/CuddlyBarnacleAgent/utils/datastruct"
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

// CreateToolParamList creates param list for tool
func CreateToolParamList(
	params map[string]agents.ParameterDescription,
	data any,
) (map[string]*schema.ParameterInfo, error) {
	result := make(map[string]*schema.ParameterInfo)
	for parameterName, description := range params {
		fieldType, isFieldExists := datastruct.GetField(
			data,
			parameterName,
		)
		if !isFieldExists {
			return result, fmt.Errorf(
				"%s field does not exists in the data",
				parameterName,
			)
		}
		einoType := ReflectTypeToEinoType(fieldType)
		if einoType == schema.Null {
			return result, fmt.Errorf(
				"Data type for field %s is not supported (type: %s)",
				parameterName,
				fieldType.String(),
			)
		}
		result[parameterName] = NewParam(
			einoType,
			description.Description,
			description.Required,
		)
	}
	return result, nil
}

// CreateParamList creates description for a list
func CreateParamList(
	subparams map[string]*schema.ParameterInfo,
	subparamsDescription string,
	required bool,
	description string,
) *schema.ParameterInfo {
	return &schema.ParameterInfo{
		Type:     schema.Array,
		Desc:     description,
		Required: required,
		ElemInfo: &schema.ParameterInfo{
			Type:      schema.Object,
			Desc:      subparamsDescription,
			SubParams: subparams,
		},
	}
}
