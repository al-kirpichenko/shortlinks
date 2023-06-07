package models

import (
	"github.com/al-kirpichenko/shortlinks/internal/database/pg"
	"github.com/al-kirpichenko/shortlinks/internal/entities"
	"log"
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

func (l *Link) Insert(link *entities.Link) (*entities.Link, error) {
	if err := l.CreateTable(); err != nil {
		return nil, err
	}
	if err := l.Store.DB.QueryRow(
		"INSERT INTO links (short, original) VALUES ($1,$2) RETURNING id",
		link.Short, link.Original,
	).Scan(&link.ID); err != nil {
		return nil, err
	}
	return link, nil
}

func (l *Link) GetOriginal(link *entities.Link) (*entities.Link, error) {

	if err := l.Store.DB.QueryRow("SELECT original FROM links WHERE short = $1", link.Short).Scan(&link.Original); err != nil {
		log.Println(err)
		return nil, err
	}
	return link, nil
}