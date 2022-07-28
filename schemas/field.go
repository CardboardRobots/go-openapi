package schemas

type Field struct {
	Name string
	Type Type
	Tag  string
}

func (f *Field) String() string {
	if f.Tag != "" {
		return Line(1, f.Name, f.Type.String(), Tag(f.Tag))
	} else {
		return Line(1, f.Name, f.Type.String())
	}
}
