package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"pixstall-artwork/domain/artwork/model"
	"time"
)

func NewUpdaterFromArtworkUpdater(a model.ArtworkUpdater) bson.D {
	updater := bson.D{}

	setter := bson.D{}
	if a.ArtistName != nil {
		setter = append(setter, bson.E{Key: "artistName", Value: a.ArtistName})
	}
	if a.ArtistProfilePath != nil {
		setter = append(setter, bson.E{Key: "artistProfilePath", Value: a.ArtistProfilePath})
	}
	if a.RequesterName != nil {
		setter = append(setter, bson.E{Key: "requesterName", Value: a.RequesterName})
	}
	if a.RequesterProfilePath != nil {
		setter = append(setter, bson.E{Key: "requesterProfilePath", Value: a.RequesterProfilePath})
	}
	if a.State != nil {
		setter = append(setter, bson.E{Key: "state", Value: a.State})
	}
	if a.Title != nil {
		setter = append(setter, bson.E{Key: "title", Value: a.Title})
	}
	if a.TextContent != nil {
		setter = append(setter, bson.E{Key: "textContent", Value: a.TextContent})
	}
	if a.Favor != nil {
		setter = append(setter, bson.E{Key: "favors." + *a.Favor, Value: true})
	}
	setter = append(setter, bson.E{Key: "lastUpdateTime", Value: time.Now()})

	unsetter := bson.D{}
	if a.Unfavor != nil {
		unsetter = append(unsetter, bson.E{Key: "favors." + *a.Unfavor, Value: ""})
	}

	putter := bson.D{}

	if len(setter) > 0 {
		updater = append(updater, bson.E{Key: "$set", Value: setter})
	}
	if len(unsetter) > 0 {
		updater = append(updater, bson.E{Key: "$unset", Value: unsetter})
	}
	if len(putter) > 0 {
		updater = append(updater, bson.E{Key: "$push", Value: putter})
	}

	return updater
}
