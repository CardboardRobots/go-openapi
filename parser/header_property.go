package parser

import (
	"github.com/CardboardRobots/go-array"
	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

func GetHeader(
	operation *openapi3.Operation,
	securitySchemes []entity.Security,
) map[string]entity.HeaderProperty {
	header := make(map[string]entity.HeaderProperty)
	parameters := array.Map(operation.Parameters, func(ref *openapi3.ParameterRef) *openapi3.Parameter {
		return ref.Value
	})
	parameters = array.Filter(parameters, func(parameter *openapi3.Parameter) bool {
		return parameter.In == openapi3.ParameterInHeader
	})
	array.ForEach(parameters, func(parameter *openapi3.Parameter) {
		if parameter.Schema != nil && parameter.Schema.Value != nil {
			schema := parameter.Schema.Value
			name := GetPropertyName(parameter.Name)
			header[name] = entity.HeaderProperty{
				Type: GetPropertyType(schema.Type),
				Name: name,
				Key:  parameter.Name,
			}
		}
	})
	array.ForEach(securitySchemes, func(securityScheme entity.Security) {
		authorization := "authorization"
		Authorization := "Authorization"
		switch securityScheme.Type {
		case entity.SECURITY_TYPE_BASIC:
			header[Authorization] = entity.HeaderProperty{
				Type: "string",
				Name: Authorization,
				Key:  authorization,
			}
		case entity.SECURITY_TYPE_BEARER:
			header[Authorization] = entity.HeaderProperty{
				Type: "string",
				Name: Authorization,
				Key:  authorization,
			}
		case entity.SECURITY_TYPE_COOKIE:
		default:
		}
	})
	return header
}
