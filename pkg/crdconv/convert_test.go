package crdconv

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "reflect"
    "testing"

    "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/apimachinery/pkg/runtime/serializer"
)

// Helper function to load a CRD YAML file for testing
func loadCRDYAML(t *testing.T, filename string) *v1.CustomResourceDefinition {
    t.Helper()

    crdFile := filepath.Join("testdata", filename)
    yamlBytes, err := ioutil.ReadFile(crdFile)
    if err != nil {
        t.Fatalf("Error reading CRD file: %v", err)
    }

    scheme := runtime.NewScheme()
    v1.AddToScheme(scheme)
    codecs := serializer.NewCodecFactory(scheme)
    decode := codecs.UniversalDeserializer().Decode

    obj, _, err := decode(yamlBytes, nil, nil)
    if err != nil {
        t.Fatalf("Error decoding CRD YAML: %v", err)
    }

    crd, ok := obj.(*v1.CustomResourceDefinition)
    if !ok {
        t.Fatalf("Error: not a CustomResourceDefinition")
    }

    return crd
}

func TestEnsureAdditionalProperties(t *testing.T) {
    crd := loadCRDYAML(t, "examplecrd.yaml")

    for _, version := range crd.Spec.Versions {
        if version.Schema != nil && version.Schema.OpenAPIV3Schema != nil {
            EnsureAdditionalProperties(&version.Schema.OpenAPIV3Schema.Properties)
            for key, prop := range version.Schema.OpenAPIV3Schema.Properties {
                if prop.Type == "object" && prop.Properties != nil && len(prop.Properties) > 0 {
                    if prop.AdditionalProperties == nil || prop.AdditionalProperties.Allows {
                        t.Errorf("Failed to set additionalProperties to false for %s", key)
                    }
                }
            }
        }
    }
}

func TestConvertCRDToJSONSchema(t *testing.T) {
    outputDir := "testdata/schemas"
    crdFile := "testdata/examplecrd.yaml"

    defer os.RemoveAll(outputDir) // Clean up after test

    err := ConvertCRDToJSONSchema(crdFile, outputDir)
    if err != nil {
        t.Fatalf("Conversion failed: %v", err)
    }

    expectedOutputFile := filepath.Join(outputDir, "sriovnetwork.openshift.io", "v1", "sriovoperatorconfig.json")
    outputFile := expectedOutputFile

    expectedOutput, err := ioutil.ReadFile(expectedOutputFile)
    if err != nil {
        t.Fatalf("Error reading expected output file: %v", err)
    }

    actualOutput, err := ioutil.ReadFile(outputFile)
    if err != nil {
        t.Fatalf("Error reading actual output file: %v", err)
    }

    if !reflect.DeepEqual(expectedOutput, actualOutput) {
        t.Errorf("Output mismatch. Expected:\n%s\nGot:\n%s", string(expectedOutput), string(actualOutput))
    }
}

