package model

import "time"

type Artwork struct {
	ID               string `json:"id" bson:"id"`
	CommissionID     string `json:"commissionId" bson:"commissionId"`
	OpenCommissionID string `json:"openCommissionId" bson:"openCommissionId"`

	ArtistID             string  `json:"artistId" bson:"artistId"`
	ArtistName           string  `json:"artistName" bson:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath" bson:"artistProfilePath,omitempty"`
	RequesterID          string  `json:"requesterId" bson:"requesterId"`
	RequesterName        string  `json:"requesterName" bson:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath" bson:"requesterProfilePath,omitempty"`

	Price      Price   `json:"price" bson:"price"`

	Size       Size    `json:"size" bson:"size"`
	Volume     float64 `json:"volume" bson:"volume"`
	Resolution float64 `json:"resolution" bson:"resolution"`
	Format     string  `json:"format" bson:"format"`

	IsR18      bool    `json:"isR18" bson:"isR18"`
	Anonymous  bool    `json:"anonymous" bson:"anonymous"`

	DisplayImagePath   string  `json:"displayImagePath,omitempty" bson:"displayImagePath,omitempty"`
	CompletionFilePath string  `json:"completionFilePath,omitempty" bson:"completionFilePath,omitempty"`
	Rating             int     `json:"rating,omitempty" bson:"rating,omitempty"`
	Comment            *string `json:"comment,omitempty" bson:"comment,omitempty"`

	CreateTime     time.Time    `json:"createTime" bson:"createTime"`
	StartTime      time.Time   `json:"startTime" bson:"startTime"`
	CompleteTime   time.Time   `json:"completeTime" bson:"completeTime"`
	LastUpdateTime time.Time    `json:"lastUpdateTime" bson:"lastUpdateTime"`
	State          ArtworkState `json:"state" bson:"state"`
}

type ArtworkState string

const (
	ArtworkStateActive    ArtworkState = "Active"
	ArtworkStateHidden    ArtworkState = "Hidden"
	ArtworkStateRemoved   ArtworkState = "Removed"
	ArtworkStateForbidden ArtworkState = "Forbidden"
)
