package storage

import (
	"database/sql"
	"time"

	"url-shortening-service/internal/models"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}
func (s *Storage) Add(url models.Url) (int, error) {
	query := "INSERT INTO urls (url, short_url, created_at, updated_at) VALUES (:url, :short_url, :created_at, :updated_at)"
	res, err := s.db.Exec(query,
		sql.Named("url", url.Url),
		sql.Named("short_url", url.ShortUrl),
		sql.Named("created_at", url.CreatedAt),
		sql.Named("updated_at", url.UpdatedAt),
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
func (s *Storage) GetByShortUrl(shortUrl string) (*models.Url, error) {
	var url models.Url
	query := "SELECT id, url, short_url, created_at, updated_at FROM urls WHERE short_url = :short_url"
	row := s.db.QueryRow(query, sql.Named("short_url", shortUrl))
	err := row.Scan(&url.ID, &url.Url, &url.ShortUrl, &url.CreatedAt, &url.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &url, nil
}
func (s *Storage) Update(shortUrl, newUrl string) (*models.Url, error) {
	var updatedUrl models.Url

	query := "UPDATE urls SET url = :newUrl,updated_at = :updated_at WHERE short_url = :shortUrl RETURNING id, url, short_url, created_at, updated_at"
	row, err := s.db.Query(query,
		sql.Named("newUrl", newUrl),
		sql.Named("updated_at", time.Now().Unix()),
		sql.Named("shortUrl", shortUrl),
	)
	if err != nil {
		return nil, err
	}
	err = row.Scan(&updatedUrl.ID, &updatedUrl.Url, &updatedUrl.ShortUrl, &updatedUrl.CreatedAt, &updatedUrl.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &updatedUrl, nil
}
