// @title DPB Online API
// @version 0.1.0
// @description API for checking NIK in the voter database
package main

import (
	"dpb-online/db"
	"dpb-online/log"
	"dpb-online/server"
	"log/slog"
	"os"

	_ "dpb-online/docs" // Import generated swagger docs

	_ "github.com/mattn/go-sqlite3"
)

var logger *slog.Logger

func main() {
	logger = log.NewLogger(log.GetLevel(), log.GetOutput())

	db, err := db.NewDatabase(logger, getDBPath())
	if err != nil {
		logger.Error("failed to open/create database", "error", err)
		os.Exit(1)
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("failed to close database", "error", err)
		}
		logger.Info("database closed")
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// populate the generated swag docs info at runtime so the UI shows the correct host
	// docs.SwaggerInfo.Host = "localhost:" + port
	// docs.SwaggerInfo.BasePath = "/"

	routerHandler := server.NewRouter(logger, db)
	server := server.NewServer(logger, ":"+port, server.WithRouter(routerHandler))
	server.StartAndWait()
}

func getDBPath() string {
	dbPath := os.Getenv("DATABASE_URL")
	if dbPath == "" {
		dbPath = "database.sqlite"
	}
	slog.Info("using database", "path", dbPath)
	return dbPath
}
