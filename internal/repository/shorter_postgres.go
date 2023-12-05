package repository

import (
	"database/sql"
	"fmt"

	urlshorter "github.com/stil4004/url-shorter"
)

type ShorterPostgres struct {
	db *sql.DB
}

func NewShorterPostgres(db *sql.DB) *ShorterPostgres{
	return &ShorterPostgres{db: db}
}


func (s *ShorterPostgres) CreateShortURL(url urlshorter.ShortURL) (string, error){

	var short_temp string

	query := fmt.Sprintf("INSERT INTO %s (alias, url) values ($1, $2) RETURNING id", "url")
	row := s.db.QueryRow(query, url.Short_url, url.Long_url)
	if err := row.Scan(&short_temp); err != nil{
		return "", err
	}
	
	return url.Short_url, nil
}

