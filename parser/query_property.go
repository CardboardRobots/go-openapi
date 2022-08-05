package parser

import (
	"github.com/CardboardRobots/go-array"
	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

func (p *SchemaParser) GetQuery(operation *openapi3.Operation) map[string]entity.QueryProperty {
	query := make(map[string]entity.QueryProperty)
	parameters := array.Map(operation.Parameters, func(ref *openapi3.ParameterRef) *openapi3.Parameter {
		return ref.Value
	})
	parameters = array.Filter(parameters, func(parameter *openapi3.Parameter) bool {
		return parameter.In == openapi3.ParameterInQuery
	})
	array.ForEach(parameters, func(parameter *openapi3.Parameter) {
		if parameter.Schema != nil && parameter.Schema.Value != nil {
			name := GetPropertyName(parameter.Name)
			schema := p.Add(name, parameter.Schema, false)
			query[name] = entity.QueryProperty{
				Schema: schema,
				Name:   parameter.Name,
				Key:    "",
			}
		}
	})
	return query
}
