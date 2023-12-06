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


func (s *ShorterPostgres) CreateShortURL(urlToSave  *urlshorter.ShortURL) (string, error){

	var short_temp string

	query := fmt.Sprintf("INSERT INTO %s (alias, url) values ($1, $2) RETURNING id", "url")
	row := s.db.QueryRow(query, urlToSave.Short_url, urlToSave.Long_url)
	if err := row.Scan(&short_temp); err != nil{
		return "", err
	}

	return urlToSave.Short_url, nil
}

func (s *ShorterPostgres) GetLongURL(urlToGet *urlshorter.ShortURL) (string, error){
	var url_temp string

	query := fmt.Sprintf("SELECT url FROM %s WHERE alias = $1", "url")

	row := s.db.QueryRow(query, urlToGet.Short_url)
	
	err := row.Scan(&url_temp)

	return url_temp, err
}



