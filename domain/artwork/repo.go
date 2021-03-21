package artwork

import (
	"context"
	"pixstall-artwork/domain/artwork/model"
)

type Repo interface {
	AddArtwork(ctx context.Context, creator model.ArtworkCreator) (*model.Artwork, error)
	GetArtwork(ctx context.Context, artworkID string) (*model.Artwork, error)
	GetArtworks(ctx context.Context, filter model.ArtworkFilter, sorter model.ArtworkSorter) (*[]model.Artwork, error)
	UpdaterArtwork(ctx context.Context, artworkUpdater model.ArtworkUpdater) error
}