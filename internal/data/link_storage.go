package data

import (
	"database/sql"
	"errors"

	"github.com/teris-io/shortid"
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
	shortURL, err := shortid.Generate()
	if err != nil {
		return "", err
	}

	_, err = s.db.Exec("INSERT INTO links (original_url, short_url) VALUES ($1, $2)", url, shortURL)
	if err != nil {
		return "", err
	}

	return shortURL, nil
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
