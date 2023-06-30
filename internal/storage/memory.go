package storage

import (
	"errors"

	"github.com/al-kirpichenko/shortlinks/internal/models"
)

type InMemoryStorage struct {
	urls map[string]string
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		urls: make(map[string]string),
	}
}

func (s *InMemoryStorage) Load(data map[string]string) {
	s.urls = data
}

func (s *InMemoryStorage) Insert(link *models.Link) error {
	s.urls[link.Short] = link.Original
	return nil
}

func (s *InMemoryStorage) InsertLinks(links []*models.Link) error {

	for _, v := range links {
		s.urls[v.Short] = v.Original
	}
	return nil
}

func (s *InMemoryStorage) GetOriginal(short string) (*models.Link, error) {

	var ok bool
	link := &models.Link{
		Short: short,
	}
	link.Original, ok = s.urls[link.Short]
	if ok {
		return link, nil
	}
	return link, errors.New("id not found")
}

func (s *InMemoryStorage) GetShort(original string) (*models.Link, error) {

	link := &models.Link{
		Original: original,
	}

	for k, v := range s.urls {
		if v == original {
			link.Short = k
			return link, nil
		}
	}
	return link, errors.New("id not found")
}
func (s *InMemoryStorage) GetAllByUserID(userID string) ([]models.Link, error) {
	return nil, nil
}

func (s *InMemoryStorage) DelURL(shortURLs []string, userid string) error {
	return nil
}
