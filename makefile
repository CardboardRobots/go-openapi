.PHONY: build
build:
	go run ./cmd > gen.go
	go fmt gen.go
