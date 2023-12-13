package repository

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	LongUrl(unit *RepUrl) error
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
	Id       int
	Long_url string
	CreateAt time.Time
}
