package fetchers

import (
	"context"
	"strconv"
	"time"

	"github.com/mjthecoder65/crypto-market-snapshots/models"
)

func FetchCandles(symbol string, interval string, limit int, candleChannel chan<- models.Candle) {
	klineService := client.NewKlinesService().Symbol(symbol).Interval(interval).Limit(limit)
	klines, err := klineService.Do(context.Background())

	if err != nil {
		return
	}

	for _, kline := range klines {
		open, _ := strconv.ParseFloat(kline.Open, 64)
		close, _ := strconv.ParseFloat(kline.Close, 64)
		high, _ := strconv.ParseFloat(kline.High, 64)
		low, _ := strconv.ParseFloat(kline.Low, 64)
		volume, _ := strconv.ParseFloat(kline.Volume, 64)

		candle := models.Candle{
			Symbol:    symbol,
			Interval:  interval,
			OpenTime:  kline.OpenTime,
			CloseTime: kline.CloseTime,
			Open:      open,
			Close:     close,
			High:      high,
			Low:       low,
			Volume:    volume,
		}
		candleChannel <- candle
	}
}

func FetchCandleEveryTwoMinutes(symbol string, interval string, limit int, candleChannel chan models.Candle) {
	ticker := time.NewTicker(time.Minute / 5)
	defer ticker.Stop()
	for range ticker.C {
		FetchCandles(symbol, interval, limit, candleChannel)
	}
}
