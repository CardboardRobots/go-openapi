package parser

import (
	"context"
	"log"

	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

type ParseOptions struct {
	Input   string
	Output  string
	Package string
}

func ParseDocument(ctx context.Context, options ParseOptions) entity.TemplateData {
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

	return entity.TemplateData{
		Package:   options.Package,
		Structs:   schemaParser.GetSchemas(),
		Endpoints: schemaParser.GetEndpoints(),
	}
}
