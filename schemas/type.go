package schemas

type Type struct {
	Name    string
	Pointer bool
}

func (t *Type) String() string {
	if t.Pointer {
		return "*" + t.Name
	} else {
		return t.Name
	}
}

func NewType(name string, pointer bool) Type {
	switch name {
	case "string":
	case "number":
	case "integer":
		name = "number"
	case "boolean":
	case "array":
	case "object":
	}
	return Type{
		name,
		pointer,
	}
}
