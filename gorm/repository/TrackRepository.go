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
		Preload("TrackVersions.Points").
		Preload("OptimizedPoints", func(db *gorm.DB) *gorm.DB {
			return db.Order("`order` ASC")
		})
}

// there is no filter by visibility. We won't support "private" visibility, only "unlisted"
func (r *TrackRepository) GetBySlugOrId(ctx context.Context, id string) (*gormModel.Track, error) {
	// @TODO search slug table for "old" slug, e.g. edited

	var track gormModel.Track
	if err := r.generateQuery(ctx).Where("id = ?", id).First(&track).Error; err == nil {
		return &track, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err := r.generateQuery(ctx).Where("slug = ?", id).First(&track).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("track " + id + " not found")
		}
		return nil, err
	}

	return &track, nil
}

func (r *TrackRepository) FindAllPublicNotInId(ctx context.Context, ids []string, limit int, neLat, swLat, neLon, swLon float64) ([]gormModel.Track, error) {
	var tracks []gormModel.Track

	q := r.generateQuery(ctx).
		Where("visibility = ?", gormModel.VisibilityPublic).
		Where("point_north_east_lat >= ?", swLat).
		Where("point_south_west_lat <= ?", neLat).
		Where("point_north_east_lng >= ?", swLon).
		Where("point_south_west_lng <= ?", neLon)

	if len(ids) > 0 {
		q = q.Where("id NOT IN ?", ids)
	}

	tx := q.Limit(limit).Find(&tracks)

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
