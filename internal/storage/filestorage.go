package storage

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type FileStorage struct {
	Short    string `json:"short_url"`
	Original string `json:"original_url"`
}

func NewFileStorage() *FileStorage {
	return &FileStorage{}
}

func SaveToFile(fs *FileStorage, fileName string) error {

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(fs)
	return err
}

func LoadFromFile(fileName string) (map[string]string, error) {

	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := make(map[string]string)

	for scanner.Scan() {
		var d FileStorage
		// Декодируем строку из формата json
		err = json.Unmarshal(scanner.Bytes(), &d)
		if err != nil {
			log.Println(err)
		}

		data[d.Short] = d.Original
	}
	return data, nil
}
