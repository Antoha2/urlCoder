package service

import (
	"fmt"

	"github.com/antoha2/urlCoder/repository"
)

func (sImpl *servImpl3) Create(unit *ServBuyUnit) error {
	fmt.Println("serv ", unit)

	repUnit := new(repository.RepBuyUnit)
	repUnit.Id = unit.Id
	repUnit.UserId = unit.Id
	repUnit.Text = unit.Text
	repUnit.IsDone = unit.IsDone

	err := sImpl.rep.Create(repUnit)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
