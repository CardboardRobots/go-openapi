package parser

import (
	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

func (p *SchemaParser) AddBoolean(
	ref string,
	name string,
	schema *openapi3.Schema,
) *entity.Schema {
	if name == "" {
		name = GetSchemaName(ref)
	}

	item := p.GetBySchema(schema)
	if item != nil {
		return item
	}

	newItem := entity.NewBooleanSchema(ref, name)

	p.SetByName(schema, &newItem)
	return &newItem
}

func (p *SchemaParser) AddInteger(
	ref string,
	name string,
	schema *openapi3.Schema,
) *entity.Schema {
	if name == "" {
		name = GetSchemaName(ref)
	}

	item := p.GetBySchema(schema)
	if item != nil {
		return item
	}

	newItem := entity.NewIntegerSchema(ref, name)

	p.SetByName(schema, &newItem)
	return &newItem
}

func (p *SchemaParser) AddFloat(
	ref string,
	name string,
	schema *openapi3.Schema,
) *entity.Schema {
	if name == "" {
		name = GetSchemaName(ref)
	}

	item := p.GetBySchema(schema)
	if item != nil {
		return item
	}

	newItem := entity.NewFloatSchema(ref, name)

	p.SetByName(schema, &newItem)
	return &newItem
}

func (p *SchemaParser) AddString(
	ref string,
	name string,
	schema *openapi3.Schema,
) *entity.Schema {
	if name == "" {
		name = GetSchemaName(ref)
	}

	item := p.GetBySchema(schema)
	if item != nil {
		return item
	}

	newItem := entity.NewStringSchema(ref, name)

	p.SetByName(schema, &newItem)
	return &newItem
}

func (p *SchemaParser) AddObject(
	ref string,
	name string,
	schema *openapi3.Schema,
) *entity.Schema {
	if name == "" {
		name = GetSchemaName(ref)
	}
	item := p.GetBySchema(schema)
	if item != nil {
		return item
	}

	fields := make([]entity.Field, len(schema.Properties))
	index := 0
	for key, schemaRef := range schema.Properties {
		schema := schemaRef.Value
		// name := schemaRef.Ref
		// if name is empty, this is not a true ref
		fields[index] = entity.Field{
			Name: GetPropertyName(key),
			Tag:  key,
			Type: GetPropertyType(schema.Type),
		}
		index++
	}

	newItem := entity.NewObjectSchema(ref, name, fields)

	p.SetByName(schema, &newItem)
	return &newItem
}

func (p *SchemaParser) AddArray(
	ref string,
	name string,
	schema *openapi3.Schema,
) *entity.Schema {
	if name == "" {
		name = GetSchemaName(ref)
	}

	item := p.GetBySchema(schema)
	if item != nil {
		return item
	}

	items := p.Add("", schema.Items)

	newItem := entity.NewArraySchema(ref, name, items)

	p.SetByName(schema, &newItem)
	return &newItem
}
