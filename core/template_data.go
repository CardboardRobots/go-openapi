package core

import "github.com/cardboardrobots/go-openapi/entity"

type TemplateData struct {
	Package   string
	Structs   []*entity.Schema
	Endpoints []*entity.Endpoint
}
