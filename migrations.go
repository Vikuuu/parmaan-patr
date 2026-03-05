package main

import (
	"database/sql"
	"embed"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

//go:embed internal/migrations/*.sql
var embedMigrations embed.FS

func migrate(db *sql.DB) {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}
	if err := goose.Up(db, "internal/migrations"); err != nil {
		panic(err)
	}
	slog.Info("Up Migration Done")
}

func setUpDatabase(dbPath string) *sql.DB {
	dbPath = filepath.Join(dbPath, "db")
	err := os.MkdirAll(dbPath, 0o755)
	if err != nil {
		panic(err)
	}

	db := newDatabase(dbPath)
	migrate(db)

	return db
}

func newDatabase(dbPath string) *sql.DB {
	if dbPath == "" {
		dbPath = "./db/my.db"
	} else {
		dbPath = filepath.Join(dbPath, "sqlite.db")
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	var sqliteVersion string
	err = db.QueryRow("SELECT sqlite_version()").Scan(&sqliteVersion)
	if err != nil {
		panic(err)
	}

	slog.Info("DB Connection Success", "version", sqliteVersion)
	return db
}
