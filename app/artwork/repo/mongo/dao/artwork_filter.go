package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"pixstall-artwork/domain/artwork/model"
)

func NewFilterFormDomainArtworkFilter(d model.ArtworkFilter) bson.D {
	filter := bson.D{}

	if d.ArtistID != nil {
		filter = append(filter, bson.E{Key: "artistId", Value: d.ArtistID})
	}
	if d.RequesterID != nil {
		filter = append(filter, bson.E{Key: "requesterId", Value: d.RequesterID})
	}

	return filter
}