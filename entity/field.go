package entity

type Field struct {
	Name string
	Type string
	Tag  string
}

func NewField(name string, Type string, tag string) Field {
	return Field{
		Name: name,
		Type: Type,
		Tag:  tag,
	}
}
