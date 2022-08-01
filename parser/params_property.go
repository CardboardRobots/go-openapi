package parser

import (
	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

func GetParams(operation *openapi3.Operation) []entity.ParamProperty {
	parameters := make([]entity.ParamProperty, 0)
	for _, parameterRef := range operation.Parameters {
		parameter := parameterRef.Value
		if parameter.In == openapi3.ParameterInPath {
			parameters = append(parameters, entity.ParamProperty{
				Name: GetPropertyName(parameter.Name),
				// Type: GetPropertyType(parameter.Schema.Value.Type),
				Key: parameter.Name,
			})
		}
	}
	return parameters
}
