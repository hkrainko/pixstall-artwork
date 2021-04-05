package update_artwork

type Response struct {
	ArtworkID string `json:"artworkId"`
}

func NewResponse(artworkID string) *Response {
	return &Response{
		ArtworkID: artworkID,
	}
}