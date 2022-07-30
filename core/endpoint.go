package core

import (
	"strings"
	"text/template"

	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

func GetEndpoint(key string, operation *openapi3.Operation, s map[string]*entity.Struct, t *template.Template) Endpoint {
	return Endpoint{
		OperationId: GetPropertyName(operation.OperationID),
		Path:        KeyToPath(key),
		Operation:   operation,
		Parameters:  GetParams(operation),
		Query:       GetQuery(operation),
		Body:        GetBody(operation),
		Responses:   GetResponses(operation, s),
	}
}

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
