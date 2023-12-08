package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
	urlshorter "github.com/stil4004/url-shorter"
)

// Out shorter interface
type ShorterURL interface {
	CreateShortURL(urlToSave *urlshorter.ShortURL) (string, error)
	GetLongURL(urlToGet *urlshorter.ShortURL) (string, error)
}

// Init services
type Repository struct {
	ShorterURL
	// Add repo here
}

func NewRepository(db *sql.DB, storageType string) *Repository {

	// If user wants to create Postgre DB (adding via argument in command line)
	if storageType == "db"{
		return &Repository{
			ShorterURL: NewShorterPostgres(db),
		}
	}
	
	return &Repository{
		ShorterURL: NewShorterMemory(),
	}
}
