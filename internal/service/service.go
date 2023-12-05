package service

import (
	urlshorter "github.com/stil4004/url-shorter"
	"github.com/stil4004/url-shorter/internal/repository"
)

type ShorterURL interface {
	CreateShortURL(urlToSave urlshorter.ShortURL) (string, error)
}

type Service struct {
	ShorterURL
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		ShorterURL: NewShorterService(repos.ShorterURL),
	}
}