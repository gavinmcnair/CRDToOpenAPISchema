package crdconv

import (
    "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// EnsureAdditionalProperties ensures additionalProperties is set to false,
// and adds types where necessary.
func EnsureAdditionalProperties(properties *map[string]v1.JSONSchemaProps) {
    for key, prop := range *properties {
        if prop.Type == "object" {
            if prop.Properties != nil && len(prop.Properties) > 0 {
                falseVal := false
                prop.AdditionalProperties = &v1.JSONSchemaPropsOrBool{
                    Allows: falseVal,
                }
                EnsureAdditionalProperties(&prop.Properties)
            }
            if prop.Description != "" {
                prop.Type = "object"
            }
        }

        // Add missing properties
        if prop.Properties != nil && len(prop.Properties) > 0 {
            prop.Type = "object"
        }

        // Ensure consistent order of keys
        ensureConsistentFieldOrder(&prop)

        (*properties)[key] = prop
    }
}

// ensureConsistentFieldOrder reorders properties to ensure consistency.
func ensureConsistentFieldOrder(prop *v1.JSONSchemaProps) {
    newProp := v1.JSONSchemaProps{
        Description: prop.Description,
        Type:        prop.Type,
        Properties:  prop.Properties,
        Items:       prop.Items,
        Required:    prop.Required,
        Enum:        prop.Enum,
    }
    *prop = newProp
}

