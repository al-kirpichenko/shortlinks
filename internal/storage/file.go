package storage

import (
	"go.uber.org/zap"

	"github.com/al-kirpichenko/shortlinks/internal/models"
	"github.com/al-kirpichenko/shortlinks/internal/services/file"
)

// FileStorage - файловое хранилище
type FileStorage struct {
	memStorage *InMemoryStorage
	filePATH   string
}

// NewFileStorage - конструктор
func NewFileStorage(path string) *FileStorage {
	fs := FileStorage{
		memStorage: NewInMemoryStorage(),
		filePATH:   path,
	}
	fs.ConfigureFileStorage()

	return &fs
}

// ConfigureFileStorage - конфигуратор
func (fs *FileStorage) ConfigureFileStorage() {

	data, err := file.LoadFromFile(fs.filePATH)

	if err != nil {
		zap.L().Fatal("Don't load from file!", zap.Error(err))
	}

	fs.Load(data)
}

// Load - загрузка данных из mem storage
func (fs *FileStorage) Load(data map[string]string) {
	fs.memStorage.Load(data)
}

// Insert - вставка записи
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

// InsertLinks - массовая вставка
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

// GetOriginal - получение оригинального url по короткому
func (fs *FileStorage) GetOriginal(short string) (*models.Link, error) {

	link, err := fs.memStorage.GetOriginal(short)

	if err != nil {
		zap.L().Error("Don't get original URL", zap.Error(err))
		return link, err
	}
	return link, nil
}

// GetShort получение короткого по оригинальному
func (fs *FileStorage) GetShort(original string) (*models.Link, error) {

	link, err := fs.memStorage.GetShort(original)

	if err != nil {
		zap.L().Error("Don't get short URL", zap.Error(err))
		return link, err
	}
	return link, nil
}

// GetAllByUserID - получение всех записей пользователя
func (fs *FileStorage) GetAllByUserID(userID string) ([]models.Link, error) {
	return nil, nil
}

// DelURL удаление записей
func (fs *FileStorage) DelURL(shortURLs []string) error {
	return nil
}

// GetCountURLs - возвращает число записей в таблице
func (fs *FileStorage) GetCountURLs() (int, error) {
	return 0, nil
}

// GetCountUsers - возвращает число записей в таблице
func (fs *FileStorage) GetCountUsers() (int, error) {
	return 0, nil
}

// Ping -
func (fs *FileStorage) Ping() error {
	return nil
}
