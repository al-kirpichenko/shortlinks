package storage

import (
	"errors"

	"github.com/al-kirpichenko/shortlinks/internal/models"
)

var ErrConflict = errors.New("conflict on inserting new record")

type Storage interface {
	Insert(link *models.Link) error
	InsertLinks(links []*models.Link) error
	GetOriginal(short string) (*models.Link, error)
	GetShort(original string) (*models.Link, error)
	GetAllByUserID(userID string) ([]*models.Link, error)
}
