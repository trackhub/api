package model

type TrackSlug struct {
	ID    string `gorm:"type:uuid;primaryKey"`
	Track string
	Slug  string
}

func (TrackSlug) TableName() string {
	return "track_slug"
}
