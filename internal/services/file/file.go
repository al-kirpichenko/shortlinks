package file

import (
	"bufio"
	"encoding/json"
	"os"

	"go.uber.org/zap"

	"github.com/al-kirpichenko/shortlinks/internal/models"
)

// SaveToFile сохраняет экземпляр *models.Link в файл
func SaveToFile(link *models.Link, fileName string) error {

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(link)
	return err
}

// AllSaveToFile сохраняет слайс экземпляров *models.Link в файл
func AllSaveToFile(links []*models.Link, fileName string) error {

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	for _, v := range links {
		err = encoder.Encode(v)
	}
	return err
}

// LoadFromFile загружает из файла слайс пар сокращенный url[оригинальный url]
func LoadFromFile(fileName string) (map[string]string, error) {

	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := make(map[string]string)

	for scanner.Scan() {

		var d models.Link
		err = json.Unmarshal(scanner.Bytes(), &d)
		if err != nil {
			zap.L().Error("error scan ", zap.Error(err))
		}

		data[d.Short] = d.Original
	}
	return data, nil
}
