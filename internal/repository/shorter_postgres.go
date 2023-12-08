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
	// Creating var to check that all is ok
	var short_temp string

	// sql that inserting data in DB, and returning id
	query := fmt.Sprintf("INSERT INTO %s (alias, url) values ($1, $2) RETURNING id", "url")

	// Process query
	row := s.db.QueryRow(query, urlToSave.Short_url, urlToSave.Long_url)
	if err := row.Scan(&short_temp); err != nil{
		return "", err
	}

	return urlToSave.Short_url, nil
}

func (s *ShorterPostgres) GetLongURL(urlToGet *urlshorter.ShortURL) (string, error){
	
	// Var to hold data from DB
	var url_temp string

	// Sql for selecting our data from db
	query := fmt.Sprintf("SELECT url FROM %s WHERE alias = $1", "url")

	// Process sqls
	row := s.db.QueryRow(query, urlToGet.Short_url)
	
	// Scaning given data into our var
	err := row.Scan(&url_temp)

	return url_temp, err
}



