package core

import (
	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

func (p *SchemaParser) AddBoolean(
	id string,
	schema *openapi3.Schema,
) *entity.Schema {
	item := p.GetById(id)
	if item != nil {
		return item
	}

	name := GetSchemaName(id)
	newItem := entity.NewBooleanSchema(id, name)

	p.SetById(id, &newItem)
	return &newItem
}

func (p *SchemaParser) AddInteger(
	id string,
	schema *openapi3.Schema,
) *entity.Schema {
	item := p.GetById(id)
	if item != nil {
		return item
	}

	name := GetSchemaName(id)
	newItem := entity.NewIntegerSchema(id, name)

	p.SetById(id, &newItem)
	return &newItem
}

func (p *SchemaParser) AddFloat(
	id string,
	schema *openapi3.Schema,
) *entity.Schema {
	item := p.GetById(id)
	if item != nil {
		return item
	}

	name := GetSchemaName(id)
	newItem := entity.NewFloatSchema(id, name)

	p.SetById(id, &newItem)
	return &newItem
}

func (p *SchemaParser) AddString(
	id string,
	schema *openapi3.Schema,
) *entity.Schema {
	item := p.GetById(id)
	if item != nil {
		return item
	}

	name := GetSchemaName(id)
	newItem := entity.NewFloatSchema(id, name)

	p.SetById(id, &newItem)
	return &newItem
}

func (p *SchemaParser) AddObject(
	id string,
	schema *openapi3.Schema,
) *entity.Schema {
	item := p.GetById(id)
	if item != nil {
		return item
	}

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

	newItem := entity.NewObjectSchema(id, name, fields)

	p.SetById(id, &newItem)
	return &newItem
}

func (p *SchemaParser) AddArray(
	id string,
	schema *openapi3.Schema,
) *entity.Schema {
	item := p.GetById(id)
	if item != nil {
		return item
	}

	name := GetSchemaName(id)
	newItem := entity.NewArraySchema(id, name, nil)

	p.SetById(id, &newItem)
	return &newItem
}
