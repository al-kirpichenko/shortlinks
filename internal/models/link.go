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
	if _, err := l.Store.DB.Exec("CREATE TABLE IF NOT EXISTS links (id SERIAL PRIMARY KEY , short CHAR (20), original CHAR (255));"); err != nil {
		return err
	}
	return nil
}

func (l *Link) Insert(link *entities.Link) error {
	if err := l.CreateTable(); err != nil {
		return err
	}
	if _, err := l.Store.DB.Exec(
		"INSERT INTO links (short, original) VALUES ($1,$2) RETURNING id",
		link.Short, link.Original); err != nil {
		return err
	}
	return nil
}

func (l *Link) InsertLinks(links []*entities.Link) error {
	return nil
}

func (l *Link) GetOriginal(short string) (*entities.Link, error) {

	link := entities.Link{
		Short: short,
	}
	if err := l.Store.DB.QueryRow("SELECT original FROM links WHERE short = $1", link.Short).Scan(&link.Original); err != nil {
		log.Println(err)
		return nil, err
	}
	return &link, nil
}
