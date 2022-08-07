package parser

import (
	"fmt"
	"strings"

	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

func (p *SchemaParser) AddBoolean(
	ref string,
	name string,
	schema *openapi3.Schema,
	display bool,
) *entity.Schema {
	item := p.GetBySchema(schema)
	if item != nil {
		item.Show(display)
		return item
	}

	newItem := entity.NewBooleanSchema(ref, name, display)

	p.SetByName(schema, &newItem)
	return &newItem
}

func (p *SchemaParser) AddInteger(
	ref string,
	name string,
	schema *openapi3.Schema,
	display bool,
) *entity.Schema {
	item := p.GetBySchema(schema)
	if item != nil {
		item.Show(display)
		return item
	}

	newItem := entity.NewIntegerSchema(ref, name, display)

	p.SetByName(schema, &newItem)
	return &newItem
}

func (p *SchemaParser) AddFloat(
	ref string,
	name string,
	schema *openapi3.Schema,
	display bool,
) *entity.Schema {
	item := p.GetBySchema(schema)
	if item != nil {
		item.Show(display)
		return item
	}

	newItem := entity.NewFloatSchema(ref, name, display)

	p.SetByName(schema, &newItem)
	return &newItem
}

func (p *SchemaParser) AddString(
	ref string,
	name string,
	schema *openapi3.Schema,
	display bool,
) *entity.Schema {
	item := p.GetBySchema(schema)
	if item != nil {
		item.Show(display)
		return item
	}

	var newItem entity.Schema
	if len(schema.Enum) > 0 {
		options := make([]entity.EnumOption, len(schema.Enum))
		for index, value := range schema.Enum {
			stringValue, ok := value.(string)
			key := strings.ToUpper(fmt.Sprintf("%v_%v", name, stringValue))
			if ok {
				options[index] = entity.EnumOption{
					Key:   key,
					Value: stringValue,
				}
			} else {
				// TODO: Fix this case
				options[index] = entity.EnumOption{}
			}
		}
		newItem = entity.NewEnumSchema(ref, name, true, &entity.Enum{
			Options: options,
		})
	} else {
		newItem = entity.NewStringSchema(ref, name, display)
	}

	p.SetByName(schema, &newItem)
	return &newItem
}

func (p *SchemaParser) AddObject(
	ref string,
	name string,
	schema *openapi3.Schema,
	display bool,
) *entity.Schema {
	item := p.GetBySchema(schema)
	if item != nil {
		item.Show(display)
		return item
	}

	fields := make([]entity.Field, len(schema.Properties))
	index := 0
	for key, schemaRef := range schema.Properties {
		// name := schemaRef.Ref
		// if name is empty, this is not a true ref
		name := GetPropertyName(key)
		schema := p.Add(key, schemaRef, false)
		fields[index] = entity.NewField(
			name,
			schema,
			key,
		)
		index++
	}

	newItem := entity.NewObjectSchema(ref, name, fields, display)

	p.SetByName(schema, &newItem)
	return &newItem
}

func (p *SchemaParser) AddArray(
	ref string,
	name string,
	schema *openapi3.Schema,
	display bool,
) *entity.Schema {
	item := p.GetBySchema(schema)
	if item != nil {
		item.Show(display)
		return item
	}

	items := p.Add("", schema.Items, false)

	newItem := entity.NewArraySchema(ref, name, items, display)

	p.SetByName(schema, &newItem)
	return &newItem
}
