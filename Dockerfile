# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

# Install necessary build tools
RUN apk add --no-cache gcc musl-dev

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the entire project source into the container
COPY . .

# Build the Go app, specifying the directory containing the entry point
RUN go build -o ./main ./cmd/api

## Stage 2: A minimal image to run the Go binary
FROM alpine:latest
#
## Set the Current Working Directory inside the container
WORKDIR /root/
#
## Copy the Pre-built binary file from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
