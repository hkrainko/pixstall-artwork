package artwork

import (
	"context"
	"pixstall-artwork/domain/artwork/model"
	model2 "pixstall-artwork/domain/user/model"
)

type UseCase interface {
	AddArtwork(ctx context.Context, creator model.ArtworkCreator) (*model.Artwork, error)
	GetArtwork(ctx context.Context, artworkID string) (*model.Artwork, error)
	GetArtworks(ctx context.Context, filter model.ArtworkFilter, sorter model.ArtworkSorter) (*model.GetArtworksResult, error)
	UpdateArtwork(ctx context.Context, requesterID string, artworkUpdater model.ArtworkUpdater) error
	UpdateArtworkFavor(ctx context.Context, requesterID string, artworkID string, favor bool) error
	UpdateArtworkByUserUpdatedEvent(ctx context.Context, updater model2.UserUpdater) error
}
