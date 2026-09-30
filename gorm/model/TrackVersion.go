package model

type TrackVersion struct {
	ID      string `gorm:"type:uuid;primaryKey"`
	TrackId string `gorm:"type:uuid"`
	Name    string
	FileId  string
}

func (TrackVersion) TableName() string {
	return "version"
}
