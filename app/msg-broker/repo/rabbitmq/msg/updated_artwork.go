package msg

import (
	"pixstall-artwork/domain/artwork/model"
	"time"
)

type UpdatedArtwork struct {
	ID                   string  `json:"id"`
	ArtistName           *string `json:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath"`
	RequesterName        *string `json:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath"`

	Title       *string `json:"title"`
	TextContent *string `json:"textContent"`
	Views       *int    `json:"views"`
	FavorCount  *int    `json:"favorCount"`

	LastUpdateTime *time.Time          `json:"lastUpdateTime"`
	State          *model.ArtworkState `json:"state"`
}

func NewUpdatedArtwork(updater model.ArtworkUpdater, views *int, favorCount *int) UpdatedArtwork {
	return UpdatedArtwork{
		ID:                   updater.ID,
		ArtistName:           updater.ArtistName,
		ArtistProfilePath:    updater.ArtistProfilePath,
		RequesterName:        updater.RequesterName,
		RequesterProfilePath: updater.RequesterProfilePath,
		Title:                updater.Title,
		TextContent:          updater.TextContent,
		Views:                views,
		FavorCount:           favorCount,
		LastUpdateTime:       updater.LastUpdateTime,
		State:                updater.State,
	}

}
