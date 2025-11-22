package models

type Url struct {
	ID        int    `db:"id"`
	Url       string `db:"url"`
	ShortUrl  string `db:"short_url"`
	CreatedAt int    `db:"created_at"`
	UpdatedAt int    `db:"updated_at"`
}
