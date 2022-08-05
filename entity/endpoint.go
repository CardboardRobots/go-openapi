package entity

type Endpoint struct {
	Name     string
	Path     string
	Verb     Verb
	Params   []ParamProperty
	Query    map[string]QueryProperty
	Header   map[string]HeaderProperty
	Body     BodyProperty
	Response Response
}

type ParamProperty struct {
	Name string
	Key  string
	Type string
}

type QueryProperty struct {
	Name string
	Key  string
	Type *Schema
}

type HeaderProperty struct {
	Name string
	Key  string
	Type string
}

type BodyProperty struct {
	Schema   *Schema
	Encoding Encoding
}

type Encoding string

const ENCODING_JSON Encoding = "json"
const ENCODING_XML Encoding = "xml"
const ENCODING_URL Encoding = "form"
const ENCODING_TEXT Encoding = "json"

type ResponseOption struct {
	Code int
	Name string
	Type *Schema
}

type Response struct {
	Options     []ResponseOption
	Default     bool
	DefaultCode int
}
