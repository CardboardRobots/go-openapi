package core

import (
	"strings"

	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

type SchemaParser struct {
	schemas   map[string]*entity.Schema
	endpoints []*entity.Endpoint
}

func NewSchemaParser() SchemaParser {
	return SchemaParser{
		schemas:   make(map[string]*entity.Schema),
		endpoints: make([]*entity.Endpoint, 0),
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

func (p *SchemaParser) GetEndpoints() []*entity.Endpoint {
	return p.endpoints
}

func (p *SchemaParser) Parse(doc *openapi3.T) {
	for name, schemaRef := range doc.Components.Schemas {
		p.Add(name, schemaRef)
	}
	for key, path := range doc.Paths {
		p.AddEndpoint(key, path, p.schemas)
	}
}

func (p *SchemaParser) AddEndpoint(key string, path *openapi3.PathItem, s map[string]*entity.Schema) {
	// fmt.Printf("path: %v\n", key)
	// printOperation("Connect", path.Connect)
	// printOperation("Delete", path.Delete)
	endpoint := p.GetEndpoint(key, path.Get)
	p.endpoints = append(p.endpoints, &endpoint)
	// printOperation("Head", path.Head)
	// printOperation("Options", path.Options)
	// printOperation("Patch", path.Patch)
	// printOperation("Post", path.Post)
	// printOperation("Put", path.Put)
	// printOperation("Trace", path.Trace)
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

func (p *SchemaParser) GetEndpoint(key string, operation *openapi3.Operation) entity.Endpoint {
	return entity.Endpoint{
		Name:     GetPropertyName(operation.OperationID),
		Path:     KeyToPath(key),
		Params:   GetParams(operation),
		Query:    GetQuery(operation),
		Body:     GetBody(operation),
		Response: GetResponses(operation, p.schemas),
	}
}

func KeyToPath(key string) string {
	key = strings.Replace(key, "}", "", -1)
	key = strings.Replace(key, "{", ":", -1)
	return key
}
