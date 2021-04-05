package msg_broker

import (
	"context"
	"pixstall-artwork/domain/artwork/model"
)

type Repo interface {
	SendArtworkCreatedMessage(ctx context.Context, artwork model.Artwork) error
	SendArtworkUpdatedMessage(ctx context.Context, updater model.ArtworkUpdater, views *int, favorCount *int) error
}
