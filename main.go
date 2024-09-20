package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "path/filepath"

    "github.com/gavinmcnair/CRDToOpenAPISchema/pkg/crdconv"
)

func main() {
    crdsDir := flag.String("crds", "crds", "Directory containing the CRD YAML files")
    outputDir := flag.String("output", "schemas", "Output directory for the JSON schemas")
    flag.Parse()

    if *crdsDir == "" {
        fmt.Println("Missing CRDs directory parameter.")
        return
    }

    // Get the list of CRD files
    files, err := ioutil.ReadDir(*crdsDir)
    if err != nil {
        fmt.Printf("Error reading CRDs directory: %v\n", err)
        return
    }

    for _, file := range files {
        if file.IsDir() {
            continue
        }

        crdFile := filepath.Join(*crdsDir, file.Name())
        fmt.Printf("Processing CRD file: %s\n", crdFile)
        
        err := crdconv.ConvertCRDToJSONSchema(crdFile, *outputDir)
        if err != nil {
            fmt.Printf("Error processing CRD file %s: %v\n", crdFile, err)
        }
    }
}

