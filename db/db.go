package db

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewDatabase(logger *slog.Logger, url string) (*Database, error) {
	db, err := sql.Open("sqlite3", url)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{
		db:     db,
		logger: logger,
	}, nil
}

func (d *Database) DB() *sql.DB {
	return d.db
}

func (d *Database) Close() error {
	return d.db.Close()
}
