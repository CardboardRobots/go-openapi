package entity

type Endpoint struct {
	Name     string
	Path     string
	Params   []ParamProperty
	Query    []QueryProperty
	Body     []BodyProperty
	Response []ResponseOption
}

type ParamProperty struct {
	Name string
	Key  string
	Type *Schema
}

type QueryProperty struct {
	Name string
	Key  string
	Type *Schema
}

type BodyProperty struct {
	Name string
	Key  string
	Type *Schema
}

type ResponseOption struct {
	Code int
	Name string
	Type *Schema
}
