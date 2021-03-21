package msg

import (
	"pixstall-artwork/domain/artwork/model"
	"time"
)

type CreatedArtwork struct {
	ID               string `json:"id" bson:"id"`
	CommissionID     string `json:"commissionId" bson:"commissionId"`
	OpenCommissionID string `json:"openCommissionId" bson:"openCommissionId"`

	ArtistID             string  `json:"artistId" bson:"artistId"`
	ArtistName           string  `json:"artistName" bson:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath" bson:"artistProfilePath,omitempty"`
	RequesterID          string  `json:"requesterId" bson:"requesterId"`
	RequesterName        string  `json:"requesterName" bson:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath" bson:"requesterProfilePath,omitempty"`

	IsR18     bool `json:"isR18" bson:"isR18"`
	Anonymous bool `json:"anonymous" bson:"anonymous"`

	DisplayImagePath string `json:"displayImagePath,omitempty" bson:"displayImagePath,omitempty"`
	Rating           int    `json:"rating,omitempty" bson:"rating,omitempty"`

	CompleteTime time.Time          `json:"completeTime" bson:"completeTime,omitempty"`
	State        model.ArtworkState `json:"state" bson:"state"`
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
		DisplayImagePath:     artwork.DisplayImagePath,
		Rating:               artwork.Rating,
		CompleteTime:         artwork.CompleteTime,
		State:                artwork.State,
	}
}
