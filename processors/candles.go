package processors

import (
	"log"
	"sync"

	"github.com/mjthecoder65/crypto-market-snapshots/models"
	"gorm.io/gorm"
)

func SaveCandle(candle models.Candle, db *gorm.DB) {
	result := db.Create(&candle)

	if result.Error != nil {
		log.Printf("failed to save data to the database: %v", result.Error)
	} else {
		log.Println("INFO: saved candle to the database")
	}
}

func AddOrUpdate(candle models.Candle, db *gorm.DB) {
	var existingCandle models.Candle
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

func ProcessCandles(db *gorm.DB, candleChannel <-chan models.Candle, wg *sync.WaitGroup) {
	defer wg.Done()
	for candle := range candleChannel {
		SaveCandle(candle, db)
	}
}
