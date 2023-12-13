package repository

import "gorm.io/gorm"

type Repository interface {
	Create(unit *RepUrl) error
}

type repositoryImplDB struct {
	Repository
	rep *gorm.DB
}

func NewRepository(dbx *gorm.DB) *repositoryImplDB {
	return &repositoryImplDB{
		rep: dbx,
	}
}

type RepUrl struct {
	Id     int
	UserId int
	Text   string
	IsDone bool
}
