package mapper

import (
	gormModel "github.com/trackhub/api/gorm/model"
	"github.com/trackhub/api/graph/model"
	graphModel "github.com/trackhub/api/graph/model"
	placeService "github.com/trackhub/api/service/place"
)

type PlaceMapper struct {
	ImageDetector placeService.ImageDetector
}

func (m *PlaceMapper) GormToGraph(p *gormModel.Place, locale *model.Locale) *graphModel.Place {
	var icon *string
	calculatedIcon, err := m.ImageDetector.DetectImage(int(p.Type))
	if err == nil {
		icon = &calculatedIcon
	}

	place := model.Place{
		ID:          p.ID,
		SlugOrID:    p.SlugOrId(),
		Slug:        p.Slug,
		Name:        &p.NameEn, // @TODO localize
		Description: &p.DescriptionEn,
		Attraction:  p.IsAttraction,
		Icon:        icon,
		Lat:         p.Lat,
		Lng:         p.Lng,
	}

	return &place
}
