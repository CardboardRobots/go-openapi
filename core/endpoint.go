package core

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type Endpoint struct {
	OperationId string
	Path        string
	Operation   *openapi3.Operation
	Parameters  []ParamProperty
	Query       map[string]QueryProperty
	Body        map[string]BodyProperty
	Responses   map[string]ResponseOption
}

func KeyToPath(key string) string {
	key = strings.Replace(key, "}", "", -1)
	key = strings.Replace(key, "{", ":", -1)
	return key
}
