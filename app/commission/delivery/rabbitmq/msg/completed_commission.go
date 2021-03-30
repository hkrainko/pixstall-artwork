package msg

import (
	"pixstall-artwork/domain/artwork/model"
	"time"
)

type CompletedCommission struct {
	ID               string `json:"id"`
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

	DisplayImage       DisplayImage `json:"displayImage"`
	CompletionFilePath string       `json:"completionFilePath"`
	Rating             int          `json:"rating"`
	Comment            *string      `json:"comment,omitempty"`

	CreateTime    time.Time `json:"createTime" bson:"createTime"`
	CompletedTime time.Time `json:"completedTime" bson:"completedTime,omitempty"`
}

func (c CompletedCommission) ToArtworkCreator() model.ArtworkCreator {

	return model.ArtworkCreator{
		CommissionID:         c.ID,
		OpenCommissionID:     c.OpenCommissionID,
		ArtistID:             c.ArtistID,
		ArtistName:           c.ArtistName,
		ArtistProfilePath:    c.ArtistProfilePath,
		RequesterID:          c.RequesterID,
		RequesterName:        c.RequesterName,
		RequesterProfilePath: c.RequesterProfilePath,
		DayUsed:              c.DayUsed,
		IsR18:                c.IsR18,
		Anonymous:            c.Anonymous,
		Path:                 c.DisplayImage.Path,
		Volume:               c.DisplayImage.Volume,
		Size:                 c.DisplayImage.Size,
		ContentType:          c.DisplayImage.ContentType,
		CompletionFilePath:   c.CompletionFilePath,
		Rating:               c.Rating,
		Comment:              c.Comment,
		CompletedTime:        c.CompletedTime,
	}
}

type DisplayImage struct {
	Path        string     `json:"path" bson:"path"`
	Volume      int64      `json:"volume" bson:"volume"`
	Size        model.Size `json:"size" bson:"size"`
	ContentType string     `json:"contentType" bson:"contentType"`
}
