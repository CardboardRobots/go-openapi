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

	schemaParser := NewSchemaParser()
	schemaParser.Parse(doc)

	log.Println("generating output...")
	WriteTemplate(fsys, options.Output, TemplateData{
		Package:   options.Package,
		Structs:   schemaParser.GetSchemas(),
		Endpoints: schemaParser.GetEndpoints(),
	})
}

func WriteTemplate(fsys fs.FS, output string, data TemplateData) {
	t, err := template.ParseFS(fsys, "templates/*.tmpl")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	buffer := bytes.NewBufferString("")
	err = t.ExecuteTemplate(buffer, "main.tmpl", data)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	bytes, err := format.Source(buffer.Bytes())
	if err != nil {
		log.Print(buffer)
		log.Fatalf("%v", err)
		return
	}

	f, err := os.Create(output)
	if err != nil {
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	writer.Write(bytes)
	writer.Flush()
}
