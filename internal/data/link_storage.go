package data

import (
	"database/sql"
	"errors"
)

type LinkStorage interface {
	CreateLink(url string) (string, error)
	GetLink(shortUrl string) (string, error)
}

type PostgresLinkStorage struct {
	db *sql.DB
}

func NewPostgresLinkStorage(db *sql.DB) *PostgresLinkStorage {
	return &PostgresLinkStorage{db: db}
}

func (s *PostgresLinkStorage) CreateLink(url string) (string, error) {
	/*shortURL :=

	_, err := s.db.Exec("INSERT INTO links (short_url, original_url) VALUES ($1, $2)", shortURL, url)
	if err != nil {
		return "", err
	}

	return shortURL, nil
	*/
}

func (s *PostgresLinkStorage) GetLink(shortURL string) (string, error) {
	var originalURL string
	err := s.db.QueryRow("SELECT original_url FROM links WHERE short_url = $1", shortURL).Scan(&originalURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("link not found")
		}
		return "", err
	}

	return originalURL, nil
}
