package entity

type Schema struct {
	Id     string
	Name   string
	Type   SchemaType
	Fields []Field
	Slice  bool
}

func NewFloatSchema(
	id string,
	name string,
	slice bool,
) Schema {
	return Schema{
		Id:    id,
		Name:  name,
		Type:  TYPE_NUMBER,
		Slice: slice,
	}
}

func NewIntegerSchema(
	id string,
	name string,
	slice bool,
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
	slice bool,
) Schema {
	return Schema{
		Id:   id,
		Name: name,
		Type: TYPE_STRING,
	}
}

func NewStructSchema(
	id string,
	name string,
	slice bool,
	fields []Field,
) Schema {
	return Schema{
		Id:     id,
		Name:   name,
		Type:   TYPE_OBJECT,
		Fields: fields,
	}
}

func NewArraySchema(
	id string,
	name string,
	slice bool,
)

type Struct struct {
	Name   string
	Fields []Field
}
