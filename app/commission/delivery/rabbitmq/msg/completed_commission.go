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

	DayUsed    time.Duration `json:"dayUsed"`
	Size       model.Size    `json:"size"`
	Volume     float64       `json:"volume"`
	Resolution float64       `json:"resolution"`
	Format     string        `json:"format"`
	IsR18      bool          `json:"isR18"`
	Anonymous  bool          `json:"anonymous"`

	DisplayImagePath   string  `json:"displayImagePath"`
	CompletionFilePath string  `json:"completionFilePath"`
	Rating             int     `json:"rating"`
	Comment            *string `json:"comment,omitempty"`

	CreateTime   time.Time `json:"createTime" bson:"createTime"`
	CompleteTime time.Time `json:"completeTime" bson:"completeTime,omitempty"`
}

func (c CompletedCommission) ToArtworkCreator() model.ArtworkCreator {

	dayUsed := c.CompleteTime.Sub(c.CreateTime)

	return model.ArtworkCreator{
		CommissionID:         c.ID,
		OpenCommissionID:     c.OpenCommissionID,
		ArtistID:             c.ArtistID,
		ArtistName:           c.ArtistName,
		ArtistProfilePath:    c.ArtistProfilePath,
		RequesterID:          c.RequesterID,
		RequesterName:        c.RequesterName,
		RequesterProfilePath: c.RequesterProfilePath,
		DayUsed:              dayUsed,
		Size:                 c.Size,
		Volume:               0,
		Resolution:           c.Resolution,
		Format:               c.Format,
		IsR18:                c.IsR18,
		Anonymous:            c.Anonymous,
		DisplayImagePath:     c.DisplayImagePath,
		CompletionFilePath:   c.CompletionFilePath,
		Rating:               c.Rating,
		Comment:              c.Comment,
		CompleteTime:         c.CreateTime,
	}
}
