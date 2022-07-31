package core

import (
	"bufio"
	"bytes"
	"context"
	"go/format"
	"io/fs"
	"log"
	"os"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/cardboardrobots/go-openapi/entity"
)

type ParseOptions struct {
	Input   string
	Output  string
	Package string
}

func ParseDocument(ctx context.Context, fsys fs.FS, options ParseOptions) {
	loader := openapi3.Loader{Context: ctx}

	log.Printf("Loading %v...\n", options.Input)
	doc, err := loader.LoadFromFile(options.Input)
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

	schemaNames := make(map[string]*entity.Schema)

	structs := make([]*entity.Schema, 0)
	for key, schemaRef := range doc.Components.Schemas {
		schema := schemaRef.Value
		switch schema.Type {
		case "string":
		case "number":
		case "integer":
		case "object":
			s := SchemaToObject(key, schema)
			name := "#/components/schemas/" + s.Name
			schemaNames[name] = &s
			structs = append(structs, &s)
			// fmt.Printf("// %v\n%v ", name, s.String())
		case "array":
		}
	}

	// createTemplate("openapi")
	endpoints := make([]Endpoint, 0)
	for key, path := range doc.Paths {
		e := printPath(key, path, schemaNames, t)
		endpoints = append(endpoints, e...)
	}

	data := TemplateData{
		Package:   options.Package,
		Structs:   structs,
		Endpoints: endpoints,
	}

	buffer := bytes.NewBufferString("")
	t.ExecuteTemplate(buffer, "main.tmpl", data)
	bytes, err := format.Source(buffer.Bytes())
	if err != nil {
		return
	}

	f, err := os.Create(options.Output)
	if err != nil {
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	writer.Write(bytes)
	writer.Flush()
}

func printPath(key string, path *openapi3.PathItem, s map[string]*entity.Schema, t *template.Template) []Endpoint {
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
