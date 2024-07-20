FROM golang:1.21-alpine AS build

# Update and install git for fetch package
RUN apk update && apk add --no-cache git

# Set the Current Working Directory of the container
WORKDIR /app

# Copy go mod and go sum for install the package
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy all project to container
COPY . .

# Ensure all packages (my personla pkg too) fetched
RUN go mod tidy

# Build the entrypoint file (main.go)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/app/main.go

# Start a new stage
FROM alpine:latest

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/main /app/main

# Copy the .env file
COPY --from=build /app/.env /app/.env

# Expose port 3333 to the outside world
EXPOSE 3333

# Command to run the executable
ENTRYPOINT ["/app/main"]
