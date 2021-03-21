package model

type ArtworkFilter struct {
	ArtistID       *string
	RequesterID    *string
	Count  int
	Offset int
}
