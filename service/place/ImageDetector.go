package place

import (
	"errors"
	"strconv"

	model "github.com/trackhub/api/gorm/model"
)

type ImageDetector struct {
}

func (i ImageDetector) DetectImage(placeType int) (string, error) {
	switch placeType {
	case model.PLACE_TYPE_DRINKING_FOUNTAIN:
		return "/images/trackhub/water/icon.png", nil
	case model.PLACE_TYPE_RESTAURANT:
		return "/images/trackhub/restaurant/icon.png", nil
	}

	return "", errors.New("Unknown type " + strconv.Itoa(placeType))
}
