package core

import (
	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

type SchemaParser struct {
	schemas map[string]*entity.Schema
}

func (p *SchemaParser) GetById(id string) *entity.Schema {
	schema, ok := p.schemas[id]
	if !ok {
		return nil
	}
	return schema
}

func (p *SchemaParser) SetById(id string, schema *entity.Schema) {
	p.schemas[id] = schema
}

func (p *SchemaParser) AddSchema(id string, schema *openapi3.Schema) entity.Schema {
	return entity.Schema{}
}

func SchemaToBoolean(name string, schema *openapi3.Schema) entity.Schema {
	return entity.NewBooleanSchema(name, name)
}

func (p *SchemaParser) AddBoolean(id string, schema *openapi3.Schema) entity.Schema {
	value := SchemaToBoolean(id, schema)
	p.schemas[id] = &value
	return value
}

func SchemaToInteger(name string, schema *openapi3.Schema) entity.Schema {
	return entity.NewIntegerSchema(name, name)
}

func SchemaToFloat(name string, schema *openapi3.Schema) entity.Schema {
	return entity.NewFloatSchema(name, name)
}

func SchhemaToString(name string, schema *openapi3.Schema) entity.Schema {
	return entity.NewFloatSchema(name, name)
}

func SchemaToObject(name string, schema *openapi3.Schema) entity.Schema {
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
	return entity.NewObjectSchema(name, name, fields)
}

func SchemaToArray(name string, schema *openapi3.Schema, s map[string]*entity.Schema) entity.Schema {
	return entity.NewArraySchema(
		name, name, nil,
	)
}
