package model

import "time"

type ArtworkCreator struct {
	CommissionID     string `json:"commissionId" bson:"commissionId"`
	OpenCommissionID string `json:"openCommissionId" bson:"openCommissionId"`

	ArtistID             string  `json:"artistId" bson:"artistId"`
	ArtistName           string  `json:"artistName" bson:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath" bson:"artistProfilePath,omitempty"`
	RequesterID          string  `json:"requesterId" bson:"requesterId"`
	RequesterName        string  `json:"requesterName" bson:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath" bson:"requesterProfilePath,omitempty"`

	DayUsed    time.Duration `json:"dayUsed" bson:"dayUsed"`
	Size       Size          `json:"size" bson:"size"`
	Volume     float64       `json:"volume" bson:"volume"`
	Resolution float64       `json:"resolution" bson:"resolution"`
	Format     string        `json:"format" bson:"format"`
	IsR18      bool          `json:"isR18" bson:"isR18"`
	Anonymous  bool          `json:"anonymous" bson:"anonymous"`

	DisplayImagePath   string  `json:"displayImagePath,omitempty" bson:"displayImagePath,omitempty"`
	CompletionFilePath string  `json:"completionFilePath,omitempty" bson:"completionFilePath,omitempty"`
	Rating             int     `json:"rating,omitempty" bson:"rating,omitempty"`
	Comment            *string `json:"comment,omitempty" bson:"comment,omitempty"`

	CompleteTime time.Time `json:"completeTime" bson:"completeTime,omitempty"`
}
