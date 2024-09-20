package crdconv

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"

    "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/apimachinery/pkg/runtime/serializer"
)

// ConvertCRDToJSONSchema converts a CRD file to JSON Schema files in the specified directory structure.
func ConvertCRDToJSONSchema(crdFile string, outputDir string) error {
    yamlBytes, err := ioutil.ReadFile(crdFile)
    if err != nil {
        return fmt.Errorf("error reading CRD file: %v", err)
    }

    scheme := runtime.NewScheme()
    v1.AddToScheme(scheme)
    codecs := serializer.NewCodecFactory(scheme)
    decode := codecs.UniversalDeserializer().Decode

    obj, _, err := decode(yamlBytes, nil, nil)
    if err != nil {
        return fmt.Errorf("error decoding CRD YAML: %v", err)
    }

    crd, ok := obj.(*v1.CustomResourceDefinition)
    if !ok {
        return fmt.Errorf("error: not a CustomResourceDefinition")
    }

    for _, version := range crd.Spec.Versions {
        if version.Schema != nil && version.Schema.OpenAPIV3Schema != nil {
            group := crd.Spec.Group
            name := strings.ToLower(crd.Spec.Names.Kind)
            versionName := version.Name

            // Create the directory path
            dirPath := filepath.Join(outputDir, group, versionName)
            if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
                return fmt.Errorf("error creating directory path: %v", err)
            }

            // Schema file path
            filename := filepath.Join(dirPath, fmt.Sprintf("%s.json", name))

            // Check if the schema file already exists
            if _, err := os.Stat(filename); err == nil {
                // File exists, skip processing
                fmt.Printf("Schema already exists, skipping: %s\n", filename)
                continue
            }

            EnsureAdditionalProperties(&version.Schema.OpenAPIV3Schema.Properties)

            jsonSchema, err := json.MarshalIndent(version.Schema.OpenAPIV3Schema, "", "  ")
            if err != nil {
                return fmt.Errorf("error creating JSON Schema: %v", err)
            }

            // Write the JSON Schema to the file
            if err := ioutil.WriteFile(filename, jsonSchema, 0644); err != nil {
                return fmt.Errorf("error writing JSON Schema file: %v", err)
            }

            fmt.Printf("JSON schema written to %s\n", filename)
        } else {
            fmt.Printf("No OpenAPIv3Schema found in version: %s\n", version.Name)
        }
    }

    return nil
}

