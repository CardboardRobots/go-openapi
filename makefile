.PHONY: build
build:
	go run .
	go fmt gen.go
