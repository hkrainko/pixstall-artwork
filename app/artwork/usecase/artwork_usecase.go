package usecase

import (
	"context"
	"pixstall-artwork/domain/artwork"
	"pixstall-artwork/domain/artwork/model"
	msgBroker "pixstall-artwork/domain/msg-broker"
)

type artworkUseCase struct {
	artworkRepo artwork.Repo
	msgBrokerRepo msgBroker.Repo
}

func NewArtworkUseCase(artworkRepo artwork.Repo, msgBrokerRepo msgBroker.Repo) artwork.UseCase {
	return &artworkUseCase{
		artworkRepo: artworkRepo,
		msgBrokerRepo: msgBrokerRepo,
	}
}

func (a artworkUseCase) AddArtwork(ctx context.Context, creator model.ArtworkCreator) (*model.Artwork, error) {
	return a.artworkRepo.AddArtwork(ctx, creator)
}

func (a artworkUseCase) GetArtwork(ctx context.Context, artworkID string) (*model.Artwork, error) {
	return a.artworkRepo.GetArtwork(ctx, artworkID)
}

func (a artworkUseCase) GetArtworks(ctx context.Context, filter model.ArtworkFilter, sorter model.ArtworkSorter) (*[]model.Artwork, error) {
	return a.artworkRepo.GetArtworks(ctx, filter, sorter)
}

func (a artworkUseCase) UpdaterArtwork(ctx context.Context, artworkUpdater model.ArtworkUpdater) error {
	return a.artworkRepo.UpdaterArtwork(ctx, artworkUpdater)
}

