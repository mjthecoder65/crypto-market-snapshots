package models

import (
	"errors"
	"log"
	"time"

	"github.com/mjthecoder65/crypto-market-snapshots/common"
	"gorm.io/gorm"
)

type Candle struct {
	ID        uint      `db:"id" json:"id" gorm:"primary_key"`
	Symbol    string    `db:"symbol" json:"symbol" gorm:"index"`
	Interval  string    `db:"symbol" json:"interval"`
	OpenTime  time.Time `db:"open_time" json:"open_time"`
	CloseTime time.Time `db:"close_time" json:"close_time"`
	Open      float64   `db:"open" json:"open"`
	Close     float64   `db:"close" json:"close"`
	High      float64   `db:"high" json:"high"`
	Low       float64   `db:"low" json:"low"`
	Volume    float64   `db:"volume" json:"volume"`
}

func (candle *Candle) Save(db *gorm.DB) {
	result := db.Create(candle)

	if result.Error != nil {
		log.Fatalf("failed to save candle: %v\n", result.Error)
	} else {
		log.Printf("%s success: added candle: %+v %s\n", common.Yellow, candle, common.Reset)
	}
}

func (candle *Candle) AddOrUpdate(db *gorm.DB) {
	var existingCandle Candle

	result := db.Where("symbol = ? AND interval = ? AND open_time = ? AND close_time = ?",
		candle.Symbol, candle.Interval, candle.OpenTime, candle.CloseTime).
		First(&existingCandle)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Printf("Error Checking for existing candles: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		if err := db.Create(&candle).Error; err != nil {
			log.Printf("failed to create new candle: %v", err)
		} else {
			log.Println("New candle added.")
		}
	} else {
		existingCandle.Open = candle.Open
		existingCandle.High = candle.High
		existingCandle.Low = candle.Low
		existingCandle.Close = candle.Close
		existingCandle.Volume = candle.Volume
		existingCandle.CloseTime = candle.CloseTime

		if err := db.Save(&existingCandle).Error; err != nil {
			log.Printf("Error updating candle: %v", err)
		} else {
			log.Println("Candle has been updated.")
		}
	}
}

func GetLatestCandle(symbol string, interval string, db *gorm.DB) (Candle, error) {
	var candle Candle
	result := db.Where("symbol = ? AND interval= ?", symbol, interval).Order("open_time DESC").First(&candle)

	if result.Error != nil {
		return Candle{}, errors.New("failed to get the latest candle")
	}

	return candle, nil
}
