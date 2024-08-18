package fetchers

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/mjthecoder65/crypto-market-snapshots/models"
)

func FetchOrderBooks(symbol string, client *binance.Client) models.OrderBook {
	depthService := client.NewDepthService().Symbol(symbol).Limit(1)
	depth, err := depthService.Do(context.Background())

	if err != nil {
		fmt.Printf("error: failed to fetch order books for %s\n", symbol)
	}

	var orderBook models.OrderBook
	orderBook.LastUpdatedID = depth.LastUpdateID

	for _, ask := range depth.Asks {
		orderBook.Asks = append(orderBook.Asks, models.Ask{
			Price:    ask.Price,
			Quantity: ask.Quantity,
		})
	}

	for _, bid := range depth.Bids {
		orderBook.Bids = append(orderBook.Bids, models.Bid{
			Price:    bid.Price,
			Quantity: bid.Quantity,
		})
	}

	return orderBook
}
