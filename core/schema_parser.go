package core

import (
	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

type SchemaParser struct {
	schemas map[string]*entity.Schema
}

func NewSchemaParser() SchemaParser {
	return SchemaParser{
		schemas: make(map[string]*entity.Schema),
	}
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

func (p *SchemaParser) Parse(doc *openapi3.T) {
	for key, schemaRef := range doc.Components.Schemas {
		schema := schemaRef.Value
		switch schema.Type {
		case "boolean":
			p.AddBoolean(key, schema)
		case "string":
			p.AddString(key, schema)
		case "number":
			p.AddFloat(key, schema)
		case "integer":
			p.AddInteger(key, schema)
		case "object":
			p.AddObject(key, schema)
		case "array":
			p.AddArray(key, schema)
		}
	}
}
