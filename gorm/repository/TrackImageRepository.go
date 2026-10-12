package repository

import (
	"context"

	gormModel "github.com/trackhub/api/gorm/model"
	"gorm.io/gorm"
)

type TrackImageRepository struct {
	db *gorm.DB
}

func NewTrackImageRepository(db *gorm.DB) *TrackImageRepository {
	return &TrackImageRepository{db: db}
}

func (r *TrackImageRepository) generateQuery(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).
		Model(&gormModel.TrackImage{})
}

func (r *TrackImageRepository) FindLatest(ctx context.Context, limit int) ([]gormModel.TrackImage, error) {
	var tracks []gormModel.TrackImage

	q := r.generateQuery(ctx).
		Order("created_at DESC").
		Limit(limit)

	tx := q.Find(&tracks)

	return tracks, tx.Error
}
