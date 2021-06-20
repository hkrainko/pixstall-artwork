package get_artwork

import (
	"pixstall-artwork/domain/artwork/model"
	"time"
)

type Response struct {
	Artwork `json:",inline"`
}

type Artwork struct {
	ID               string  `json:"id"`
	CommissionID     *string `json:"commissionId,omitempty"`
	OpenCommissionID string  `json:"openCommissionId"`

	ArtistID             string  `json:"artistId"`
	ArtistName           string  `json:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath"`
	RequesterID          *string `json:"requesterId,omitempty"`
	RequesterName        *string `json:"requesterName,omitempty"`
	RequesterProfilePath *string `json:"requesterProfilePath,omitempty"`

	DayUsed   time.Duration `json:"dayUsed"`
	IsR18     bool          `json:"isR18"`
	Anonymous bool          `json:"anonymous"`

	Path        string     `json:"path"`
	Volume      int64      `json:"volume"`
	Size        model.Size `json:"size"`
	ContentType string     `json:"contentType"`
	Rating      int        `json:"rating"`

	Title       string `json:"title"`
	TextContent string `json:"textContent"`
	Views       int    `json:"views"`
	Favored     bool   `json:"favored"`
	FavorCount  int    `json:"favorCount"`

	CreateTime     time.Time          `json:"createTime"`
	StartTime      time.Time          `json:"startTime"`
	CompletedTime  time.Time          `json:"completedTime"`
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
	if !a.Anonymous {
		requesterID = &a.RequesterID
		requesterName = &a.RequesterName
		requesterProfilePath = a.RequesterProfilePath
	}
	favored := false
	//if userID != nil {
	//	if _, ok := a.Favors[*userID]; ok {
	//		favored = true
	//	}
	//}
	return Artwork{
		ID:                   a.ID,
		OpenCommissionID:     a.OpenCommissionID,
		ArtistID:             a.ArtistID,
		ArtistName:           a.ArtistName,
		ArtistProfilePath:    a.ArtistProfilePath,
		RequesterID:          requesterID,
		RequesterName:        requesterName,
		RequesterProfilePath: requesterProfilePath,

		DayUsed:   a.DayUsed,
		IsR18:     a.IsR18,
		Anonymous: a.Anonymous,

		Path:        a.Path,
		Volume:      a.Volume,
		Size:        a.Size,
		ContentType: a.ContentType,
		Rating:      a.Rating,

		Title:       a.Title,
		TextContent: a.TextContent,
		Views:       a.Views,
		Favored:     favored,
		FavorCount:  len(a.Favors),

		CreateTime:     a.CreateTime,
		StartTime:      a.StartTime,
		CompletedTime:  a.CompletedTime,
		LastUpdateTime: a.LastUpdateTime,
		State:          a.State,
	}
}
