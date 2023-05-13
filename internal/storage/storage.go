package storage

import "errors"

type Storage struct {
	urls map[string]string
}

func NewStorage() *Storage {
	return &Storage{
		urls: make(map[string]string),
	}
}

func (s *Storage) SetURL(key, value string) {
	s.urls[key] = value
}

func (s *Storage) GetURL(key string) (string, error) {

	url, ok := s.urls[key]
	if ok {
		return url, nil
	}
	return "", errors.New("id not found")
}
