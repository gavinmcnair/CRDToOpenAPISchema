# CRD to OpenAPI Schema Converter

This tool processes Kubernetes Custom Resource Definitions (CRDs) and converts them into OpenAPI JSON Schema files. The generated schema files can be used with tools like `kubeval` or `kubeconform` for validating CRDs.

## Table of Contents

- [Features](#features)
- [Directory Structure](#directory-structure)
- [Getting Started](#getting-started)
    - [Dependencies](#dependencies)
    - [Building the Project](#building-the-project)
    - [Running the Project](#running-the-project)
- [Usage](#usage)
- [Makefile Commands](#makefile-commands)
- [Docker Usage](#docker-usage)
    - [Build Docker Image](#build-docker-image)
    - [Run Docker Container](#run-docker-container)
- [Cleaning Up](#cleaning-up)
- [License](#license)

## Features

- Converts CRDs to OpenAPI JSON Schema files.
- Supports processing multiple CRDs in a specified directory.
- Skips schema generation if the schema file already exists.
- Outputs schemas into a structured directory `schemas/<group name>/<api version>/<name>.json`.

## Directory Structure

```
crd-extractor/
├── crdconv/
│   ├── convert.go
│   ├── convert_test.go
│   ├── utils.go
├── crds/             # Contains CRD files to be processed
├── schemas/          # Output directory for JSON schema files
├── go.mod
├── go.sum
├── main.go
├── Dockerfile
├── Makefile
└── README.md
```

## Getting Started

### Dependencies

Ensure you have the following installed:

- Go (1.16 or later)
- Docker
- GNU Make

### Building the Project

To build the project, run:

```
make build
```

This will compile the Go application and generate the `crd-extractor` binary.

### Running the Project

To run the project, execute the following command:

```
make docker-run
```

This will process all CRD files in the `crds` directory and generate the corresponding JSON schema files in the `schemas` directory, skipping files that already exist.

## Usage

1. Place your CRD files in the `crds` directory.
2. Ensure the `schemas` directory exists or is created by the tool.
3. Run the tool using:

   ```
   make docker-run
   ```

4. The generated JSON schema files will be available in the `schemas` directory, organized by group, version, and name.

## Makefile Commands

- `make build`: Compiles the Go application and generates the `crd-extractor` binary.
- `make docker-build`: Builds the Docker image.
- `make docker-run`: Runs the Docker container and processes the CRD files.
- `make clean`: Cleans up the generated schema files.
- `make all`: Combines `build` and `docker-build` commands to build the Go binary and Docker image.

## Docker Usage

### Build Docker Image

To build the Docker image, run:

```
make docker-build
```

### Run Docker Container

To run the Docker container, execute:

```
make docker-run
```

This mounts the local `schemas` directory to the container and runs the CRD processing tool inside the container.

## Cleaning Up

To clean up the generated schema files, run:

```
make clean
```

This will remove the files in the `schemas` directory.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
