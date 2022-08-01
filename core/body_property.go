package core

import (
	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

func GetBody(operation *openapi3.Operation) map[string]entity.BodyProperty {
	body := make(map[string]entity.BodyProperty)
	return body
}
