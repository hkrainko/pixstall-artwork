package model

import "time"

type ArtworkUpdater struct {
	ID                   string  `json:"id"`
	ArtistName           *string `json:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath"`
	RequesterName        *string `json:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath"`

	LastUpdateTime *time.Time `json:"lastUpdateTime" bson:"lastUpdateTime"`
	State          *ArtworkState
}
