package entity

type Endpoint struct {
	Name     string
	Path     string
	Verb     Verb
	Params   []ParamProperty
	Query    map[string]QueryProperty
	Body     map[string]BodyProperty
	Response map[string]ResponseOption
}

type ParamProperty struct {
	Name string
	Key  string
	Type string
}

type QueryProperty struct {
	Name string
	Key  string
	Type string
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
