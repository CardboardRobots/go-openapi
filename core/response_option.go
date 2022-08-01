package core

import (
	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

func GetResponses(operation *openapi3.Operation, s map[string]*entity.Schema) map[string]entity.ResponseOption {
	responseOptions := make(map[string]entity.ResponseOption)

	for code, ref := range operation.Responses {
		response := ref.Value
		for key, mediaType := range response.Content {
			if mediaType.Schema != nil {
				name := GetSchemaName(mediaType.Schema.Ref)
				schema, ok := s[name]
				if ok {
					name := GetResponseName(code, key)
					responseOptions[name] = entity.ResponseOption{
						Name: schema.Name,
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
