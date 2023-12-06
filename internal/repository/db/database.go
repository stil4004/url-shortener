package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	urlshorter "github.com/stil4004/url-shorter"
)

const(
	urlTable = "url"
)

var(
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists = errors.New("url already exists")
)

type Config struct{
	Host    string
	Port     string
	Username string
	Password string
	DBName 	 string
	SSLMode  string
}

type DataBase struct {
	db *sql.DB
}

func New(cfg Config) (*sql.DB, error){

	// Формирование строки подключения
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName,
	)

	// Подключение к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// TODO: если проект расширять то можно добавить миграцию
	migr, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS url(
		id SERIAL PRIMARY KEY,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL UNIQUE);
	`)
	if err != nil{
		return nil, err
	}

	_, err = migr.Exec()
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	dataB := &DataBase{db: db}
	All_alias = make(map[string]bool)

	temp_allias_arr, err := dataB.GetAllShortURLS()
	if err != nil{
		log.Println("couldn't reach alias")
	}
	for _, a := range temp_allias_arr{
		All_alias[a] = true
	}

	return db, nil
}

func (d *DataBase) SaveURL(urlToSave urlshorter.ShortURL) (int, error){

	var id int

	query := fmt.Sprintf("INSERT INTO %s (alias, url) values ($1, $2) RETURNING id", urlTable)
	row := d.db.QueryRow(query, urlToSave.Short_url, urlToSave.Long_url)
	if err := row.Scan(&id); err != nil{
		return 0, err
	}
	return id, nil
}

func (d *DataBase) GetURLbyAlias(alias string) (string, error){

	query := fmt.Sprintf("SELECT url FROM %s WHERE alias = $1", urlTable)

	row := d.db.QueryRow(query, alias)
	
	var long_url string
	err := row.Scan(&long_url)
	return long_url, err
}

func (d *DataBase) DeleteURLbyAlias(alias string) error{
	query := fmt.Sprintf("DELETE FROM %s WHERE alias = $1;", urlTable)

	_, err := d.db.Exec(query, alias)
	if err != nil {
  		return err
	}
	return nil
}

var All_alias map[string]bool

func (d *DataBase) GetAllShortURLS() ([]string, error){
	allias_arr := []string{}

	query := fmt.Sprintf("SELECT alias FROM %s WHERE", "url")

	rows, err := d.db.Query(query)
	if err != nil {
        return nil, err
    }
    defer rows.Close()
	for rows.Next() {
        var al string
        if err := rows.Scan(&al); err != nil {
            return allias_arr, err
        }
        allias_arr = append(allias_arr, al)
    }
    if err = rows.Err(); err != nil {
        return allias_arr, err
    }
	return allias_arr, nil
}