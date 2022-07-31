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

	t, err := template.ParseFS(fsys, "templates/*.tmpl")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	schemaParser := NewSchemaParser()
	schemaParser.Parse(doc)

	data := TemplateData{
		Package:   options.Package,
		Structs:   schemaParser.GetSchemas(),
		Endpoints: schemaParser.GetEndpoints(),
	}

	log.Println("generating output...")
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

	f, err := os.Create(options.Output)
	if err != nil {
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	writer.Write(bytes)
	writer.Flush()
}
