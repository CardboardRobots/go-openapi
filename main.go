package main

import (
	"context"
	"embed"
	"log"

	"github.com/cardboardrobots/go-openapi/core"
)

//go:embed templates/*
var content embed.FS

func main() {

	log.SetPrefix("[OpenApi Loader] ")
	log.Println("Starting...")
	ctx := context.Background()

	core.ParseDocument(ctx, content)
}
