package models

import (
	"log"

	"github.com/al-kirpichenko/shortlinks/internal/database/pg"
	"github.com/al-kirpichenko/shortlinks/internal/entities"
)

type Link struct {
	Store *pg.PG
}

func (l *Link) CreateTable() error {
	if _, err := l.Store.DB.Exec("CREATE TABLE IF NOT EXISTS links (id SERIAL PRIMARY KEY , short CHAR (20) UNIQUE, original CHAR (255) UNIQUE );"); err != nil {
		return err
	}
	return nil
}

func (l *Link) Insert(link *entities.Link) (*entities.Link, error) {
	if err := l.CreateTable(); err != nil {
		return link, err
	}
	if _, err := l.Store.DB.Exec(
		"INSERT INTO links (short, original) VALUES ($1,$2)",
		link.Short, link.Original); err != nil {

		//if sqlerror.GetSQLState(err) == "23505" {
		//	link2, err := l.GetShort(link.Original)
		//	if err != nil {
		//		log.Println("Don't read data from table")
		//		log.Println(err)
		//		return link2, err
		//	}
		//	return link2, nil
		//}
		return link, err
	}
	return link, nil
}

func (l *Link) InsertLinks(links []*entities.Link) error {
	if err := l.CreateTable(); err != nil {
		return err
	}
	tx, err := l.Store.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(
		"INSERT INTO links (short, original) VALUES($1,$2) ON CONFLICT (original, short) DO NOTHING ")
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

func (l *Link) GetOriginal(short string) (*entities.Link, error) {

	link := &entities.Link{
		Short: short,
	}
	if err := l.Store.DB.QueryRow("SELECT original FROM links WHERE short = $1", link.Short).Scan(&link.Original); err != nil {
		log.Println(err)
		return link, err
	}
	return link, nil
}

func (l *Link) GetShort(original string) (*entities.Link, error) {

	link := entities.Link{
		Original: original,
	}
	if err := l.Store.DB.QueryRow("SELECT short FROM links WHERE original = $1", link.Original).Scan(&link.Short); err != nil {
		log.Println(err)
		return &link, err
	}
	return &link, nil
}
