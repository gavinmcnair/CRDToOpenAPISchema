# Start with the official Golang image as a build stage
FROM golang:1.23.1 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o crd-extractor main.go

# Start a new stage from scratch
FROM alpine:latest

# Add Maintainer info
LABEL maintainer="Your Name <your-email@example.com>"

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/crd-extractor .

# Ensure the schemas directory is present
RUN mkdir -p /root/schemas

# Define the command to run the executable
CMD ["./crd-extractor", "-crds", "/root/crds", "-output", "/root/schemas"]

