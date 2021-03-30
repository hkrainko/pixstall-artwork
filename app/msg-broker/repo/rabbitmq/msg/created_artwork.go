package msg

import (
	"pixstall-artwork/domain/artwork/model"
	"time"
)

type CreatedArtwork struct {
	ID                   string             `json:"id" bson:"id"`
	CommissionID         string             `json:"commissionId" bson:"commissionId"`
	OpenCommissionID     string             `json:"openCommissionId" bson:"openCommissionId"`
	ArtistID             string             `json:"artistId" bson:"artistId"`
	ArtistName           string             `json:"artistName" bson:"artistName"`
	ArtistProfilePath    *string            `json:"artistProfilePath" bson:"artistProfilePath,omitempty"`
	RequesterID          string             `json:"requesterId" bson:"requesterId"`
	RequesterName        string             `json:"requesterName" bson:"requesterName"`
	RequesterProfilePath *string            `json:"requesterProfilePath" bson:"requesterProfilePath,omitempty"`
	IsR18                bool               `json:"isR18" bson:"isR18"`
	Anonymous            bool               `json:"anonymous" bson:"anonymous"`
	Path                 string             `json:"path"`
	Volume               int64              `json:"volume"`
	Size                 model.Size         `json:"size"`
	ContentType          string             `json:"contentType"`
	Rating               int                `json:"rating,omitempty" bson:"rating,omitempty"`
	CompletedTime        time.Time          `json:"completedTime" bson:"completedTime,omitempty"`
	State                model.ArtworkState `json:"state" bson:"state"`
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
		IsR18:                artwork.IsR18,
		Anonymous:            artwork.Anonymous,
		Path:                 artwork.Path,
		Volume:               artwork.Volume,
		Size:                 artwork.Size,
		ContentType:          artwork.ContentType,
		Rating:               artwork.Rating,
		CompletedTime:        artwork.CompletedTime,
		State:                artwork.State,
	}
}
