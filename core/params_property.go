package core

import "github.com/getkin/kin-openapi/openapi3"

func GetParams(operation *openapi3.Operation) []ParamProperty {
	parameters := make([]ParamProperty, 0)
	for _, parameterRef := range operation.Parameters {
		parameter := parameterRef.Value
		if parameter.In == openapi3.ParameterInPath {
			parameters = append(parameters, ParamProperty{
				Name: GetPropertyName(parameter.Name),
				Type: GetPropertyType(parameter.Schema.Value.Type),
				Key:  parameter.Name,
			})
		}
	}
	return parameters
}

type ParamProperty struct {
	Name string
	Type string
	Key  string
}
