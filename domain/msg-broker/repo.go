package msg_broker

import (
	"context"
	"pixstall-artwork/domain/artwork/model"
)

type Repo interface {
	SendArtworkCreatedMessage(ctx context.Context, artwork model.Artwork) error
}
