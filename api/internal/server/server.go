// Package server wires the HTTP routes of the API.
//
// It uses only the Go standard library (net/http with method-aware routing,
// Go 1.22+). Business logic does not live here; handlers will delegate to
// domain packages as they appear.
package server

import (
	"encoding/json"
	"net/http"

	"github.com/digitaldemocracy-app/platform/api/openapi"
)

// Version is the build version. It is overridden at build time with
// -ldflags "-X github.com/digitaldemocracy-app/platform/api/internal/server.Version=<tag>".
var Version = "0.0.1-dev"

// Server holds what the handlers need. Nothing yet beyond identity.
type Server struct {
	name    string
	version string
}

// New returns a Server identified by name and version.
func New(name, version string) *Server {
	return &Server{name: name, version: version}
}

type healthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Version string `json:"version"`
}

type helloResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Docs    string `json:"docs"`
	OpenAPI string `json:"openapi"`
}

// Handler returns the HTTP handler with all routes registered.
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", s.handleHello)
	mux.HandleFunc("GET /healthz", s.handleHealth)
	mux.HandleFunc("GET /openapi.yaml", s.handleOpenAPI)
	return mux
}

func (s *Server) handleHello(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, helloResponse{
		Name:    s.name,
		Message: "Digital Democracy API. Nothing here yet beyond this hello.",
		Docs:    "https://github.com/digitaldemocracy-app/docs",
		OpenAPI: "/openapi.yaml",
	})
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{
		Status:  "ok",
		Service: s.name,
		Version: s.version,
	})
}

func (s *Server) handleOpenAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/yaml; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(openapi.Spec)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
