package get_artwork

import (
	"pixstall-artwork/domain/artwork/model"
	"time"
)

type Response struct {
	Artwork Artwork `json:",inline"`
}

type Artwork struct {
	ID               string `json:"id"`
	OpenCommissionID string `json:"openCommissionId"`

	ArtistID             string  `json:"artistId"`
	ArtistName           string  `json:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath"`
	RequesterID          *string `json:"requesterId,omitempty"`
	RequesterName        *string `json:"requesterName,omitempty"`
	RequesterProfilePath *string `json:"requesterProfilePath,omitempty"`

	Size       model.Size `json:"size"`
	Volume     float64    `json:"volume"`
	Resolution float64    `json:"resolution"`
	Format     string     `json:"format"`

	IsR18     bool `json:"isR18" bson:"isR18"`
	Anonymous bool `json:"anonymous" bson:"anonymous"`

	DisplayImagePath string `json:"displayImagePath"`
	Rating           int    `json:"rating"`

	CreateTime     time.Time          `json:"createTime"`
	StartTime      time.Time          `json:"startTime"`
	CompleteTime   time.Time          `json:"completeTime"`
	LastUpdateTime time.Time          `json:"lastUpdateTime"`
	State          model.ArtworkState `json:"state"`
}

func NewResponse(a model.Artwork) *Response {
	return &Response{
		Artwork: NewRespArtworkFormDomainArtwork(a),
	}
}

func NewRespArtworkFormDomainArtwork(a model.Artwork) Artwork {
	var requesterID *string
	var requesterName *string
	var requesterProfilePath *string
	if a.Anonymous {
		requesterID = &a.RequesterID
		requesterName = &a.RequesterName
		requesterProfilePath = a.RequesterProfilePath
	}
	return Artwork{
		ID:                   a.ID,
		OpenCommissionID:     a.OpenCommissionID,
		ArtistID:             a.ArtistID,
		ArtistName:           a.ArtistName,
		ArtistProfilePath:    a.ArtistProfilePath,
		RequesterID:          requesterID,
		RequesterName:        requesterName,
		RequesterProfilePath: requesterProfilePath,
		Size:                 a.Size,
		Volume:               a.Volume,
		Resolution:           a.Resolution,
		Format:               a.Format,
		IsR18:                a.IsR18,
		Anonymous:            a.Anonymous,
		DisplayImagePath:     a.DisplayImagePath,
		Rating:               a.Rating,
		CreateTime:           a.CreateTime,
		StartTime:            a.StartTime,
		CompleteTime:         a.CompleteTime,
		LastUpdateTime:       a.LastUpdateTime,
		State:                a.State,
	}
}