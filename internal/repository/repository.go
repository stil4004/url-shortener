package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
	urlshorter "github.com/stil4004/url-shorter"
)

type ShorterURL interface {
	CreateShortURL(urlToSave *urlshorter.ShortURL) (string, error)
	GetLongURL(urlToGet *urlshorter.ShortURL) (string, error)
}

type Repository struct {
	ShorterURL
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		ShorterURL: NewShorterPostgres(db),
	}
}
