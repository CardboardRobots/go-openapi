package core

import (
	"bufio"
	"context"
	"io/fs"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/CardboardRobots/go-array"
	"github.com/cardboardrobots/go-openapi/schemas"
)

func ParseDocument(ctx context.Context, fsys fs.FS) {
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

	// content holds our static web server content.
	t, err := template.ParseFS(
		fsys,
		// os.DirFS("./templates"),
		"templates/*.tmpl",
	)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	schemaNames := make(map[string]*schemas.Struct)

	structs := make([]*schemas.Struct, 0)
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
			structs = append(structs, &s)
			// fmt.Printf("// %v\n%v ", name, s.String())
		case "array":
		}
	}

	// createTemplate("openapi")
	// fmt.Println("func Route(router *gin.Engine) {")
	endpoints := make([]Endpoint, 0)
	for key, path := range doc.Paths {
		e := printPath(key, path, schemaNames, t)
		endpoints = append(endpoints, e...)
	}

	data := TemplateData{
		Structs:   structs,
		Endpoints: endpoints,
	}

	f, err := os.Create("./gen.go")
	if err != nil {
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	t.ExecuteTemplate(writer, "main.tmpl", data)
	// fmt.Println("}")
	writer.Flush()

}

type TemplateData struct {
	Structs   []*schemas.Struct
	Endpoints []Endpoint
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

func printPath(key string, path *openapi3.PathItem, s map[string]*schemas.Struct, t *template.Template) []Endpoint {
	endpoints := make([]Endpoint, 0)
	// fmt.Printf("path: %v\n", key)
	// printOperation("Connect", path.Connect)
	// printOperation("Delete", path.Delete)
	endpoint := GetEndpoint(key, path.Get, s, t)
	endpoints = append(endpoints, endpoint)
	// printOperation("Head", path.Head)
	// printOperation("Options", path.Options)
	// printOperation("Patch", path.Patch)
	// printOperation("Post", path.Post)
	// printOperation("Put", path.Put)
	// printOperation("Trace", path.Trace)
	return endpoints
}

func GetEndpoint(key string, operation *openapi3.Operation, s map[string]*schemas.Struct, t *template.Template) Endpoint {
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

func GetParams(operation *openapi3.Operation) []ParamProperty {
	parameters := make([]ParamProperty, 0)
	for _, parameterRef := range operation.Parameters {
		parameter := parameterRef.Value
		if parameter.In == openapi3.ParameterInPath {
			parameters = append(parameters, ParamProperty{
				Name: GetPropertyName(parameter.Name),
				Type: GetPropertyType(parameter.Schema.Value.Type),
				Key:  parameter.Name,
			})
		}
	}
	return parameters
}

type ParamProperty struct {
	Name string
	Type string
	Key  string
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

type BodyProperty struct {
	Type     string
	Property string
}

func GetBody(operation *openapi3.Operation) map[string]BodyProperty {
	body := make(map[string]BodyProperty)
	return body
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
