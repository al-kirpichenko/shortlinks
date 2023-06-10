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

func (fs *FileStorage) Insert(link *entities.Link) (*entities.Link, error) {

	_, err := fs.memStorage.Insert(link)
	if err != nil {
		return link, err
	}
	err2 := file.SaveToFile(link, fs.filePATH)
	if err2 != nil {
		return link, err2
	}
	return link, nil
}

func (fs *FileStorage) InsertLinks(links []*entities.Link) error {

	err := fs.memStorage.InsertLinks(links)
	if err != nil {
		return err
	}
	err2 := file.AllSaveToFile(links, fs.filePATH)
	if err2 != nil {
		return err2
	}
	return err
}

func (fs *FileStorage) GetOriginal(short string) (*entities.Link, error) {

	link, err := fs.memStorage.GetOriginal(short)

	if err != nil {
		log.Println("")
		return link, err
	}
	return link, nil
}

func (fs *FileStorage) GetShort(original string) (*entities.Link, error) {

	link, err := fs.memStorage.GetShort(original)

	if err != nil {
		log.Println("")
		return link, err
	}
	return link, nil
}
