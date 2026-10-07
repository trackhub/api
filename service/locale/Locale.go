package locale

import (
	gormModel "github.com/trackhub/api/gorm/model"
	graphModel "github.com/trackhub/api/graph/model"
)

func TrackName(t gormModel.Track, locale *graphModel.Locale) *string {
	if locale == nil {
		return t.NameEn
	}

	if locale.String() == graphModel.LocaleBg.String() {
		return t.NameBg
	}

	return t.NameEn
}
