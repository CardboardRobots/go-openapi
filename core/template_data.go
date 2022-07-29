package core

import "github.com/cardboardrobots/go-openapi/schemas"

type TemplateData struct {
	Package   string
	Structs   []*schemas.Struct
	Endpoints []Endpoint
}
