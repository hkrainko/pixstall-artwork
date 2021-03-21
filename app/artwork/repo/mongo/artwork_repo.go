package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pixstall-artwork/app/artwork/repo/mongo/dao"
	"pixstall-artwork/domain/artwork"
	"pixstall-artwork/domain/artwork/model"
	error2 "pixstall-artwork/domain/error"
)

type mongoArtworkRepo struct {
	db         *mongo.Database
	collection *mongo.Collection
}

const (
	ArtworkCollection = "Artworks"
)

func NewMongoArtworkRepo(db *mongo.Database) artwork.Repo {
	return &mongoArtworkRepo{
		db:         db,
		collection: db.Collection(ArtworkCollection),
	}
}

func (m mongoArtworkRepo) AddArtwork(ctx context.Context, creator model.ArtworkCreator) (*model.Artwork, error) {
	newArtwork := dao.NewFromArtworkCreator(creator)
	result, err := m.collection.InsertOne(ctx, newArtwork)
	if err != nil {
		fmt.Printf("AddArtwork error %v\n", err)
		return nil, err
	}
	fmt.Printf("AddArtwork %v", result.InsertedID)
	dArtwork := newArtwork.ToDomainArtwork()
	return &dArtwork, nil
}

func (m mongoArtworkRepo) GetArtwork(ctx context.Context, artworkID string) (*model.Artwork, error) {
	daoArtwork := dao.Artwork{}
	err := m.collection.FindOne(ctx, bson.M{"id": artworkID}, nil).Decode(&daoArtwork)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, error2.NotFoundError
		default:
			return nil, error2.UnknownError
		}
	}
	dArtwork := daoArtwork.ToDomainArtwork()
	return &dArtwork, nil
}

func (m mongoArtworkRepo) GetArtworks(ctx context.Context, filter model.ArtworkFilter, sorter model.ArtworkSorter) (*[]model.Artwork, error) {
	daoFilter := dao.NewFilterFormDomainArtworkFilter(filter)
	opts := options.FindOptions{}
	os := int64(filter.Offset)
	opts.Skip = &os
	c := int64(filter.Count)
	opts.Limit = &c

	cursor, err := m.collection.Find(ctx, daoFilter, &opts)
	if err != nil {
		return nil, error2.UnknownError
	}
	defer cursor.Close(ctx)
	var dArtwork []model.Artwork
	for cursor.Next(ctx) {
		var r dao.Artwork
		if err := cursor.Decode(&r); err != nil {
			dArtwork = append(dArtwork, r.ToDomainArtwork())
		}
	}
	return &dArtwork, nil
}

func (m mongoArtworkRepo) UpdaterArtwork(ctx context.Context, artworkUpdater model.ArtworkUpdater) error {
	filter := bson.M{
		"id": artworkUpdater.ID,
	}
	updater := dao.NewUpdaterFromArtworkUpdater(artworkUpdater)
	_, err := m.collection.UpdateOne(ctx, filter, updater)
	if err != nil {
		return error2.UnknownError
	}
	return nil
}