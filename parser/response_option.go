package parser

import (
	"net/http"
	"strconv"
	"strings"

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
						Type: s[schema.Name],
						Code: GetStatus(code),
					}
				}
			}
		}
	}

	return responseOptions
}

func GetStatus(status string) int {
	status = strings.ToLower(status)
	if status == "default" {
		return http.StatusBadRequest
	}
	code, err := strconv.Atoi(status)
	if err != nil {
		return http.StatusBadRequest
	}
	return code
}

func GetResponseName(code string, key string) string {
	switch key {
	case "application/json":
		key = "Json"
	}
	return GetPropertyName(key) + GetPropertyName(code)
}
