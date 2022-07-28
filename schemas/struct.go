package schemas

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type Struct struct {
	Name   string
	Fields []Field
}

func (s *Struct) String() string {
	sb := strings.Builder{}
	sb.WriteString(Line(0, "type", s.Name, "struct {"))
	for _, field := range s.Fields {
		sb.WriteString(field.String())
	}
	sb.WriteString("}\n\n")
	return sb.String()
}

func NewStruct(name string, schema *openapi3.Schema) Struct {
	fields := make([]Field, len(schema.Properties))
	index := 0
	for key, schemaRef := range schema.Properties {
		schema := schemaRef.Value
		fields[index] = Field{
			Name: key,
			Type: NewType(schema.Type, false),
		}
		index++
	}
	return Struct{
		Name:   name,
		Fields: fields,
	}
}
