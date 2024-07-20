# Use the official Golang image as a build stage
FROM golang:1.21-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and go sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Set GO111MODULE environment variable to enable Go modules
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

# Build the Go app, specifying the entry point
RUN go build -o main ./cmd/app/main.go

# Start a new stage from scratch
FROM alpine:latest

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/main /app/main
COPY --from=build /app/.env /app/.env

# Expose port 3333 to the outside world
EXPOSE 3333

# Command to run the executable
ENTRYPOINT ["/app/main"]
