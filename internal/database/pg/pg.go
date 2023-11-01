package pg

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// PG - структура бд
type PG struct {
	databaseConf string
	DB           *sql.DB
}

// NewDB конструктор
func NewDB(conf string) *PG {
	return &PG{
		databaseConf: conf,
	}
}

// InitDB - инициализация
func InitDB(conn string) *PG {
	if conn == "" {
		return nil
	}
	db := NewDB(conn)
	if err := db.Open(); err != nil {
		log.Println("Don't connect DataBase")
		log.Fatal(err)
		return nil
	}
	if err := db.PingDB(); err != nil {
		log.Println("Don't ping DataBase")
		log.Fatal(err)
		return nil
	}
	return db
}

// Open - открытие соединения с бд
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

// Close - закрытие соединения
func (pg *PG) Close() {
	if pg.DB != nil {
		pg.DB.Close()
	}
}

// PingDB - пинг бд
func (pg *PG) PingDB() error {
	if err := pg.DB.Ping(); err != nil {
		log.Println("don't ping Database")
		return err
	}
	return nil
}
