package parser

import (
	"strings"

	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

type SchemaParser struct {
	schemas   map[*openapi3.Schema]*entity.Schema
	endpoints []*entity.Endpoint
}

func NewSchemaParser() SchemaParser {
	return SchemaParser{
		schemas:   make(map[*openapi3.Schema]*entity.Schema),
		endpoints: make([]*entity.Endpoint, 0),
	}
}

func (p *SchemaParser) GetBySchema(oapiSchema *openapi3.Schema) *entity.Schema {
	schema, ok := p.schemas[oapiSchema]
	if !ok {
		return nil
	}
	return schema
}

// func (p *SchemaParser) GetByName(name string) *entity.Schema {
// 	schema, ok := p.schemas[name]
// 	if !ok {
// 		return nil
// 	}
// 	return schema
// }

// func (p *SchemaParser) GetByRef(ref string) *entity.Schema {
// 	name := GetSchemaName(ref)
// 	if name == "" {
// 		return nil
// 	}
// 	return p.GetByName(name)
// }

func (p *SchemaParser) SetByName(oapiSchema *openapi3.Schema, schema *entity.Schema) {
	p.schemas[oapiSchema] = schema
}

// func (p *SchemaParser) SetByName(name string, schema *entity.Schema) {
// 	p.schemas[name] = schema
// }

// func (p *SchemaParser) SetByRef(ref string, schema *entity.Schema) error {
// 	name := GetSchemaName(ref)
// 	if name == "" {
// 		return errors.New("unable to parse ref")
// 	}
// 	p.schemas[name] = schema
// 	return nil
// }

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
		p.AddEndpoint(key, path)
	}
}

func (p *SchemaParser) AddEndpoint(key string, path *openapi3.PathItem) {
	if path.Connect != nil {
		p.CreateEndpoint(key, entity.VERB_CONNECT, path.Connect)
	}
	if path.Delete != nil {
		p.CreateEndpoint(key, entity.VERB_DELETE, path.Delete)
	}
	if path.Get != nil {
		p.CreateEndpoint(key, entity.VERB_GET, path.Get)
	}
	if path.Head != nil {
		p.CreateEndpoint(key, entity.VERB_HEAD, path.Head)
	}
	if path.Options != nil {
		p.CreateEndpoint(key, entity.VERB_OPTIONS, path.Options)
	}
	if path.Patch != nil {
		p.CreateEndpoint(key, entity.VERB_PATCH, path.Patch)
	}
	if path.Post != nil {
		p.CreateEndpoint(key, entity.VERB_POST, path.Post)
	}
	if path.Put != nil {
		p.CreateEndpoint(key, entity.VERB_PUT, path.Put)
	}
	if path.Trace != nil {
		p.CreateEndpoint(key, entity.VERB_TRACE, path.Trace)
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

func (p *SchemaParser) CreateEndpoint(key string, verb entity.Verb, operation *openapi3.Operation) *entity.Endpoint {
	endpoint := &entity.Endpoint{
		Verb:     verb,
		Name:     GetPropertyName(operation.OperationID),
		Path:     KeyToPath(key),
		Params:   GetParams(operation),
		Query:    GetQuery(operation),
		Header:   GetHeader(operation),
		Body:     p.GetBody(operation),
		Response: p.GetResponses(operation, p.schemas),
	}
	p.endpoints = append(p.endpoints, endpoint)
	return endpoint
}

func KeyToPath(key string) string {
	key = strings.Replace(key, "}", "", -1)
	key = strings.Replace(key, "{", ":", -1)
	return key
}
