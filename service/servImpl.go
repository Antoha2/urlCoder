package service

import (
	"fmt"
	"time"

	"github.com/antoha2/urlCoder/repository"
)

func (sImpl *servImpl3) AddLongUrl(url *ServUrl) error {

	repUrl := new(repository.RepLongUrl)
	repUrl.Id = url.Id
	repUrl.Long_url = url.Long_url
	repUrl.CreateAt = time.Now()
	repUrl.Token = sImpl.hashid(url.Id)

	err := sImpl.rep.RepAddLongUrl(repUrl)
	if err != nil {
		fmt.Println(err)
		return err
	}
	url.Id = repUrl.Id
	url.Token = repUrl.Token
	return nil
}

// func (sImpl *servImpl3) hashid(id int) string {
// 	hd := hashids.NewData()
// 	hd.Salt = cfg.HashSalt
// 	hd.MinLength = cfg.HashMinLength

// 	h := hashids.NewWithData(hd)
// 	token, _ := h.Encode([]int{id})

// 	// fmt.Println(e)
// 	// d, _ := h.DecodeWithError(e)
// 	// fmt.Println(d)
// 	return token
// }

func (sImpl *servImpl3) ServGenTokens(q int) error {
	err := sImpl.rep.RepGenTokens(q)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
