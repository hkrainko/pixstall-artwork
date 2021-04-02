package dao

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pixstall-artwork/domain/artwork/model"
	"time"
)

type Artwork struct {
	ObjectID         primitive.ObjectID `bson:"_id,omitempty"`
	ID               string             `json:"id" bson:"id"`
	CommissionID     string             `json:"commissionId" bson:"commissionId"`
	OpenCommissionID string             `json:"openCommissionId" bson:"openCommissionId"`

	ArtistID             string  `json:"artistId" bson:"artistId"`
	ArtistName           string  `json:"artistName" bson:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath,omitempty" bson:"artistProfilePath,omitempty"`
	RequesterID          string  `json:"requesterId" bson:"requesterId"`
	RequesterName        string  `json:"requesterName" bson:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath,omitempty" bson:"requesterProfilePath,omitempty"`

	DayUsed   time.Duration `json:"dayUsed" bson:"dayUsed"`
	IsR18     bool          `json:"isR18" bson:"isR18"`
	Anonymous bool          `json:"anonymous" bson:"anonymous"`

	Path               string     `json:"path" bson:"path"`
	Volume             int64      `json:"volume" bson:"volume"`
	Size               model.Size `json:"size" bson:"size"`
	ContentType        string     `json:"contentType" bson:"contentType"`
	CompletionFilePath string     `json:"completionFilePath" bson:"completionFilePath"`
	Rating             int        `json:"rating" bson:"rating"`
	Comment            *string    `json:"comment,omitempty" bson:"comment,omitempty"`

	Title       string          `json:"title"`
	TextContent string          `json:"textContent"`
	Views       int             `json:"views"`
	Favors      map[string]bool `json:"favors"`

	CreateTime     time.Time          `json:"createTime" bson:"createTime"`
	StartTime      time.Time          `json:"startTime" bson:"startTime"`
	CompletedTime  time.Time          `json:"completedTime" bson:"completedTime"`
	LastUpdateTime time.Time          `json:"lastUpdateTime" bson:"lastUpdateTime"`
	State          model.ArtworkState `json:"state" bson:"state"`
}

func NewFromArtworkCreator(a model.ArtworkCreator) Artwork {
	now := time.Now()
	return Artwork{
		ObjectID:         primitive.ObjectID{},
		ID:               "aw-" + uuid.NewString(),
		CommissionID:     a.CommissionID,
		OpenCommissionID: a.OpenCommissionID,

		ArtistID:             a.ArtistID,
		ArtistName:           a.ArtistName,
		ArtistProfilePath:    a.ArtistProfilePath,
		RequesterID:          a.RequesterID,
		RequesterName:        a.RequesterName,
		RequesterProfilePath: a.RequesterProfilePath,

		DayUsed:   a.DayUsed,
		IsR18:     a.IsR18,
		Anonymous: a.Anonymous,

		Path:               a.Path,
		Volume:             a.Volume,
		Size:               a.Size,
		ContentType:        a.ContentType,
		CompletionFilePath: a.CompletionFilePath,
		Rating:             a.Rating,
		Comment:            a.Comment,

		Title:       "",
		TextContent: "",
		Views:       0,
		Favors:      make(map[string]bool),

		CreateTime:     now,
		StartTime:      a.StartTime,
		CompletedTime:  a.CompletedTime,
		LastUpdateTime: now,
		State:          model.ArtworkStateActive,
	}
}

func (a Artwork) ToDomainArtwork() model.Artwork {
	return model.Artwork{
		ID:                   a.ID,
		CommissionID:         a.CommissionID,
		OpenCommissionID:     a.OpenCommissionID,
		ArtistID:             a.ArtistID,
		ArtistName:           a.ArtistName,
		ArtistProfilePath:    a.ArtistProfilePath,
		RequesterID:          a.RequesterID,
		RequesterName:        a.RequesterName,
		RequesterProfilePath: a.RequesterProfilePath,

		DayUsed:   a.DayUsed,
		IsR18:     a.IsR18,
		Anonymous: a.Anonymous,

		Path:               a.Path,
		Volume:             a.Volume,
		Size:               a.Size,
		ContentType:        a.ContentType,
		CompletionFilePath: a.CompletionFilePath,
		Rating:             a.Rating,
		Comment:            a.Comment,

		Title:       a.Title,
		TextContent: a.TextContent,
		Views:       a.Views,
		Favors:      a.Favors,
		FavorCount:  len(a.Favors),

		CreateTime:     a.CreateTime,
		StartTime:      a.StartTime,
		CompletedTime:  a.CompletedTime,
		LastUpdateTime: a.LastUpdateTime,
		State:          a.State,
	}
}
