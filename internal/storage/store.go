package storage

import (
	"errors"

	"github.com/al-kirpichenko/shortlinks/internal/models"
)

// ErrConflict - для ошибок вставки
var ErrConflict = errors.New("conflict on inserting new record")

// Storage интерфейс хранилища
type Storage interface {
	Insert(link *models.Link) error
	InsertLinks(links []*models.Link) error
	GetOriginal(short string) (*models.Link, error)
	GetShort(original string) (*models.Link, error)
	GetAllByUserID(userID string) ([]models.Link, error)
	DelURL(shortURLs []string) error
	GetCountURLs() (int, error)
	GetCountUsers() (int, error)
	Ping() error
}
