package model

type Size struct {
	Width  float64 `json:"width" bson:"width"`
	Height float64 `json:"height" bson:"height"`
	Unit   string  `json:"unit" bson:"unit"`
}