package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

func NewPostgresConnection(connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("postgres: can`t connect to database. err: %v", err)
	}

	for i := 0; i < 5; i++ {
		time.Sleep(30 * time.Millisecond)
		if err = db.Ping(); err != nil {
			continue
		}
		break
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("postgtes: don`t ping connection. err: %v", err)
	}

	if err = goose.Up(db.DB, "database/migrations"); err != nil {
		return nil, fmt.Errorf("goose: can't run migrations. err: %v", err)
	}

	return db, nil
}
