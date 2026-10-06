package repository

import (
	"context"

	gormModel "github.com/trackhub/api/gorm/model"
	"gorm.io/gorm"
)

type PlaceRepository struct {
	db *gorm.DB
}

func NewPlaceRepository(db *gorm.DB) *PlaceRepository {
	return &PlaceRepository{db: db}
}

func (r *PlaceRepository) FindAllPublicNotInId(ctx context.Context, ids []string, limit int, neLat, swLat, neLon, swLon float64) ([]gormModel.Place, error) {
	var places []gormModel.Place

	q := r.db.WithContext(ctx).
		Model(&gormModel.Place{}).
		Where("lat <= ?", neLat).
		Where("lat >= ?", swLat).
		Where("lng <= ?", neLon).
		Where("lng >= ?", swLon)

	if len(ids) > 0 {
		q = q.Where("id NOT IN ?", ids)
	}

	tx := q.Find(&places)

	return places, tx.Error
}
