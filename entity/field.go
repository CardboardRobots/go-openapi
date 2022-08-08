package entity

import "fmt"

type Field struct {
	*Schema
	Name     string
	TypeName string
	Tag      string
}

func NewField(name string, schema *Schema, tag string) Field {
	typeName := ""
	switch schema.Type {
	case TYPE_BOOLEAN:
		if schema.Enum != nil {
			typeName = schema.Name
		} else {
			typeName = string(schema.Type)
		}
	case TYPE_INTEGER:
		if schema.Enum != nil {
			typeName = schema.Name
		} else {
			typeName = string(schema.Type)
		}
	case TYPE_NUMBER:
		if schema.Enum != nil {
			typeName = schema.Name
		} else {
			typeName = string(schema.Type)
		}
	case TYPE_STRING:
		if schema.Enum != nil {
			typeName = schema.Name
		} else {
			typeName = string(schema.Type)
		}
	case TYPE_ARRAY:
		typeName = fmt.Sprintf("[]%v", schema.Items.Name)
	case TYPE_OBJECT:
		typeName = schema.Name
	}
	return Field{
		Schema:   schema,
		Name:     name,
		TypeName: typeName,
		Tag:      tag,
	}
}
