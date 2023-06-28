package storage

import (
	"go.uber.org/zap"

	"github.com/al-kirpichenko/shortlinks/internal/models"
	"github.com/al-kirpichenko/shortlinks/internal/services/file"
)

type FileStorage struct {
	memStorage *InMemoryStorage
	filePATH   string
}

func NewFileStorage(path string) *FileStorage {
	fs := FileStorage{
		memStorage: NewInMemoryStorage(),
		filePATH:   path,
	}
	fs.ConfigureFileStorage()

	return &fs
}

func (fs *FileStorage) ConfigureFileStorage() {

	data, err := file.LoadFromFile(fs.filePATH)

	if err != nil {
		zap.L().Fatal("Don't load from file!", zap.Error(err))
	}

	fs.Load(data)
}

func (fs *FileStorage) Load(data map[string]string) {
	fs.memStorage.Load(data)
}

func (fs *FileStorage) Insert(link *models.Link) error {

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

func (fs *FileStorage) InsertLinks(links []*models.Link) error {

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

func (fs *FileStorage) GetOriginal(short string) (*models.Link, error) {

	link, err := fs.memStorage.GetOriginal(short)

	if err != nil {
		zap.L().Error("Don't get original URL", zap.Error(err))
		return link, err
	}
	return link, nil
}

func (fs *FileStorage) GetShort(original string) (*models.Link, error) {

	link, err := fs.memStorage.GetShort(original)

	if err != nil {
		zap.L().Error("Don't get short URL", zap.Error(err))
		return link, err
	}
	return link, nil
}
