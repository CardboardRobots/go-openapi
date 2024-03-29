package parser

import (
	"net/http"
	"sort"
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
						schema = p.add(GetPropertyName(operation.OperationID+code), mediaType.Schema, true)
					}
					if schema == nil {
						// TODO: Handle this error
					} else {
						responseOption := entity.ResponseOption{
							Name:     schema.Name,
							Type:     schema,
							Code:     GetStatus(code),
							Redirect: IsRediret(code),
						}
						responseOptions = append(responseOptions, responseOption)
					}
				}
			}
		}
	}

	sort.Slice(responseOptions, func(i, j int) bool {
		return responseOptions[i].Name < responseOptions[j].Name
	})

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

func IsRediret(status string) bool {
	code := GetStatus(status)
	if code >= 300 && code < 400 {
		return true
	} else {
		return false
	}
}
