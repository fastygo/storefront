SHELL := /bin/bash

APP := bin/storefront

.PHONY: dev build css generate test lint

dev:
	templ generate ./...
	bun run build:css
	go run ./cmd/server

build:
	templ generate ./...
	bun run build:css
	go build -o $(APP) ./cmd/server

css:
	bun run build:css

generate:
	templ generate ./...

test:
	go test ./...

lint:
	go test ./...
