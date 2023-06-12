package storage

import (
	"errors"
	"log"

	"github.com/jackc/pgerrcode"

	"github.com/al-kirpichenko/shortlinks/internal/database/pg"
	"github.com/al-kirpichenko/shortlinks/internal/models"
	"github.com/al-kirpichenko/shortlinks/internal/services/sqlerror"
)

var ErrConflict = errors.New("conflict on inserting new record")

type Link struct {
	Store *pg.PG
}

func (l *Link) CreateTable() error {
	if _, err := l.Store.DB.Exec("CREATE TABLE IF NOT EXISTS links (id SERIAL PRIMARY KEY , short CHAR (20) UNIQUE, original CHAR (255) UNIQUE );"); err != nil {
		return err
	}
	return nil
}

func (l *Link) Insert(link *models.Link) error {
	if err := l.CreateTable(); err != nil {
		return err
	}
	if _, err := l.Store.DB.Exec(
		"INSERT INTO links (short, original) VALUES ($1,$2)",
		link.Short, link.Original); err != nil {
		if sqlerror.GetSQLState(err) == pgerrcode.UniqueViolation {
			err = ErrConflict
		}
		return err
	}
	return nil
}

func (l *Link) InsertLinks(links []*models.Link) error {
	if err := l.CreateTable(); err != nil {
		return err
	}
	tx, err := l.Store.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(
		"INSERT INTO links (short, original) VALUES($1,$2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, v := range links {
		_, err := stmt.Exec(v.Short, v.Original)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (l *Link) GetOriginal(short string) (*models.Link, error) {

	link := &models.Link{
		Short: short,
	}
	if err := l.Store.DB.QueryRow("SELECT original FROM links WHERE short = $1", link.Short).Scan(&link.Original); err != nil {
		log.Println(err)
		return link, err
	}
	return link, nil
}

func (l *Link) GetShort(original string) (*models.Link, error) {

	link := models.Link{
		Original: original,
	}
	if err := l.Store.DB.QueryRow("SELECT short FROM links WHERE original = $1", link.Original).Scan(&link.Short); err != nil {
		log.Println(err)
		return &link, err
	}
	return &link, nil
}
