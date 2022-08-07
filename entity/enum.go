package entity

type Enum struct {
	Options []EnumOption
}

type EnumOption struct {
	Key   string
	Value string
}
