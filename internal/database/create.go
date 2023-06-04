package database

func (store *DBStore) CreateTable() error {
	if _, err := store.DB.Exec("CREATE TABLE links (id SERIAL PRIMARY KEY , short CHAR (20), original CHAR (255));"); err != nil {
		return err
	}
	return nil
}
