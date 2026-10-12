package model

import (
	"time"
)

type TrackImage struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	TrackId   string
	Filepath  string
	CreatedAt time.Time
}

func (i *TrackImage) TableName() string {
	return "track_image"
}
