package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func (db *DB) createTables() error {
	schemas := db.getSchemas()

	for _, schemaStatement := range schemas {
		_, err := db.db.Exec(schemaStatement)
		if err != nil {
			return err
		}
	}

	return nil
}

func New(dsn string) (*DB, error) {
	db, err := sql.Open("sqlite3", "./"+dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	conn := DB{db: db}

	if err := conn.createTables(); err != nil {
		return nil, err
	}

	return &conn, nil
}
