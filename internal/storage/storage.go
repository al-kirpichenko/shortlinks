package storage

import "errors"

type Storage struct {
	Urls map[string]string
}

func (s *Storage) SetURL(key, value string) {
	s.Urls[key] = value
}

func (s *Storage) GetURL(key string) (string, error) {

	url, ok := s.Urls[key]
	if ok {
		return url, nil
	}
	return "", errors.New("id not found")
}
