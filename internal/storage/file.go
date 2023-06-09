package storage

import (
	"log"

	"github.com/al-kirpichenko/shortlinks/internal/entities"
	"github.com/al-kirpichenko/shortlinks/internal/services/file"
)

type FileStorage struct {
	memStorage *InMemoryStorage
	filePATH   string
}

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{
		memStorage: NewInMemoryStorage(),
		filePATH:   path,
	}
}

func (fs *FileStorage) Load(data map[string]string) {
	fs.memStorage.Load(data)
}

func (fs *FileStorage) Insert(link *entities.Link) error {

	err := fs.memStorage.Insert(link)
	if err != nil {
		return err
	}
	err2 := file.SaveToFile(link, fs.filePATH)
	if err2 != nil {
		return err2
	}
	return nil
}

func (fs *FileStorage) InsertLinks(links []*entities.Link) error {
	return nil
}

func (fs *FileStorage) GetOriginal(short string) (*entities.Link, error) {

	link, err := fs.memStorage.GetOriginal(short)

	if err != nil {
		log.Println("")
		return nil, err
	}
	return link, nil
}
