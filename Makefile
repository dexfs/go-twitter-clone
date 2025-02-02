test:
	@go test ./... -v

build:
	@go build -o bin/api cmd/api/main.go

start: build
	./bin/api

seed:
	@go run cmd/seeds/main.go