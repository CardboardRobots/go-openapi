package entity

type TemplateData struct {
	Package   string
	Structs   []*Schema
	Endpoints []*Endpoint
}
