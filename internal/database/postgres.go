package database

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBStore struct {
	DatabaseConf string
	SqlDB        *sql.DB
}

func NewDB(conf string) *DBStore {
	return &DBStore{
		DatabaseConf: conf,
	}
}

func (store *DBStore) Open() error {

	db, err := sql.Open("pgx", store.DatabaseConf)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	store.SqlDB = db
	return nil
}

func (store *DBStore) Close() {
	store.SqlDB.Close()
}
