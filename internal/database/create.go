package database

import "log"

func (store *DBStore) CreateTable() error {
	if _, err := store.DB.Exec("CREATE TABLE IF NOT EXISTS links (id SERIAL PRIMARY KEY , short CHAR (20), original CHAR (255));"); err != nil {
		return err
	}
	log.Println("the table has been created!")
	return nil
}
