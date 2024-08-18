package workers

import (
	"sync"

	"github.com/mjthecoder65/crypto-market-snapshots/models"
	"github.com/mjthecoder65/crypto-market-snapshots/processors"
	"gorm.io/gorm"
)

func StartCandleWorkers(db *gorm.DB, candleChannel chan models.Candle, wg *sync.WaitGroup) {
	for i := 0; i < NUMBER_OF_WORKERS; i++ {
		wg.Add(1)
		go processors.ProcessCandles(db, candleChannel, wg)
	}
}

func StartOrderBookWorkers(db *gorm.DB, orderBookChannel chan models.OrderBook, wg *sync.WaitGroup) {
	for i := 0; i < NUMBER_OF_WORKERS; i++ {
		wg.Add(1)
		go processors.ProcessOrderBook(db, orderBookChannel, wg)
	}
}
