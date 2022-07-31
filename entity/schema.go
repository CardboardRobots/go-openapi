package entity

type Schema struct {
	Id     string
	Name   string
	Type   SchemaType
	Fields []Field
	Items  *Schema
}

type Field struct {
	Name string
	Type *Schema
	Tag  string
}

func NewBooleanSchema(
	id string,
	name string,
) Schema {
	return Schema{
		Id:   id,
		Name: name,
		Type: TYPE_BOOLEAN,
	}
}

func NewFloatSchema(
	id string,
	name string,
) Schema {
	return Schema{
		Id:   id,
		Name: name,
		Type: TYPE_NUMBER,
	}
}

func NewIntegerSchema(
	id string,
	name string,
) Schema {
	return Schema{
		Id:   id,
		Name: name,
		Type: TYPE_INTEGER,
	}
}

func NewStringSchema(
	id string,
	name string,
) Schema {
	return Schema{
		Id:   id,
		Name: name,
		Type: TYPE_STRING,
	}
}

func NewObjectSchema(
	id string,
	name string,
	fields []Field,
) *Schema {
	return &Schema{
		Id:     id,
		Name:   name,
		Type:   TYPE_OBJECT,
		Fields: fields,
	}
}

func NewArraySchema(
	id string,
	name string,
	items *Schema,
) Schema {
	return Schema{
		Id:    id,
		Name:  name,
		Type:  TYPE_ARRAY,
		Items: items,
	}
}
