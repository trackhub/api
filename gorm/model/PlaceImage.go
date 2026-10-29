package model

import "time"

type PlaceImage struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	PlaceId   string `gorm:"type:uuid"`
	SendById  int
	Filepath  string
	CreatedAt time.Time
}

func (PlaceImage) TableName() string {
	return "place_image"
}
