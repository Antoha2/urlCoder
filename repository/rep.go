package repository

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	RepAddLongUrl(url *RepLongUrl) error
	RepGenTokens(q int) error
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

type RepLongUrl struct {
	Id       int `gorm:"primaryKey"`
	Long_url string
	Token    string
	CreateAt time.Time
}
