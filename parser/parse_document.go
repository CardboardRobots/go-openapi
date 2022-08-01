package parser

import (
	"context"

	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

type ParseOptions struct {
	Input   string
	Output  string
	Package string
}

func ParseDocument(ctx context.Context, options ParseOptions) (*entity.TemplateData, error) {
	loader := openapi3.Loader{Context: ctx}
	doc, err := loader.LoadFromFile(options.Input)
	if err != nil {
		return nil, err
	}

	err = doc.Validate(ctx)
	if err != nil {
		return nil, err
	}

	schemaParser := NewSchemaParser()
	schemaParser.Parse(doc)

	return &entity.TemplateData{
		Package:   options.Package,
		Structs:   schemaParser.GetSchemas(),
		Endpoints: schemaParser.GetEndpoints(),
	}, nil
}
