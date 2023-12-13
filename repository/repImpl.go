package repository

import "fmt"

func (*repositoryImplDB) Create(unit *RepUrl) error {
	fmt.Println("rep ", unit)
	return nil
}
