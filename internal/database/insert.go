package database

func (store *DBStore) Insert(short, original string) error {
	if _, err := store.DB.Exec("INSERT INTO links VALUES ($1,$2)", short, original); err != nil {
		return err
	}
	return nil
}
