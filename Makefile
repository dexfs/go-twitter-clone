test:
	@go test ./... -v

build:
	@go build -o bin/api cmd/api/api.go

start: build
	./bin/api
