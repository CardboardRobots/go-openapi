package main

import (
	"context"
	"log"

	"github.com/cardboardrobots/go-openapi/core"
)

func main() {
	log.SetPrefix("[OpenApi Loader] ")
	log.Println("Starting...")
	ctx := context.Background()

	core.ParseDocument(ctx)
}
