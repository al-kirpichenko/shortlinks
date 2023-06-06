package pg

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

type PG struct {
	databaseConf string
	DB           *sql.DB
}

func NewDB(conf string) *PG {
	return &PG{
		databaseConf: conf,
	}
}

func (pg *PG) Open() error {

	db, err := sql.Open("pgx", pg.databaseConf)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	pg.DB = db

	return nil
}

func (pg *PG) Close() {
	pg.DB.Close()
}

func (pg *PG) PingDB() error {
	if err := pg.DB.Ping(); err != nil {
		log.Println("don't ping Database")
		return err
	}
	return nil
}
