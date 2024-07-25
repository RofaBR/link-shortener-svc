package storage

import (
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

type Link struct {
	ID          int    `db:"id"`
	OriginalURL string `db:"original_url"`
	ShortURL    string `db:"short_url"`
	CreatedAt   string `db:"created_at"`
}

func (s *Storage) CreateLink(link *Link) error {
	_, err := s.db.NamedExec(`INSERT INTO links (original_url, short_url) VALUES (:original_url, :short_url)`, link)
	return err
}

func (s *Storage) GetLinkByShortURL(shortURL string) (*Link, error) {
	var link Link
	err := s.db.Get(&link, `SELECT * FROM links WHERE short_url=$1`, shortURL)
	return &link, err
}
