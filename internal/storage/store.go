package storage

import "github.com/al-kirpichenko/shortlinks/internal/entities"

type Storage interface {
	Insert(link entities.Link) error
	InsertLinks(links []entities.Link) error
	GetOriginal(short string) (entities.Link, error)
}
