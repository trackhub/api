package model

type Place struct {
	ID           string `gorm:"type:uuid;primaryKey"`
	NameEn       string
	NameBg       string
	Slug         *string
	Lat          float64
	Lng          float64
	Type         int16
	IsAttraction bool
}

func (Place) TableName() string {
	return "place"
}

func (t Place) SlugOrId() string {
	if t.Slug != nil {
		return *t.Slug
	}

	return t.ID
}
