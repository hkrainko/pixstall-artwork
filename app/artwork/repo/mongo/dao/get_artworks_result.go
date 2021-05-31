package dao

import "pixstall-artwork/domain/artwork/model"

type GetArtworksResult struct {
	Artworks []Artwork `bson:"artworks"`
	Total    int       `bson:"total"`
}

func (g GetArtworksResult) ToDomainGetArtworksResult() *model.GetArtworksResult {
	dArtworks := make([]model.Artwork, 0)
	for _, artwork := range g.Artworks {
		dArtworks = append(dArtworks, artwork.ToDomainArtwork())
	}

	return &model.GetArtworksResult{
		Artwork: dArtworks,
		Total:   g.Total,
	}
}