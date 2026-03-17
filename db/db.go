package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Open(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("db: open: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("db: ping: %w", err)
	}

	return db, nil
}
