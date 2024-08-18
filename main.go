package main

import (
	"log"
	"sync"

	"github.com/mjthecoder65/crypto-market-snapshots/config"
	"github.com/mjthecoder65/crypto-market-snapshots/db"
	"github.com/mjthecoder65/crypto-market-snapshots/fetchers"
	"github.com/mjthecoder65/crypto-market-snapshots/models"
	"github.com/mjthecoder65/crypto-market-snapshots/workers"
)

func main() {
	var wg sync.WaitGroup

	candleChannel := make(chan models.Candle, config.CANDLE_CHANNEL_BUFFER_COUNT)
	orderBookChannel := make(chan models.OrderBook, config.ORDER_BOOK_CHANNEL_BUFFER_COUNT)

	db, err := db.Connect(config.Settings.DatabaseDSN)

	if err != nil {
		panic("Failed to connect to the database")
	} else {
		log.Printf("connected to the databaser...")
	}

	// Create candle table if not exist.
	db.AutoMigrate(&models.Candle{})

	// Start candle workers
	workers.StartCandleWorkers(db, candleChannel, &wg)

	// Start order book workers (example with 1 worker, adjust as needed)
	workers.StartOrderBookWorkers(db, orderBookChannel, &wg)

	// Fetching candles every two minutes
	go fetchers.FetchCandleEveryTwoMinutes("BTCUSDT", "1m", 1000, candleChannel)

	// Wait for all workers to finish
	wg.Wait()
}
