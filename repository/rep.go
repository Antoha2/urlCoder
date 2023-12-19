package repository

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	RepAddLongUrl(url *RepLongUrl) error
	RepGenTokens() error
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
	Id        int `gorm:"primaryKey, column:url_id"`
	Long_url  string
	Token_id  int
	Token     string
	Create_at time.Time
	Expiry_at time.Time
}

type RepToken struct {
	Id    int `gorm:"column:token_id"`
	Token string
	Used  bool `gorm:"default:false"`
}
