package main

import (
	"context"
	"embed"
	"flag"
	"log"

	"github.com/cardboardrobots/go-openapi/core"
)

//go:embed templates/*
var content embed.FS

func main() {
	options := getOptions()

	log.SetPrefix("[OpenApi Loader] ")
	log.Println("Starting...")

	ctx := context.Background()
	core.ParseDocument(ctx, content, options)
}

func getOptions() core.ParseOptions {
	Input := flag.String("i", "openapi.yml", "Input location")
	Output := flag.String("o", "gen.go", "Output location")
	Package := flag.String("p", "definition", "Package name")

	flag.Parse()

	return core.ParseOptions{
		Input:   *Input,
		Output:  *Output,
		Package: *Package,
	}
}
