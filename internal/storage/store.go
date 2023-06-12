package storage

import (
	"github.com/al-kirpichenko/shortlinks/internal/models"
)

type Storage interface {
	Insert(link *models.Link) error
	InsertLinks(links []*models.Link) error
	GetOriginal(short string) (*models.Link, error)
	GetShort(original string) (*models.Link, error)
}
