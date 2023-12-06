package repository

import (
	"errors"

	urlshorter "github.com/stil4004/url-shorter"
)

type ShorterMemory struct {
	long_to_short map[string]string
	short_to_long map[string]string
}

func NewShorterMemory() *ShorterMemory{
	var shorter ShorterMemory
	shorter.long_to_short = make(map[string]string)
	shorter.short_to_long = make(map[string]string)
	return &shorter
}


func (s *ShorterMemory) CreateShortURL(urlToSave  *urlshorter.ShortURL) (string, error){

	if _, found := s.long_to_short[urlToSave.Long_url]; found{
		return "", errors.New("url already in database")
	}
	s.long_to_short[urlToSave.Long_url] = urlToSave.Short_url
	s.short_to_long[urlToSave.Short_url] = urlToSave.Long_url

	return urlToSave.Short_url, nil
}

func (s *ShorterMemory) GetLongURL(urlToGet *urlshorter.ShortURL) (string, error){
	if long_url, found := s.short_to_long[urlToGet.Short_url]; found{
		return long_url, nil
	}
	return "", errors.New("no url connected to this alias")
}



