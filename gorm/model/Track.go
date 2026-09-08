package model

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
	// SELECT * FROM @@table
	FindAll() ([]T, error)
}