package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

var (
	version  = "dev"
	revision = "unknown"
)

const message = "Hello from task-4 (21:45:40)"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/version", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]string{
			"service":  "backend",
			"version":  version,
			"revision": revision,
			"message":  message,
		}); err != nil {
			slog.Error("writing the version response failed", "err", err)
		}
	})
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	addr := ":8080"
	slog.Info("listening", "addr", addr, "version", version, "revision", revision)
	if err := http.ListenAndServe(addr, mux); err != nil {
		slog.Error("server stopped", "err", err)
		os.Exit(1)
	}
}
