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

func (p *SchemaParser) GetSchemas() []*entity.Schema {
	schemas := make([]*entity.Schema, len(p.schemas))
	index := 0
	for _, schema := range p.schemas {
		schemas[index] = schema
		index++
	}
	return schemas
}

func (p *SchemaParser) Parse(doc *openapi3.T) {
	for name, schemaRef := range doc.Components.Schemas {
		p.Add(name, schemaRef)
	}
}

func (p *SchemaParser) Add(name string, schemaRef *openapi3.SchemaRef) *entity.Schema {
	schema := schemaRef.Value
	ref := schemaRef.Ref
	switch schema.Type {
	case "boolean":
		return p.AddBoolean(ref, name, schema)
	case "string":
		return p.AddString(ref, name, schema)
	case "number":
		return p.AddFloat(ref, name, schema)
	case "integer":
		return p.AddInteger(ref, name, schema)
	case "object":
		return p.AddObject(ref, name, schema)
	case "array":
		return p.AddArray(ref, name, schema)
	}
	return nil
}
