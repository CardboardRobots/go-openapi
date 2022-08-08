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
		typeName = "bool"
	case TYPE_INTEGER:
		typeName = "int"
	case TYPE_NUMBER:
		typeName = "float32"
	case TYPE_STRING:
		typeName = "string"
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
