package server

import (
	"dpb-online/db"
	"dpb-online/server/features/check_nik"
	"dpb-online/server/middleware"
	"log/slog"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(logger *slog.Logger, db *db.Database) http.Handler {
	mux := http.NewServeMux()

	// Serve generated swagger UI and docs from the generated docs package.
	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)

	// Serve a simple index.html at GET /
	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "server/static/index.html")
	}))

	mux.Handle("GET /{nik}", check_nik.CheckNIKHandler(logger, db))
	mux.Handle("GET /check", check_nik.CheckNIKQueryHandler(logger, db))

	return middleware.NewLoggingMiddleware(logger, mux)
}

func newPath(method string, path string) string {
	return method + " " + path
}
