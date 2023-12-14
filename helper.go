package urlcoder

import "time"

type RecordingDB struct {
	Id       int `gorm:"primaryKey"`
	Token    string
	Long_url string
	CreateAt time.Time
}
