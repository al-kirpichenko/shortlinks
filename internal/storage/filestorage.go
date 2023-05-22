package storage

import (
	"bufio"
	"encoding/json"
	"os"
)

type FileStorage struct {
	Short    string `json:"short_url"`
	Original string `json:"original_url"`
}

func NewFileStorage() *FileStorage {
	return &FileStorage{}
}

func (fs *FileStorage) AddURL(short string, original string) {
	fs.Short = short
	fs.Original = original
}

func (fs *FileStorage) SaveToFile(fileName string) error {

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(fs)
	if err != nil {
		panic(err)
	}
	return err
}

func LoadInFile(fileName string) map[string]string {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return make(map[string]string)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := make(map[string]string)

	for scanner.Scan() {
		var d FileStorage
		// Декодируем строку из формата json
		err = json.Unmarshal(scanner.Bytes(), &d)
		if err != nil {
			panic(err)
		}

		data[d.Short] = d.Original
	}
	return data
}
