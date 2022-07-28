package core

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/CardboardRobots/go-array"
	"github.com/cardboardrobots/go-openapi/schemas"
)

func ParseDocument(ctx context.Context) {
	loader := openapi3.Loader{Context: ctx}

	fileName := "openapi.yml"
	log.Printf("Loading %v...\n", fileName)
	doc, err := loader.LoadFromFile(fileName)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	err = doc.Validate(ctx)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	t, err := template.ParseFS(os.DirFS("./templates"), "route.tmpl")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	fmt.Println("package gen")
	fmt.Println("import \"github.com/gin-gonic/gin\"")

	schemaNames := make(map[string]*schemas.Struct)

	for key, schemaRef := range doc.Components.Schemas {
		schema := schemaRef.Value
		switch schema.Type {
		case "string":
		case "number":
		case "integer":
		case "object":
			s := schemas.NewStruct(key, schema)
			name := "#/components/schemas/" + s.Name
			schemaNames[name] = &s
			fmt.Printf("// %v\n%v ", name, s.String())
		case "array":
		}
	}

	// createTemplate("openapi")
	fmt.Println("func Route(router *gin.Engine) {")
	for key, path := range doc.Paths {
		printPath(key, path, schemaNames, t)
	}
	fmt.Println("}")

}

type Types map[string]interface{}

func ParseType(
	name string,
	schema openapi3.Schema,
	pointer bool,
	types Types,
) interface{} {
	return nil
}

func printPath(key string, path *openapi3.PathItem, s map[string]*schemas.Struct, t *template.Template) {
	// fmt.Printf("path: %v\n", key)
	// printOperation("Connect", path.Connect)
	// printOperation("Delete", path.Delete)
	printGet(key, path.Get, s, t)
	// printOperation("Head", path.Head)
	// printOperation("Options", path.Options)
	// printOperation("Patch", path.Patch)
	// printOperation("Post", path.Post)
	// printOperation("Put", path.Put)
	// printOperation("Trace", path.Trace)
}

func printOperation(verb string, operation *openapi3.Operation) {
	if operation != nil {
		fmt.Printf("//\toperation: %v, %#v\n", verb, operation.Description)
		for key, responseRef := range operation.Responses {
			response := responseRef.Value
			if response != nil {
				description := *response.Description
				fmt.Printf("\t\tresponse: %v, %v\n", key, description)
				for key, content := range response.Content {
					fmt.Printf("\t\t\tcontent: %v, %v\n", key, content.Schema)
				}
			}
		}
	}
}

func printGet(key string, operation *openapi3.Operation, s map[string]*schemas.Struct, t *template.Template) {
	parameters := make([]*openapi3.Parameter, 0)
	for _, parameterRef := range operation.Parameters {
		parameter := parameterRef.Value
		if parameter.In == openapi3.ParameterInPath {
			parameters = append(parameters, parameter)
		}
	}

	t.Execute(os.Stdout, map[string]interface{}{
		"Path":       KeyToPath(key),
		"Operation":  operation,
		"Parameters": parameters,
		"Query":      GetQuery(operation),
		"Responses":  GetResponses(operation, s),
	})
}

func KeyToPath(key string) string {
	key = strings.Replace(key, "}", "", -1)
	key = strings.Replace(key, "{", ":", -1)
	return key
}

type QueryProperty struct {
	Type     string
	Property string
}

func GetQuery(operation *openapi3.Operation) map[string]QueryProperty {
	query := make(map[string]QueryProperty)
	parameters := array.Map(operation.Parameters, func(ref *openapi3.ParameterRef) *openapi3.Parameter {
		return ref.Value
	})
	parameters = array.Filter(parameters, func(parameter *openapi3.Parameter) bool {
		return parameter.In == openapi3.ParameterInQuery
	})
	array.ForEach(parameters, func(parameter *openapi3.Parameter) {
		if parameter.Schema != nil && parameter.Schema.Value != nil {
			schema := parameter.Schema.Value
			query[GetPropertyName(parameter.Name)] = QueryProperty{
				Type:     GetPropertyType(schema.Type),
				Property: parameter.Name,
			}

		}
	})
	return query
}

func GetPropertyName(name string) string {
	if len(name) < 1 {
		return ""
	}
	first := name[:1]
	rest := name[1:]
	return strings.ToUpper(first) + rest
}

func GetPropertyType(name string) string {
	switch name {
	case "number":
		return "float32"
	case "integer":
		return "int"
	}
	return name
}

type ResponseOption struct {
	Type string
}

func GetResponses(operation *openapi3.Operation, s map[string]*schemas.Struct) map[string]ResponseOption {
	responseOptions := make(map[string]ResponseOption)

	for code, ref := range operation.Responses {
		response := ref.Value
		for key, mediaType := range response.Content {
			if mediaType.Schema != nil {
				schema, ok := s[mediaType.Schema.Ref]
				if ok {
					name := GetResponseName(code, key, operation)
					responseOptions[name] = ResponseOption{
						Type: schema.Name,
					}
				}
			}
		}
	}

	return responseOptions
}

func GetResponseName(code string, key string, operation *openapi3.Operation) string {
	switch key {
	case "application/json":
		key = "Json"
	}
	return GetPropertyName(operation.OperationID) + GetPropertyName(code) + GetPropertyName(key)
}
