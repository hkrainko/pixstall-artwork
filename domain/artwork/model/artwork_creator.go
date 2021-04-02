package model

import "time"

type ArtworkCreator struct {
	CommissionID     string `json:"commissionId"`
	OpenCommissionID string `json:"openCommissionId"`

	ArtistID             string  `json:"artistId"`
	ArtistName           string  `json:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath,omitempty"`
	RequesterID          string  `json:"requesterId"`
	RequesterName        string  `json:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath,omitempty"`

	DayUsed   time.Duration `json:"dayUsed"`
	IsR18     bool          `json:"isR18"`
	Anonymous bool          `json:"anonymous"`

	Path               string  `json:"path"`
	Volume             int64   `json:"volume"`
	Size               Size    `json:"size"`
	ContentType        string  `json:"contentType"`
	CompletionFilePath string  `json:"completionFilePath"`
	Rating             int     `json:"rating"`
	Comment            *string `json:"comment,omitempty"`

	StartTime     time.Time `json:"startTime"`
	CompletedTime time.Time `json:"completedTime"`
}
