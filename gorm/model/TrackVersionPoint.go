package model

type TrackVersionPoint struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	Order     int
	Distance  float64
	VersionId string
	Elevation float64
}

func (TrackVersionPoint) TableName() string {
	return "point"
}
