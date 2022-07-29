package core

import "github.com/getkin/kin-openapi/openapi3"

type BodyProperty struct {
	Type     string
	Property string
}

func GetBody(operation *openapi3.Operation) map[string]BodyProperty {
	body := make(map[string]BodyProperty)
	return body
}
