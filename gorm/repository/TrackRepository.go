package repository

import (
	"context"

	gormModel "github.com/trackhub/api/gorm/model"
	"gorm.io/gorm"
)

type TrackRepository struct {
	db *gorm.DB
}

func NewTrackRepository(db *gorm.DB) *TrackRepository {
	return &TrackRepository{db: db}
}

func (r *TrackRepository) generateQuery(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).
		Model(&gormModel.Track{}).
		Preload("TrackVersions").
		Preload("OptimizedPoints")
}

func (r *TrackRepository) FindAllPublicNotInId(ctx context.Context, ids []string, limit int, neLat, swLat, neLon, swLon float64) ([]gormModel.Track, error) {
	var tracks []gormModel.Track

	q := r.generateQuery(ctx).
		Where("visibility = ?", gormModel.VisibilityPublic).
		Where("point_north_east_lat <= ?", neLat).
		Where("point_south_west_lat >= ?", swLat).
		Where("point_north_east_lng <= ?", neLon).
		Where("point_south_west_lng >= ?", swLon)

	if len(ids) > 0 {
		q = q.Where("id NOT IN ?", ids)
	}

	tx := q.Find(&tracks)

	return tracks, tx.Error
}

func (r *TrackRepository) FindLatest(ctx context.Context, trackType, limit int) ([]gormModel.Track, error) {
	var tracks []gormModel.Track

	q := r.generateQuery(ctx).
		Where("visibility = ?", gormModel.VisibilityPublic).
		Where("type = ?", trackType).
		Limit(limit)

	tx := q.Find(&tracks)

	return tracks, tx.Error
}
