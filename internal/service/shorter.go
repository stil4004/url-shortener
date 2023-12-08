package service

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

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

// Creating of alias, and sending it to repo
func (s *ShortService) CreateShortURL(urlToSave *urlshorter.ShortURL) (string, error){

	// Validating out url
	err := validateUrl(urlToSave.Long_url)
	if err != nil{
		return "", err
	}

	// Creating alias
	new_alias := s.GenerateAlias()

	// Inserting alias into model var to send it to repo
	urlToSave.Short_url = new_alias

	return s.repo.CreateShortURL(urlToSave)
}

func (s *ShortService) GetLongURL(urlToGet *urlshorter.ShortURL) (string, error){	
	
	// Validating that alias        (yep we could just check if it empty, but if we want to have more conditions we could add them here)
	err := validateAlias(urlToGet.Short_url)
	if err != nil{
		return "", err
	}

	return s.repo.GetLongURL(urlToGet)
}

// Alphabet for creating alias
const alphabet = "_-ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// Function to generate short url
func (s *ShortService) GenerateAlias() string{

	alias := bytes.Buffer{}
	uniqueTime := time.Now().Unix()
	_, _ = fmt.Fprintf(&alias, "%s", convert(uniqueTime, int64(len(alphabet))))

	// 10 times adding random var to alias
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


// Validating given url
func validateUrl(str string) error{
	if str == ""{
		return errors.New("requested empty url")
	}
	if !strings.HasPrefix(str, `http:\\`)  && !strings.HasPrefix(str, `https:\\`) {
		return errors.New("the given string is not url")
	}
	return nil
}


func validateAlias(str string) error{
	if str == ""{
		return errors.New("requested alias is empty")
	}
	return nil
}