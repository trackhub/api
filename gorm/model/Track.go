package model

const VisibilityPublic = 0

type Track struct {
	ID   string `gorm:"type:uuid;primaryKey"`
	Name string
	Slug *string
}

func (t Track) SlugOrId() string {
	if t.Slug != nil {
		return *t.Slug
	}

	return t.ID
}

func (Track) TableName() string {
	return "track"
}

type Query[T any] interface {
	// @TODO try to use constant from above, isntead of hardcoding "0"

	// SELECT * FROM @@table
	// WHERE (@ids IS NULL OR id NOT IN @ids)
	//   AND visibility = 0
	//   AND point_north_east_lat <= @neLat
	//   AND point_south_west_lat >= @swLat
	//   AND point_north_east_lng <= @neLon
	//   AND point_south_west_lng >= @swLon
	FindAllPublicNotInId(ids []string, limit int, neLat, swLat, neLon, swLon float64) ([]T, error)
}
