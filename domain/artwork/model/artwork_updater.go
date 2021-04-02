package model

import "time"

type ArtworkUpdater struct {
	ID                   string  `json:"id"`
	ArtistName           *string `json:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath"`
	RequesterName        *string `json:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath"`

	Title       *string `json:"title"`
	TextContext *string `json:"textContext"`
	Favor       *string `json:"favor"`
	Unfavor     *string `json:"unfavor"`

	LastUpdateTime *time.Time `json:"lastUpdateTime" bson:"lastUpdateTime"`
	State          *ArtworkState
}
