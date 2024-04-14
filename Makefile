test:
	@go test ./... -v

start:
	@go run cmd/api/api.go