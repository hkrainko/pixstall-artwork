package update_artwork_favor

type Response struct {
	ArtworkID string `json:"artworkId"`
	Favor     bool   `json:"favor"`
}

func NewResponse(artworkID string, favor bool) *Response {
	return &Response{
		ArtworkID: artworkID,
		Favor: favor,
	}
}