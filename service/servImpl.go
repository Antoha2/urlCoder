package service

import (
	"fmt"
	"time"

	"github.com/antoha2/urlCoder/repository"
)

func (sImpl *servImpl3) LongUrl(url *ServUrl) error {
	

	repUrl := new(repository.RepUrl)
	repUrl.Id = url.Id
	repUrl.Long_url = url.Long_url
	repUrl.CreateAt = time.Now()
	// // repUnit.Text = unit.Text
	// // repUnit.IsDone = unit.IsDone

	err := sImpl.rep.LongUrl(repUrl)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
