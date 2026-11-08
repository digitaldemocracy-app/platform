// Package openapi embeds the API contract so the running server can publish it.
package openapi

import _ "embed"

// Spec is the OpenAPI 3.1 document, the source of truth for every endpoint.
//
//go:embed openapi.yaml
var Spec []byte
