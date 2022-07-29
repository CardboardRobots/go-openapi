package core

import (
	"github.com/CardboardRobots/go-array"
	"github.com/getkin/kin-openapi/openapi3"
)

type QueryProperty struct {
	Type     string
	Property string
}

func GetQuery(operation *openapi3.Operation) map[string]QueryProperty {
	query := make(map[string]QueryProperty)
	parameters := array.Map(operation.Parameters, func(ref *openapi3.ParameterRef) *openapi3.Parameter {
		return ref.Value
	})
	parameters = array.Filter(parameters, func(parameter *openapi3.Parameter) bool {
		return parameter.In == openapi3.ParameterInQuery
	})
	array.ForEach(parameters, func(parameter *openapi3.Parameter) {
		if parameter.Schema != nil && parameter.Schema.Value != nil {
			schema := parameter.Schema.Value
			query[GetPropertyName(parameter.Name)] = QueryProperty{
				Type:     GetPropertyType(schema.Type),
				Property: parameter.Name,
			}

		}
	})
	return query
}
