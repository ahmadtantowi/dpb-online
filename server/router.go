package server

import (
	"dpb-online/db"
	"dpb-online/server/middleware"
	"log/slog"
	"net/http"
)

func NewRouter(logger *slog.Logger, db *db.Database) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	return middleware.NewLoggingMiddleware(logger, mux)
}

func newPath(method string, path string) string {
	return method + " " + path
}
