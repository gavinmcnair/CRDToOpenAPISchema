# Variables
APP_NAME = crd-extractor
DOCKER_IMAGE = $(APP_NAME)
CRDS_DIR = crds
OUTPUT_DIR = schemas

# Build the Go binary
build:
	go build -o $(APP_NAME) main.go

# Build the Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Run the Docker container
docker-run:
	docker run --rm \
		-v $(PWD)/$(CRDS_DIR):/root/crds \
		-v $(PWD)/$(OUTPUT_DIR):/root/schemas \
		$(DOCKER_IMAGE)

# Clean up generated schema files
clean:
	rm -rf $(OUTPUT_DIR)

# Combined build and Docker build
all: build docker-build

.PHONY: build docker-build docker-run clean all

