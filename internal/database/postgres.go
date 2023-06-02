package database

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(conn string) (*sql.DB, error) {

	db, err := sql.Open("pgx", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
