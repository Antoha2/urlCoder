package service

import (
	"fmt"
	"log"
	"time"

	"github.com/antoha2/urlCoder/repository"
)

func (sImpl *servImpl3) AddLongUrl(url *ServUrl) error {

	repUrl := new(repository.RepLongUrl)
	repUrl.Id = url.Id
	repUrl.Long_url = url.Long_url
	repUrl.Create_at = time.Now()

	err := sImpl.rep.RepAddLongUrl(repUrl)
	if err != nil {
		fmt.Println(err)
		return err
	}
	url.Id = repUrl.Id
	url.Token = repUrl.Token
	log.Println(url)
	return nil
}

func (sImpl *servImpl3) ServGenTokens() error {
	err := sImpl.rep.RepGenTokens()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (sImpl *servImpl3) ServRedirect(url *ServUrl) error {
	repUrl := new(repository.RepLongUrl)
	repUrl.Token = url.Token
	err := sImpl.rep.RepRedirect(repUrl)
	if err != nil {
		//fmt.Println(err)
		return err
	}
	url.Id = repUrl.Id
	url.Long_url = repUrl.Long_url

	return nil
}
