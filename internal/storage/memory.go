package storage

import (
	"errors"

	"github.com/al-kirpichenko/shortlinks/internal/models"
)

// InMemoryStorage хранилище
type InMemoryStorage struct {
	urls map[string]string
}

// NewInMemoryStorage - конструктор
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		urls: make(map[string]string),
	}
}

// Load загрузка данных
func (s *InMemoryStorage) Load(data map[string]string) {
	s.urls = data
}

// Insert - вставка записи
func (s *InMemoryStorage) Insert(link *models.Link) error {
	s.urls[link.Short] = link.Original
	return nil
}

// InsertLinks - массовая вставка
func (s *InMemoryStorage) InsertLinks(links []*models.Link) error {

	for _, v := range links {
		s.urls[v.Short] = v.Original
	}
	return nil
}

// GetOriginal - получение оригинального url по короткому
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

// GetShort получение короткого по оригинальному
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

// GetAllByUserID - получение всех записей пользователя
func (s *InMemoryStorage) GetAllByUserID(userID string) ([]models.Link, error) {
	return nil, nil
}

// DelURL удаление записей
func (s *InMemoryStorage) DelURL(shortURLs []string) error {
	return nil
}

// GetCountURLs - возвращает число записей в таблице
func (s *InMemoryStorage) GetCountURLs() (int, error) {
	return 0, nil
}

// GetCountUsers - возвращает число записей в таблице
func (s *InMemoryStorage) GetCountUsers() (int, error) {
	return 0, nil
}

// Ping -
func (s *InMemoryStorage) Ping() error {
	return nil
}
