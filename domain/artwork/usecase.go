package artwork

import (
	"context"
	"pixstall-artwork/domain/artwork/model"
)

type UseCase interface {
	AddArtwork(ctx context.Context, creator model.ArtworkCreator) (*model.Artwork, error)
	GetArtwork(ctx context.Context, artworkID string) (*model.Artwork, error)
	GetArtworks(ctx context.Context, filter model.ArtworkFilter, sorter model.ArtworkSorter) (*[]model.Artwork, error)
	UpdateArtwork(ctx context.Context, requesterID string, artworkUpdater model.ArtworkUpdater) error
}