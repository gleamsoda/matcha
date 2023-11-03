package db

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Open(connString string) (*sql.DB, error) {
	return sql.Open("pgx", connString)
}
