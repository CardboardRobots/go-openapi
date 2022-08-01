package main

import (
	"context"
	"embed"
	"flag"
	"log"

	"github.com/cardboardrobots/go-openapi/parser"
	"github.com/cardboardrobots/go-openapi/writer"
)

//go:embed templates/*
var content embed.FS

func main() {
	options := getOptions()

	log.SetPrefix("[OpenApi Loader] ")
	log.Println("Starting...")

	ctx := context.Background()
	data := parser.ParseDocument(ctx, options)
	writer.WriteTemplate(content, options.Output, data)
}

func getOptions() parser.ParseOptions {
	Input := flag.String("i", "openapi.yml", "Input location")
	Output := flag.String("o", "gen.go", "Output location")
	Package := flag.String("p", "definition", "Package name")

	flag.Parse()

	return parser.ParseOptions{
		Input:   *Input,
		Output:  *Output,
		Package: *Package,
	}
}
