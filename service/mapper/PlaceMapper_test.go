package mapper

import (
	"testing"

	gormModel "github.com/trackhub/api/gorm/model"
	"github.com/trackhub/api/service/place"
)

func TestPlaceMapperGormToGraph(t *testing.T) {
	tests := []struct {
		name         string
		place        gormModel.Place
		wantSlugOrID string
		wantIcon     *string
	}{
		{
			name: "maps a place with slug and drinking fountain icon",
			place: gormModel.Place{
				ID:           "place-id",
				NameEn:       "Fountain",
				NameBg:       "Чешма",
				Slug:         stringPointer("city-fountain"),
				Lat:          42.6977,
				Lng:          23.3219,
				Type:         gormModel.PLACE_TYPE_DRINKING_FOUNTAIN,
				IsAttraction: true,
			},
			wantSlugOrID: "city-fountain",
			wantIcon:     stringPointer("/images/trackhub/water/icon.png"),
		},
		{
			name: "falls back to ID and leaves unknown type icon empty",
			place: gormModel.Place{
				ID:     "place-without-slug",
				NameEn: "Unknown place",
				Lat:    1.25,
				Lng:    2.5,
				Type:   99,
			},
			wantSlugOrID: "place-without-slug",
		},
	}

	mapper := PlaceMapper{ImageDetector: place.ImageDetector{}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapper.GormToGraph(&tt.place, nil)

			if got.ID != tt.place.ID {
				t.Errorf("ID = %q, want %q", got.ID, tt.place.ID)
			}
			if got.SlugOrID != tt.wantSlugOrID {
				t.Errorf("SlugOrID = %q, want %q", got.SlugOrID, tt.wantSlugOrID)
			}
			if got.Name == nil || *got.Name != tt.place.NameEn {
				t.Errorf("Name = %v, want %q", got.Name, tt.place.NameEn)
			}
			if got.Attraction != tt.place.IsAttraction {
				t.Errorf("Attraction = %t, want %t", got.Attraction, tt.place.IsAttraction)
			}
			if got.Lat != tt.place.Lat {
				t.Errorf("Lat = %v, want %v", got.Lat, tt.place.Lat)
			}
			if got.Lng != tt.place.Lng {
				t.Errorf("Lng = %v, want %v", got.Lng, tt.place.Lng)
			}
			if !sameStringPointer(got.Icon, tt.wantIcon) {
				t.Errorf("Icon = %v, want %v", got.Icon, tt.wantIcon)
			}
		})
	}
}

func stringPointer(value string) *string {
	return &value
}

func sameStringPointer(a, b *string) bool {
	if a == nil || b == nil {
		return a == b
	}
	return *a == *b
}

