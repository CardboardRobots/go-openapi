package entity

type Field struct {
	*Schema
	Name string
	Tag  string
}

func NewField(name string, schema *Schema, tag string) Field {
	return Field{
		Name:   name,
		Schema: schema,
		Tag:    tag,
	}
}
