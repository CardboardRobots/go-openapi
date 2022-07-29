package core

import "github.com/cardboardrobots/go-openapi/schemas"

type TemplateData struct {
	Structs   []*schemas.Struct
	Endpoints []Endpoint
}
