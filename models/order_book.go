package models

type Bid struct {
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

type Ask struct {
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

type OrderBook struct {
	Symbol        string `db:"symbol" json:"symbol"`
	Bids          []Bid  `db:"bids" json:"bids"`
	Asks          []Ask  `db:"asks" json:"asks"`
	LastUpdatedID int64  `db:"last_updated" json:"lastUpdatedId"`
}
