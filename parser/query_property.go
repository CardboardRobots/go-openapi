package parser

import (
	"github.com/CardboardRobots/go-array"
	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

func GetQuery(operation *openapi3.Operation) map[string]entity.QueryProperty {
	query := make(map[string]entity.QueryProperty)
	parameters := array.Map(operation.Parameters, func(ref *openapi3.ParameterRef) *openapi3.Parameter {
		return ref.Value
	})
	parameters = array.Filter(parameters, func(parameter *openapi3.Parameter) bool {
		return parameter.In == openapi3.ParameterInQuery
	})
	array.ForEach(parameters, func(parameter *openapi3.Parameter) {
		if parameter.Schema != nil && parameter.Schema.Value != nil {
			schema := parameter.Schema.Value
			query[GetPropertyName(parameter.Name)] = entity.QueryProperty{
				Type: GetPropertyType(schema.Type),
				Name: parameter.Name,
				Key:  "",
			}
		}
	})
	return query
}