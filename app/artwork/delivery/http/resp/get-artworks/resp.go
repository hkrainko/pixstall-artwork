package get_artworks

import (
	"pixstall-artwork/app/artwork/delivery/http/resp/get-artwork"
	"pixstall-artwork/domain/artwork/model"
)

type Response struct {
	Artworks []get_artwork.Artwork `json:"artworks"`
	Offset   int       `json:"offSet"`
	Count    int       `json:"count"`
}

func NewResponse(artworks []model.Artwork, userID *string, offset int, count int) *Response {
	var rArtworks []get_artwork.Artwork
	for _, a := range artworks {
		rArtworks = append(rArtworks, get_artwork.NewRespArtworkFormDomainArtwork(a, userID))
	}
	return &Response{
		Artworks: rArtworks,
		Offset: offset,
		Count: count,
	}
}
