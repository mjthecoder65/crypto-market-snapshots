package processors

import (
	"fmt"
	"sync"

	"github.com/mjthecoder65/crypto-market-snapshots/models"
	"gorm.io/gorm"
)

func ProcessOrderBook(db *gorm.DB, orderBookChannel <-chan models.OrderBook, wg *sync.WaitGroup) {
	defer wg.Done()

	for orderBook := range orderBookChannel {
		fmt.Printf("INFO: Received order book for %s ", orderBook.Symbol)
	}
}
