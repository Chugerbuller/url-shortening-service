package storage

import (
	"database/sql"

	"github.com/Chugerbuller/url-shortening-service/internal/storage/models"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}
func (s *Storage) Add(url models.Url) error {
	query := "INSERT INTO urls (url, short_url, created_at, updated_at) VALUES (:url, :short_url, :created_at, :updated_at)"
	res, err := s.db.Exec(query,
		sql.Named("url", url.URL),
		sql.Named("short_url", url.ShortUrl),
		sql.Named("address", p.Address),
		sql.Named("created_at", p.CreatedAt),
	)
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastID), nil
}
