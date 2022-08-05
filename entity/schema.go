package entity

import "sort"

type Schema struct {
	Id      string
	Name    string
	Type    SchemaType
	Fields  []Field
	Items   *Schema
	Display bool
}

func (s *Schema) Sort() {
	sort.Slice(s.Fields, func(i, j int) bool {
		return s.Fields[i].Name < s.Fields[j].Name
	})
}

func (s *Schema) Show(display bool) {
	if display {
		s.Display = true
	}
}

func NewBooleanSchema(
	id string,
	name string,
	display bool,
) Schema {
	return Schema{
		Id:      id,
		Name:    name,
		Type:    TYPE_BOOLEAN,
		Display: display,
	}
}

func NewFloatSchema(
	id string,
	name string,
	display bool,
) Schema {
	return Schema{
		Id:      id,
		Name:    name,
		Type:    TYPE_NUMBER,
		Display: display,
	}
}

func NewIntegerSchema(
	id string,
	name string,
	display bool,
) Schema {
	return Schema{
		Id:      id,
		Name:    name,
		Type:    TYPE_INTEGER,
		Display: display,
	}
}

func NewStringSchema(
	id string,
	name string,
	display bool,
) Schema {
	return Schema{
		Id:      id,
		Name:    name,
		Type:    TYPE_STRING,
		Display: display,
	}
}

func NewObjectSchema(
	id string,
	name string,
	fields []Field,
	display bool,
) Schema {
	return Schema{
		Id:      id,
		Name:    name,
		Type:    TYPE_OBJECT,
		Fields:  fields,
		Display: display,
	}
}

func NewArraySchema(
	id string,
	name string,
	items *Schema,
	display bool,
) Schema {
	return Schema{
		Id:      id,
		Name:    name,
		Type:    TYPE_ARRAY,
		Items:   items,
		Display: display,
	}
}
