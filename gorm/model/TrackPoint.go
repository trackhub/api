package model

type TrackPoint struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	Order     int
	Distance  float64
	Version   string
	Elevation float64
}

func (TrackPoint) TableName() string {
	return "track_point"
}
