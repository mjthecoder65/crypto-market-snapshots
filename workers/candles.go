package workers

import (
	"log"
	"sync"

	"github.com/mjthecoder65/crypto-market-snapshots/common"
	"github.com/mjthecoder65/crypto-market-snapshots/models"
	"gorm.io/gorm"
)

func CandleWorker(db *gorm.DB, candleChannel <-chan models.Candle, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for candle := range candleChannel {
		log.Printf("%s[INFO/Worker-%d] received candles : %+v %s\n", common.Green, id, candle, common.Reset)
		candle.Save(db)
	}
}

func StartCandleWorkers(db *gorm.DB, candleChannel chan models.Candle, wg *sync.WaitGroup) {
	for id := 0; id < NUMBER_OF_WORKERS; id++ {
		wg.Add(1)
		log.Printf("%s[INFO/Worker-%d] %s Started\n", common.Green, id, common.Reset)
		go CandleWorker(db, candleChannel, wg, id)
	}
}
