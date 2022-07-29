package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"log"

	"github.com/cardboardrobots/go-openapi/core"
)

//go:embed templates/*
var content embed.FS

func main() {

	ctx := context.Background()
	options := GetOptions()

	log.SetPrefix("[OpenApi Loader] ")
	log.Println("Starting...")

	core.ParseDocument(ctx, content, options)
}

func GetOptions() core.ParseOptions {
	options := core.ParseOptions{
		Input:   *flag.String("i", "openapi.yml", "Input location"),
		Output:  *flag.String("o", "gen.go", "Output location"),
		Package: *flag.String("p", "definition", "Package name"),
	}
	flag.Parse()
	fmt.Printf("%v", options)
	return options
}
