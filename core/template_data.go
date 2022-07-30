package core

import "github.com/cardboardrobots/go-openapi/entity"

type TemplateData struct {
	Package   string
	Structs   []*entity.Struct
	Endpoints []Endpoint
}
