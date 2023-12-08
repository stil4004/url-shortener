package service

import (
	urlshorter "github.com/stil4004/url-shorter"
	"github.com/stil4004/url-shorter/internal/repository"
)

// Stuct for our url shorter service
type ShorterURL interface {
	CreateShortURL(urlToSave *urlshorter.ShortURL) (string, error)
	GetLongURL(urlToGet *urlshorter.ShortURL) (string, error)
}


type Service struct {
	ShorterURL
	// Add services here
}

// Function for initng our service
func NewService(repos *repository.Repository) *Service {
	return &Service{
		ShorterURL: NewShorterService(repos.ShorterURL),
	}
}