package core

import (
	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

type ResponseOption struct {
	Type string
}

func GetResponses(operation *openapi3.Operation, s map[string]*entity.Schema) map[string]ResponseOption {
	responseOptions := make(map[string]ResponseOption)

	for code, ref := range operation.Responses {
		response := ref.Value
		for key, mediaType := range response.Content {
			if mediaType.Schema != nil {
				schema, ok := s[mediaType.Schema.Ref]
				if ok {
					name := GetResponseName(code, key)
					responseOptions[name] = ResponseOption{
						Type: schema.Name,
					}
				}
			}
		}
	}

	return responseOptions
}

func GetResponseName(code string, key string) string {
	switch key {
	case "application/json":
		key = "Json"
	}
	return GetPropertyName(key) + GetPropertyName(code)
}
