package repository

import (
	"context"
	"errors"

	gormModel "github.com/trackhub/api/gorm/model"
	"gorm.io/gorm"
)

type PlaceRepository struct {
	db *gorm.DB
}

func NewPlaceRepository(db *gorm.DB) *PlaceRepository {
	return &PlaceRepository{db: db}
}

func (r *PlaceRepository) GetBySlugOrId(ctx context.Context, id string) (*gormModel.Place, error) {
	// @TODO search slug table for "old" slug, e.g. edited

	var place gormModel.Place
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&place).Error; err == nil {
		return &place, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err := r.db.WithContext(ctx).Where("slug = ?", id).First(&place).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("place " + id + " not found")
		}
		return nil, err
	}

	return &place, nil
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

	tx := q.Limit(limit).Find(&places)

	return places, tx.Error
}
