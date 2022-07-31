package core

import (
	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

func (p *SchemaParser) AddBoolean(
	id string,
	schema *openapi3.Schema,
) entity.Schema {
	name := GetSchemaName(id)
	return entity.NewBooleanSchema(id, name)
}

func (p *SchemaParser) AddInteger(
	id string,
	schema *openapi3.Schema,
) entity.Schema {
	name := GetSchemaName(id)
	return entity.NewIntegerSchema(id, name)
}

func (p *SchemaParser) AddFloat(
	id string,
	schema *openapi3.Schema,
) entity.Schema {
	name := GetSchemaName(id)
	return entity.NewFloatSchema(id, name)
}

func (p *SchemaParser) AddString(
	id string,
	schema *openapi3.Schema,
) entity.Schema {
	name := GetSchemaName(id)
	return entity.NewFloatSchema(id, name)
}

func (p *SchemaParser) AddObject(
	id string,
	schema *openapi3.Schema,
) entity.Schema {
	name := GetSchemaName(id)
	fields := make([]entity.Field, len(schema.Properties))
	index := 0
	for key := range schema.Properties {
		// schema := schemaRef.Value
		// name := schemaRef.Ref
		// if name is empty, this is not a true ref
		fields[index] = entity.Field{
			Name: GetPropertyName(key),
			Tag:  key,
			// Type: GetPropertyType(schema.Type),
		}
		index++
	}
	return entity.NewObjectSchema(id, name, fields)
}

func (p *SchemaParser) AddArray(
	id string,
	schema *openapi3.Schema,
	s map[string]*entity.Schema,
) entity.Schema {
	name := GetSchemaName(id)
	return entity.NewArraySchema(
		id, name, nil,
	)
}
