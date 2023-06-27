package storage

import (
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"

	"github.com/al-kirpichenko/shortlinks/internal/database/pg"
	"github.com/al-kirpichenko/shortlinks/internal/models"
)

type Link struct {
	Store *pg.PG
}

func NewLinkStorage(db *pg.PG) *Link {
	return &Link{
		Store: db,
	}
}

func (l *Link) CreateTable() error {

	if _, err := l.Store.DB.Exec("CREATE TABLE IF NOT EXISTS links (id SERIAL PRIMARY KEY , userid CHAR (255) NULL, short CHAR (20) UNIQUE, original CHAR (255) UNIQUE, deleted BOOLEAN DEFAULT FALSE );"); err != nil {
		return err
	}
	return nil
}

func (l *Link) Insert(link *models.Link) error {
	if err := l.CreateTable(); err != nil {
		return err
	}
	if _, err := l.Store.DB.Exec(
		"INSERT INTO links (short, original, userid) VALUES ($1,$2,$3)",
		link.Short, link.Original, link.UserID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				err = ErrConflict
			}
		}
		return err
	}
	return nil
}

func (l *Link) InsertLinks(links []*models.Link) error {
	if err := l.CreateTable(); err != nil {
		return err
	}
	tx, err := l.Store.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(
		"INSERT INTO links (short, original, userid) VALUES($1,$2,$3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, v := range links {
		_, err := stmt.Exec(v.Short, v.Original, v.UserID)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (l *Link) GetOriginal(short string) (*models.Link, error) {

	link := &models.Link{
		Short: short,
	}
	if err := l.Store.DB.QueryRow("SELECT original, deleted FROM links WHERE short = $1", link.Short).Scan(&link.Original, &link.Deleted); err != nil {
		zap.L().Error("Don't get original URL", zap.Error(err))
		return link, err
	}
	return link, nil
}

func (l *Link) GetShort(original string) (*models.Link, error) {

	link := models.Link{
		Original: original,
	}
	if err := l.Store.DB.QueryRow("SELECT short FROM links WHERE original = $1", link.Original).Scan(&link.Short); err != nil {
		zap.L().Error("Don't get short URL", zap.Error(err))
		return &link, err
	}
	return &link, nil
}

func (l *Link) GetAllByUserID(userID string) ([]models.Link, error) {

	var links []models.Link

	rows, err := l.Store.DB.Query("SELECT original, short, userid FROM links WHERE userid = $1", userID)
	if err != nil {
		zap.L().Error("Don't get original URL", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var l models.Link
		err := rows.Scan(&l.Original, &l.Short, &l.UserID)
		if err != nil {
			return nil, err
		}

		links = append(links, l)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return links, nil
}

func (l *Link) DelURL(shortURL string, userid string) error {

	row, err := l.Store.DB.Query("UPDATE links SET deleted=true WHERE short=$1 AND userid=$2", shortURL, userid)
	if err != nil {
		return err
	}
	defer row.Close()
	err = row.Err()
	if err != nil {
		return err
	}
	return nil
}
