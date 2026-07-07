package pkg

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

func InitPostgres() (*Postgres, error) {
	connStr := "postgresql://neondb_owner:npg_TC2NyUYZvrB8@ep-orange-unit-atici341-pooler.c-9.us-east-1.aws.neon.tech/neondb?sslmode=require"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// проверка подключения
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{DB: db}, nil
}