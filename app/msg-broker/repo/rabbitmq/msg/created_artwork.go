package msg

import (
	"pixstall-artwork/domain/artwork/model"
	"time"
)

type CreatedArtwork struct {
	ID                   string             `json:"id" bson:"id"`
	CommissionID         string             `json:"commissionId"`
	OpenCommissionID     string             `json:"openCommissionId"`
	ArtistID             string             `json:"artistId"`
	ArtistName           string             `json:"artistName"`
	ArtistProfilePath    *string            `json:"artistProfilePath,omitempty"`
	RequesterID          string             `json:"requesterId"`
	RequesterName        string             `json:"requesterName"`
	RequesterProfilePath *string            `json:"requesterProfilePath,omitempty"`
	DayUsed              time.Duration      `json:"dayUsed"`
	IsR18                bool               `json:"isR18"`
	Anonymous            bool               `json:"anonymous"`
	Path                 string             `json:"path"`
	Volume               int64              `json:"volume"`
	Size                 model.Size         `json:"size"`
	ContentType          string             `json:"contentType"`
	Rating               int                `json:"rating"`
	StartTime            time.Time          `json:"startTime"`
	CompletedTime        time.Time          `json:"completedTime"`
	State                model.ArtworkState `json:"state"`
}

func NewCreatedArtwork(artwork model.Artwork) CreatedArtwork {
	return CreatedArtwork{
		ID:                   artwork.ID,
		CommissionID:         artwork.CommissionID,
		OpenCommissionID:     artwork.OpenCommissionID,
		ArtistID:             artwork.ArtistID,
		ArtistName:           artwork.ArtistName,
		ArtistProfilePath:    artwork.ArtistProfilePath,
		RequesterID:          artwork.RequesterID,
		RequesterName:        artwork.RequesterName,
		RequesterProfilePath: artwork.RequesterProfilePath,
		DayUsed:              artwork.DayUsed,
		IsR18:                artwork.IsR18,
		Anonymous:            artwork.Anonymous,
		Path:                 artwork.Path,
		Volume:               artwork.Volume,
		Size:                 artwork.Size,
		ContentType:          artwork.ContentType,
		Rating:               artwork.Rating,
		StartTime:            artwork.StartTime,
		CompletedTime:        artwork.CompletedTime,
		State:                artwork.State,
	}
}
