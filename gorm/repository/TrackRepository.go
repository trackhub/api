package repository

import (
	"context"
	"errors"

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

func (r *TrackRepository) GetBySlugOrId(ctx context.Context, id string) (*gormModel.Track, error) {
	var tracks []gormModel.Track

	tx := r.generateQuery(ctx).
		Where("id = ? or slug = ?", id, id).
		Limit(2).
		Find(&tracks)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if len(tracks) == 1 {
		return &tracks[0], nil
	}

	for _, track := range tracks {
		if track.ID == id {
			return &track, nil
		}
	}

	return nil, errors.New("track " + id + " not found")
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
		Order("created_at DESC").
		Limit(limit)

	tx := q.Find(&tracks)

	return tracks, tx.Error
}
