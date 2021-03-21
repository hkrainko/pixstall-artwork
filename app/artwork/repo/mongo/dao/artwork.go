package dao

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pixstall-artwork/domain/artwork/model"
	"time"
)

type Artwork struct {
	ObjectID primitive.ObjectID `bson:"_id,omitempty"`
	model.Artwork `bson:",inline"`
}

func NewFromArtworkCreator(a model.ArtworkCreator) Artwork {
	now := time.Now()
	return Artwork{
		ObjectID: primitive.ObjectID{},
		Artwork: model.Artwork{
			ID:                   "aw-" + uuid.NewString(),
			CommissionID:         a.CommissionID,
			OpenCommissionID:     a.OpenCommissionID,
			ArtistID:             a.ArtistID,
			ArtistName:           a.ArtistName,
			ArtistProfilePath:    a.ArtistProfilePath,
			RequesterID:          a.RequesterID,
			RequesterName:        a.RequesterName,
			RequesterProfilePath: a.RequesterProfilePath,
			Price: a.Price,
			DayUsed: a.DayUsed,
			Size: a.Size,
			Volume:             a.Volume,
			Resolution:         a.Resolution,
			Format:             a.Format,
			IsR18:              a.IsR18,
			Anonymous:          a.Anonymous,
			DisplayImagePath:   a.DisplayImagePath,
			CompletionFilePath: a.CompletionFilePath,
			Rating:             a.Rating,
			Comment:            a.Comment,
			CreateTime:         now,
			CompleteTime:       a.CompleteTime,
			LastUpdateTime:     now,
			State:              model.ArtworkStateActive,
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
		Price: a.Price,
		DayUsed: a.DayUsed,
		Size: a.Size,
		Volume:             a.Volume,
		Resolution:         a.Resolution,
		Format:             a.Format,
		IsR18:              a.IsR18,
		Anonymous:          a.Anonymous,
		DisplayImagePath:   a.DisplayImagePath,
		CompletionFilePath: a.CompletionFilePath,
		Rating:             a.Rating,
		Comment:            a.Comment,
		CreateTime:         a.CreateTime,
		CompleteTime:       a.CompleteTime,
		LastUpdateTime:     a.LastUpdateTime,
		State:              a.State,
	}
}