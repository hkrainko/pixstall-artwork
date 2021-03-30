package dao

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pixstall-artwork/domain/artwork/model"
	"time"
)

type Artwork struct {
	ObjectID      primitive.ObjectID `bson:"_id,omitempty"`
	model.Artwork `bson:",inline"`
}

func NewFromArtworkCreator(a model.ArtworkCreator) Artwork {
	now := time.Now()
	return Artwork{
		ObjectID: primitive.ObjectID{},
		Artwork: model.Artwork{
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

			CreateTime:     now,
			CompletedTime:  a.CompletedTime,
			LastUpdateTime: now,
			State:          model.ArtworkStateActive,
		},
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

		CreateTime:     a.CreateTime,
		CompletedTime:   a.CompletedTime,
		LastUpdateTime: a.LastUpdateTime,
		State:          a.State,
	}
}
