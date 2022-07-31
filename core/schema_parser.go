package core

import "github.com/cardboardrobots/go-openapi/entity"

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
