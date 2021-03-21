package model

type Price struct {
	Amount   float64  `json:"amount" bson:"amount"`
	Currency Currency `json:"currency" bson:"currency"`
}

type Currency string

const (
	CurrencyHKD Currency = "HKD"
	CurrencyTWD Currency = "TWD"
	CurrencyUSE Currency = "USD"
)