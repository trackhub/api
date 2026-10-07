package mapper

import (
	gormModel "github.com/trackhub/api/gorm/model"
	"github.com/trackhub/api/graph/model"
	graphModel "github.com/trackhub/api/graph/model"
)

func GormTrackToGraphTrack(t gormModel.Track) *graphModel.Track {
	track := &graphModel.Track{
		ID:       t.ID,
		SlugOrID: t.SlugOrId(),
		Type:     t.Type,
		Versions: make([]*graphModel.TrackVersion, 0),
	}

	for _, trackVersion := range t.TrackVersions {
		track.Versions = append(track.Versions, &model.TrackVersion{
			ID: trackVersion.ID,
		})
	}

	optimizedPoints := make([]*model.Point, 0)
	for _, optimizedPoint := range t.OptimizedPoints {
		optimizedPoints = append(optimizedPoints, &model.Point{
			Lat: optimizedPoint.Lat,
			Lng: optimizedPoint.Lng,
		})

		track.OptimizedPoints = make([][]*model.Point, 0)
		track.OptimizedPoints = append(track.OptimizedPoints, optimizedPoints)
	}

	return track
}
