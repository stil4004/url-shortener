package service

import (
	urlshorter "github.com/stil4004/url-shorter"
	"github.com/stil4004/url-shorter/internal/repository"
)

type ShortService struct {
	repo repository.ShorterURL
}

func NewShorterService(repo repository.ShorterURL) *ShortService{
	return &ShortService{
		repo: repo,
	}
}

func (s *ShortService) CreateShortURL(urlToSave urlshorter.ShortURL) (string, error){
	// TODO: Create alias
	urlToSave.Short_url = urlToSave.Long_url + "123"
	return s.repo.CreateShortURL(urlToSave)
}