# Use the official GoLang image as the base image
FROM golang:1.17 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Set the environment variables for Go build
ENV GOARCH=amd64
ENV GOOS=linux

# Copy go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .

# Download and install dependencies
RUN go mod download

# Copy the entire source code to the working directory
COPY . .

# Build the GoLang application
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory to /app in the container
WORKDIR /app

# Copy the compiled binary from the builder stage to the new stage
COPY --from=builder /app/app .

# Copy the config directory into the container
COPY --from=builder /app/config ./config

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./app"]
