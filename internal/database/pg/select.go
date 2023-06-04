package pg

func (store *DBStore) Select(short string) (string, error) {
	row := store.DB.QueryRow("SELECT original FROM links WHERE short = $1", short)

	var res string

	if err := row.Scan(&res); err != nil {
		return "", err
	}
	return res, nil
}
