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

	candleJobs := make(chan models.Candle, config.CANDLE_CHANNEL_BUFFER_COUNT)

	db, err := db.Connect(config.Settings.DatabaseDSN)

	if err != nil {
		panic("Failed to connect to the database")
	} else {
		log.Printf("connected to the database...")
	}

	// Create candle table if not exist.
	db.AutoMigrate(&models.Candle{})

	// Start candle workers
	workers.StartCandleWorkers(db, candleJobs, &wg)

	// Fetching candles every two minutes
	var (
		symbol   = "BTCUSDT"
		interval = "1m"
		limit    = 1000
	)

	go fetchers.StartCandleJob(symbol, interval, limit, candleJobs, db)

	// Wait for all workers to finish until the end.
	wg.Wait()
}
