package service

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	urlshorter "github.com/stil4004/url-shorter"
	"github.com/stil4004/url-shorter/internal/repository"
	"github.com/stil4004/url-shorter/internal/repository/db"
)

type ShortService struct {
	repo repository.ShorterURL
}

func NewShorterService(repo repository.ShorterURL) *ShortService{
	return &ShortService{
		repo: repo,
	}
}

func (s *ShortService) CreateShortURL(urlToSave *urlshorter.ShortURL) (string, error){
	// TODO: Create alias
	new_alias := s.GenerateAlias()
	for _, found := db.All_alias[new_alias]; found;{
		new_alias = s.GenerateAlias()
	}
	urlToSave.Short_url = new_alias
	return s.repo.CreateShortURL(urlToSave)
}

func (s *ShortService) GetLongURL(urlToGet *urlshorter.ShortURL) (string, error){	
	// TODO: Add validation or hz blya ne pridumal
	return s.repo.GetLongURL(urlToGet)
}

const alphabet = "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func (s *ShortService) GenerateAlias() string{
	alias := bytes.Buffer{}
	uniqueTime := time.Now().Unix()
	_, _ = fmt.Fprintf(&alias, "%s", convert(uniqueTime, int64(len(alphabet))))

	for len(alias.String()) < 10 {
		rand.Seed(time.Now().UnixNano())
		number := rand.Intn(len(alphabet))
		_, _ = fmt.Fprintf(&alias, "%c", alphabet[int64(number)])

	}
	return alias.String()
}

func convert(decimalNumber, n int64) string {
	buf := bytes.Buffer{}
	for decimalNumber > 0 {
		curNumber := decimalNumber % n
		decimalNumber /= n
		_, _ = fmt.Fprintf(&buf, "%c", alphabet[curNumber])
	}
	return buf.String()
}