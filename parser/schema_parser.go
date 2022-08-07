package parser

import (
	"sort"
	"strings"

	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

type SchemaParser struct {
	schemasMap map[*openapi3.Schema]*entity.Schema
	schemas    []*entity.Schema
	endpoints  []*entity.Endpoint
}

func NewSchemaParser() SchemaParser {
	return SchemaParser{
		schemasMap: make(map[*openapi3.Schema]*entity.Schema),
		endpoints:  make([]*entity.Endpoint, 0),
	}
}

func (p *SchemaParser) GetBySchema(oapiSchema *openapi3.Schema) *entity.Schema {
	schema, ok := p.schemasMap[oapiSchema]
	if !ok {
		return nil
	}
	return schema
}

func (p *SchemaParser) SetByName(oapiSchema *openapi3.Schema, schema *entity.Schema) {
	p.schemasMap[oapiSchema] = schema
}

func (p *SchemaParser) GetSchemas() []*entity.Schema {
	return p.schemas
}

func (p *SchemaParser) GetEndpoints() []*entity.Endpoint {
	return p.endpoints
}

func (p *SchemaParser) Parse(doc *openapi3.T) {
	for name, schemaRef := range doc.Components.Schemas {
		p.Add(name, schemaRef, true)
	}
	for key, path := range doc.Paths {
		p.AddEndpoint(key, path)
	}
	p.createSortedSchemas()
	p.Sort()
}

func (p *SchemaParser) createSortedSchemas() {
	p.schemas = make([]*entity.Schema, len(p.schemasMap))
	index := 0
	for _, schema := range p.schemasMap {
		p.schemas[index] = schema
		index++
	}
}

func (p *SchemaParser) Sort() {
	sort.Slice(p.schemas, func(i, j int) bool {
		return p.schemas[i].Name < p.schemas[j].Name
	})

	for _, schema := range p.schemas {
		schema.Sort()
	}

	sort.Slice(p.endpoints, func(i, j int) bool {
		return p.endpoints[i].Name < p.endpoints[j].Name
	})
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

func (p *SchemaParser) Add(name string, schemaRef *openapi3.SchemaRef, display bool) *entity.Schema {
	schema := schemaRef.Value
	ref := schemaRef.Ref
	if name == "" {
		name = GetSchemaName(ref)
	}
	name = GetPropertyName(name)

	switch schema.Type {
	case "boolean":
		return p.AddBoolean(ref, name, schema, display)
	case "string":
		return p.AddString(ref, name, schema, display)
	case "number":
		return p.AddFloat(ref, name, schema, display)
	case "integer":
		return p.AddInteger(ref, name, schema, display)
	case "object":
		return p.AddObject(ref, name, schema, true)
	case "array":
		return p.AddArray(ref, name, schema, true)
	}
	return nil
}

func (p *SchemaParser) CreateEndpoint(key string, verb entity.Verb, operation *openapi3.Operation) *entity.Endpoint {
	endpoint := &entity.Endpoint{
		Verb:     verb,
		Name:     GetEndpointName(verb, operation.OperationID, key),
		Path:     KeyToPath(key),
		Params:   GetParams(operation),
		Query:    p.GetQuery(operation),
		Header:   GetHeader(operation),
		Body:     p.GetBody(operation),
		Response: p.GetResponses(operation, p.schemasMap),
	}
	p.endpoints = append(p.endpoints, endpoint)
	return endpoint
}

func GetEndpointName(verb entity.Verb, operationId string, key string) string {
	if operationId != "" {
		return GetPropertyName(operationId)
	}
	return KeyToName(verb, key)
}

func KeyToName(verb entity.Verb, key string) string {
	key = strings.Replace(key, "}", "/", -1)
	key = strings.Replace(key, "{", "By/", -1)
	parts := strings.Split(key, "/")
	for index, part := range parts {
		parts[index] = capitalize(part)
	}
	prefix := capitalize(strings.ToLower(string(verb)))
	return prefix + strings.Join(parts, "")
}

func KeyToPath(key string) string {
	key = strings.Replace(key, "}", "", -1)
	key = strings.Replace(key, "{", ":", -1)
	return key
}
