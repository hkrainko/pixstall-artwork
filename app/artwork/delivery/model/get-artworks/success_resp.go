package get_artworks

import "pixstall-artwork/domain/artwork/model"

type Response struct {
	Artworks []model.Artwork `json:"artworks"`
	Offset   int             `json:"offSet"`
	Count    int             `json:"count"`
}

func NewResponse(artworks []model.Artwork, offset int, count int) *Response {
	return &Response{
		Artworks: artworks,
		Offset: offset,
		Count: count,
	}
}
