package service

import (
	"github.com/antoha2/urlCoder/repository"
)

type UrlService interface {
	AddLongUrl(url *ServUrl) error
	ServGenTokens() error
	ServRedirect(url *ServUrl) error
}

type servImpl3 struct {
	rep repository.Repository
	UrlService
}

func NewService(rep repository.Repository) *servImpl3 {
	return &servImpl3{
		rep: rep,
	}
}

type ServUrl struct {
	Id       int    `json:"id"`
	Long_url string `json:"long_url"`
	Token    string `token`
}
