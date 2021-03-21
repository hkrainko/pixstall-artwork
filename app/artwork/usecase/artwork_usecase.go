package usecase

import (
	"context"
	"pixstall-artwork/app/commission/delivery/rabbitmq/msg"
	"pixstall-artwork/domain/artwork"
	"pixstall-artwork/domain/artwork/model"
)

type artworkUseCase struct {
	artworkRepo artwork.Repo
}

func NewArtworkUseCase(artworkRepo artwork.Repo) artwork.UseCase {
	return &artworkUseCase{
		artworkRepo: artworkRepo,
	}
}

func (a artworkUseCase) AddFromCompletedCommission(ctx context.Context, completedCommission msg.CompletedCommission) error {


	return a.artworkRepo.
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

