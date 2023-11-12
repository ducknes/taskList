package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresConnection(connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("postgres: can`t connect to database. err: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("postgtes: don`t ping connection. err: %v", err)
	}

	return db, nil
}
