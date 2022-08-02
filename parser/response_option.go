package parser

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

func (p *SchemaParser) GetResponses(operation *openapi3.Operation, s map[*openapi3.Schema]*entity.Schema) entity.Response {
	responseOptions := make([]entity.ResponseOption, 0)
	r := entity.Response{}

	for code, ref := range operation.Responses {
		response := ref.Value
		if len(response.Content) == 0 {
			// Null response
			r.Default = true
			r.DefaultCode = GetStatus(code)
		} else {
			for _, mediaType := range response.Content {
				if mediaType.Schema != nil {
					schema, ok := s[mediaType.Schema.Value]
					if !ok {
						// Schema has not be logged
						schema = p.Add(GetPropertyName(operation.OperationID+code), mediaType.Schema)
					}

					responseOption := entity.ResponseOption{
						Name: schema.Name,
						Type: schema,
						Code: GetStatus(code),
					}
					responseOptions = append(responseOptions, responseOption)
				}
			}
		}
	}

	r.Options = responseOptions
	return r
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
