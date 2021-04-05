package usecase

import (
	"context"
	"log"
	"pixstall-artwork/domain/artwork"
	"pixstall-artwork/domain/artwork/model"
	error2 "pixstall-artwork/domain/error"
	msgBroker "pixstall-artwork/domain/msg-broker"
)

type artworkUseCase struct {
	artworkRepo   artwork.Repo
	msgBrokerRepo msgBroker.Repo
}

func NewArtworkUseCase(artworkRepo artwork.Repo, msgBrokerRepo msgBroker.Repo) artwork.UseCase {
	return &artworkUseCase{
		artworkRepo:   artworkRepo,
		msgBrokerRepo: msgBrokerRepo,
	}
}

func (a artworkUseCase) AddArtwork(ctx context.Context, creator model.ArtworkCreator) (*model.Artwork, error) {
	dArtwork, err := a.artworkRepo.AddArtwork(ctx, creator)
	if err != nil {
		return nil, err
	}
	err = a.msgBrokerRepo.SendArtworkCreatedMessage(ctx, *dArtwork)
	if err != nil {
		log.Println(err)
		// ignore error
	}
	return dArtwork, nil
}

func (a artworkUseCase) GetArtwork(ctx context.Context, artworkID string) (*model.Artwork, error) {
	return a.artworkRepo.GetArtwork(ctx, artworkID)
}

func (a artworkUseCase) GetArtworks(ctx context.Context, filter model.ArtworkFilter, sorter model.ArtworkSorter) (*[]model.Artwork, error) {
	return a.artworkRepo.GetArtworks(ctx, filter, sorter)
}

func (a artworkUseCase) UpdateArtwork(ctx context.Context, requesterID string, artworkUpdater model.ArtworkUpdater) error {
	dArtwork, err := a.artworkRepo.GetArtwork(ctx, artworkUpdater.ID)
	if err != nil {
		return err
	}
	if requesterID != dArtwork.ArtistID {
		return error2.UnAuthError
	}
	err = a.artworkRepo.UpdaterArtwork(ctx, artworkUpdater)
	if err != nil {
		return err
	}
	err = a.msgBrokerRepo.SendArtworkUpdatedMessage(ctx, artworkUpdater, nil, nil)
	if err != nil {
		log.Println(err)
		// Ignore error
	}
	return nil
}

func (a artworkUseCase) UpdateArtworkFavor(ctx context.Context, requesterID string, artworkID string, favor bool) error {
	dArtwork, err := a.artworkRepo.GetArtwork(ctx, artworkID)
	if err != nil {
		return err
	}
	if !a.isAllowUpdateFavor(dArtwork, requesterID, favor) {
		return error2.ForbiddenError
	}
	updater := model.ArtworkUpdater{
		ID: artworkID,
	}
	newFavorCount := len(dArtwork.Favors)
	if favor {
		updater.Favor = &requesterID
		newFavorCount++
	} else {
		updater.Unfavor = &requesterID
		newFavorCount--
	}
	err = a.artworkRepo.UpdaterArtwork(ctx, updater)
	if err != nil {
		return err
	}
	err = a.msgBrokerRepo.SendArtworkUpdatedMessage(ctx, updater, nil, &newFavorCount)
	if err != nil {
		log.Println(err)
		// Ignore error
	}
	return nil
}

func (a artworkUseCase) isAllowUpdateFavor(dArtwork *model.Artwork, requesterID string, favor bool) bool {
	_, favored := dArtwork.Favors[requesterID]
	if favor && favored {
		return false
	} else if !favor && !favored {
		return false
	}
	return true
}
