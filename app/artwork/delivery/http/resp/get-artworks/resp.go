package get_artworks

import (
	"pixstall-artwork/app/artwork/delivery/http/resp/get-artwork"
	"pixstall-artwork/domain/artwork/model"
)

type Response struct {
	Artworks []get_artwork.Artwork `json:"artworks"`
	Offset   int                   `json:"offSet"`
	Count    int                   `json:"count"`
	Total    int                   `json:"total"`
}

func NewResponse(userID *string, result model.GetArtworksResult, offset int) *Response {
	var rArtworks []get_artwork.Artwork
	for _, a := range result.Artwork {
		rArtworks = append(rArtworks, get_artwork.NewRespArtworkFormDomainArtwork(a, userID))
	}
	return &Response{
		Artworks: rArtworks,
		Offset:   offset,
		Count:    len(rArtworks),
		Total:    result.Total,
	}
}
