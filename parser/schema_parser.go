package parser

import (
	"fmt"
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
		p.add(name, schemaRef, true)
	}
	security := p.addSecurity(doc)
	for key, path := range doc.Paths {
		p.addEndpoint(doc, key, path, security)
	}
	p.sortSchemas()
	p.sortEndpoints()
}

func (p *SchemaParser) addSecurity(doc *openapi3.T) []*openapi3.SecurityScheme {
	security := make([]*openapi3.SecurityScheme, 0)
	for _, securityRequirement := range doc.Security {
		for name := range securityRequirement {
			securitySchema, ok := doc.Components.SecuritySchemes[name]
			if ok && securitySchema.Value != nil {
				security = append(security, securitySchema.Value)
			}
		}
	}
	return security
}

func (p *SchemaParser) sortSchemas() {
	p.schemas = make([]*entity.Schema, 0)

	// Filter schemas
	for _, schema := range p.schemasMap {
		if schema.Display {
			p.schemas = append(p.schemas, schema)
		}
	}

	sort.Slice(p.schemas, func(i, j int) bool {
		return p.schemas[i].Name < p.schemas[j].Name
	})

	for _, schema := range p.schemas {
		schema.Sort()
	}
}

func (p *SchemaParser) sortEndpoints() {
	sort.Slice(p.endpoints, func(i, j int) bool {
		return p.endpoints[i].Name < p.endpoints[j].Name
	})
}

func (p *SchemaParser) addEndpoint(
	doc *openapi3.T,
	key string,
	path *openapi3.PathItem,
	security []*openapi3.SecurityScheme,
) {
	if path.Connect != nil {
		p.createEndpoint(doc, key, entity.VERB_CONNECT, path.Connect, security)
	}
	if path.Delete != nil {
		p.createEndpoint(doc, key, entity.VERB_DELETE, path.Delete, security)
	}
	if path.Get != nil {
		p.createEndpoint(doc, key, entity.VERB_GET, path.Get, security)
	}
	if path.Head != nil {
		p.createEndpoint(doc, key, entity.VERB_HEAD, path.Head, security)
	}
	if path.Options != nil {
		p.createEndpoint(doc, key, entity.VERB_OPTIONS, path.Options, security)
	}
	if path.Patch != nil {
		p.createEndpoint(doc, key, entity.VERB_PATCH, path.Patch, security)
	}
	if path.Post != nil {
		p.createEndpoint(doc, key, entity.VERB_POST, path.Post, security)
	}
	if path.Put != nil {
		p.createEndpoint(doc, key, entity.VERB_PUT, path.Put, security)
	}
	if path.Trace != nil {
		p.createEndpoint(doc, key, entity.VERB_TRACE, path.Trace, security)
	}
}

func (p *SchemaParser) add(key string, schemaRef *openapi3.SchemaRef, display bool) *entity.Schema {
	schema := schemaRef.Value
	ref := schemaRef.Ref
	name := GetSchemaName(ref)
	if name == "" {
		name = key
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
		return p.AddArray(ref, name, schema, display)
	}
	return nil
}

func (p *SchemaParser) createEndpoint(
	doc *openapi3.T,
	key string,
	verb entity.Verb,
	operation *openapi3.Operation,
	security []*openapi3.SecurityScheme,
) *entity.Endpoint {
	securitySchemes := p.createSecuritySchemes(doc, operation, security)
	endpoint := &entity.Endpoint{
		Verb:     verb,
		Name:     GetEndpointName(verb, operation.OperationID, key),
		Path:     KeyToPath(key),
		Params:   GetParams(operation),
		Query:    p.GetQuery(operation),
		Header:   GetHeader(operation, securitySchemes),
		Body:     p.GetBody(operation),
		Response: p.GetResponses(operation, p.schemasMap),
		Security: securitySchemes,
	}
	p.endpoints = append(p.endpoints, endpoint)
	return endpoint
}

func (p *SchemaParser) createSecuritySchemes(
	doc *openapi3.T,
	operation *openapi3.Operation,
	security []*openapi3.SecurityScheme,
) []entity.Security {
	schemes := make(map[*openapi3.SecurityScheme]entity.Security)

	for _, scheme := range security {
		schemes[scheme] = p.createSecurity(scheme)
	}

	if operation.Security != nil {
		for _, securityRequirement := range *operation.Security {
			for name := range securityRequirement {
				securitySchema, ok := doc.Components.SecuritySchemes[name]
				if ok && securitySchema.Value != nil {
					scheme := securitySchema.Value
					schemes[scheme] = p.createSecurity(scheme)
				}
			}
		}
	}

	result := make([]entity.Security, len(schemes))
	index := 0
	for _, scheme := range schemes {
		result[index] = scheme
		index++
	}
	fmt.Printf("%v", result)

	return result
}

func (p *SchemaParser) createSecurity(scheme *openapi3.SecurityScheme) entity.Security {
	switch strings.ToLower(scheme.Type) {
	case "http":
		switch strings.ToLower(scheme.Scheme) {
		case "basic":
			return entity.Security{
				Type: entity.SECURITY_TYPE_BASIC,
			}
		case "bearer":
			return entity.Security{
				Type: entity.SECURITY_TYPE_BEARER,
			}
		default:
			return entity.Security{
				Type: entity.SECURITY_TYPE_BASIC,
			}
		}
	case "oauth2":
		return entity.Security{
			Type: entity.SECURITY_TYPE_BEARER,
		}
	default:
		return entity.Security{
			Type: entity.SECURITY_TYPE_BASIC,
		}
	}
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
