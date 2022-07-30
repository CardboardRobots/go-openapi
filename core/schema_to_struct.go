package core

import (
	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

func SchemaToStruct(name string, schema openapi3.Schema) entity.Struct {
	fields := make([]entity.Field, len(schema.Properties))
	index := 0
	for key, schemaRef := range schema.Properties {
		schema := schemaRef.Value
		fields[index] = entity.Field{
			Name: GetPropertyName(key),
			Tag:  key,
			Type: schema.Type,
		}
		index++
	}
	return entity.Struct{
		Name:   name,
		Fields: fields,
	}
}
