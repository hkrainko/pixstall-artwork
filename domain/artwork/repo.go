package artwork

import (
	"context"
	"pixstall-artwork/domain/artwork/model"
	model2 "pixstall-artwork/domain/user/model"
)

type Repo interface {
	AddArtwork(ctx context.Context, creator model.ArtworkCreator) (*model.Artwork, error)
	GetArtwork(ctx context.Context, artworkID string) (*model.Artwork, error)
	GetArtworks(ctx context.Context, filter model.ArtworkFilter, sorter model.ArtworkSorter) (*model.GetArtworksResult, error)
	UpdaterArtwork(ctx context.Context, artworkUpdater model.ArtworkUpdater) error
	UpdaterArtworkUser(ctx context.Context, updater model2.UserUpdater) error
}