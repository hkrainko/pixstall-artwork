package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"pixstall-artwork/app/artwork/repo/mongo/dao"
	"pixstall-artwork/domain/artwork"
	"pixstall-artwork/domain/artwork/model"
	error2 "pixstall-artwork/domain/error"
	model2 "pixstall-artwork/domain/user/model"
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
	err := m.collection.FindOne(ctx, bson.M{"id": artworkID}).Decode(&daoArtwork)
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

func (m mongoArtworkRepo) GetArtworks(ctx context.Context, filter model.ArtworkFilter, sorter model.ArtworkSorter) (*model.GetArtworksResult, error) {
	//daoFilter := dao.NewFilterFormDomainArtworkFilter(filter)

	var pipeline []bson.M
	if filter.ArtistID != nil {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"artistId": filter.ArtistID}})
	}
	if filter.RequesterID != nil {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"requesterId": filter.ArtistID}})
	}
	pipeline = append(pipeline, bson.M{
		"$facet": bson.M{
			"total": []bson.M{{
				"$count": "total",
			}},
			"artworks": bson.A{
				bson.D{{"$skip", filter.Offset}},
				bson.D{{"$limit", filter.Count}},
			},
		},
	})
	pipeline = append(pipeline, bson.M{
		"$addFields": bson.M{
			"total": bson.M{"$arrayElemAt": bson.A{"$total.total", 0}},
		},
	})

	cursor, err := m.collection.Aggregate(ctx, pipeline)
	defer cursor.Close(ctx)
	if err != nil {
		return nil, error2.UnknownError
	}
	var getArtworksResult *model.GetArtworksResult
	for cursor.Next(ctx) {
		var r dao.GetArtworksResult
		if err := cursor.Decode(&r); err != nil {
			return nil, err
		}
		getArtworksResult = r.ToDomainGetArtworksResult()
	}
	return getArtworksResult, nil
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

func (m mongoArtworkRepo) UpdaterArtworkUser(ctx context.Context, updater model2.UserUpdater) error {
	filter := bson.M{"artistId": updater.UserID}
	updateContent := bson.M{}
	if updater.UserName != nil {
		updateContent["artistName"] = updater.UserName
	}
	if updater.ProfilePath != nil {
		updateContent["artistProfilePath"] = updater.ProfilePath
	}
	update := bson.M{"$set": updateContent}

	result, err := m.collection.UpdateMany(ctx, filter, update)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	fmt.Printf("UpdateArtworkUser aritst count:%v\n", result.MatchedCount)

	filter = bson.M{"requesterId": updater.UserID}
	updateContent = bson.M{}
	if updater.UserName != nil {
		updateContent["requesterName"] = updater.UserName
	}
	if updater.ProfilePath != nil {
		updateContent["requesterProfilePath"] = updater.ProfilePath
	}
	update = bson.M{"$set": updateContent}

	result, err = m.collection.UpdateMany(ctx, filter, update)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	fmt.Printf("UpdateArtworkUser requester count:%v\n", result.MatchedCount)

	return nil
}
