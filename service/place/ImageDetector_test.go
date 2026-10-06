package place

import (
	"testing"

	model "github.com/trackhub/api/gorm/model"
)

func TestImageDetectorDetectImage(t *testing.T) {
	tests := []struct {
		name          string
		placeType     int
		wantImagePath string
		wantError     string
	}{
		{
			name:          "drinking fountain",
			placeType:     model.PLACE_TYPE_DRINKING_FOUNTAIN,
			wantImagePath: "/images/trackhub/water/icon.png",
		},
		{
			name:          "restaurant",
			placeType:     model.PLACE_TYPE_RESTAURANT,
			wantImagePath: "/images/trackhub/restaurant/icon.png",
		},
		{
			name:      "unknown type",
			placeType: 99,
			wantError: "Unknown type 99",
		},
	}

	detector := ImageDetector{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotImagePath, err := detector.DetectImage(tt.placeType)

			if tt.wantError != "" {
				if err == nil {
					t.Fatalf("DetectImage(%d) error = nil, want %q", tt.placeType, tt.wantError)
				}
				if err.Error() != tt.wantError {
					t.Errorf("DetectImage(%d) error = %q, want %q", tt.placeType, err.Error(), tt.wantError)
				}
				if gotImagePath != "" {
					t.Errorf("DetectImage(%d) image path = %q, want empty", tt.placeType, gotImagePath)
				}
				return
			}

			if err != nil {
				t.Fatalf("DetectImage(%d) unexpected error: %v", tt.placeType, err)
			}
			if gotImagePath != tt.wantImagePath {
				t.Errorf("DetectImage(%d) = %q, want %q", tt.placeType, gotImagePath, tt.wantImagePath)
			}
		})
	}
}
