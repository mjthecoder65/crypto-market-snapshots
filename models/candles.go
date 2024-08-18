package models

type Candle struct {
	ID        uint    `db:"id" json:"id" gorm:"primary_key"`
	Symbol    string  `db:"symbol" json:"symbol"`
	Interval  string  `db:"symbol" json:"interval"`
	OpenTime  int64   `db:"open_time" json:"open_time"`
	CloseTime int64   `db:"close_time" json:"close_time"`
	Open      float64 `db:"open" json:"open"`
	Close     float64 `db:"close" json:"close"`
	High      float64 `db:"high" json:"high"`
	Low       float64 `db:"low" json:"low"`
	Volume    float64 `db:"volume" json:"volume"`
}
