package core

import (
	"bufio"
	"context"
	"io/fs"
	"log"
	"os"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"

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
