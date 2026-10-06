package model

const VisibilityPublic = 0

type Track struct {
	ID              string `gorm:"type:uuid;primaryKey"`
	Name            string
	Slug            *string
	Type            int
	TrackVersions   []TrackVersion   `gorm:"foreignKey:TrackId"`
	OptimizedPoints []OptimizedPoint `gorm:"foreignKey:TrackId"`
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

type OptimizedPoint struct {
	ID           string `gorm:"type:uuid;primaryKey"`
	TrackId      string
	Order        int
	Lat          float64
	Lng          float64
	VersionIndex int
}

func (op OptimizedPoint) TableName() string {
	return "optimized_point"
}
