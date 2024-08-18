package fetchers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/mjthecoder65/crypto-market-snapshots/models"
)

func FetchCandles(symbol string, interval string, limit int) ([]models.Candle, error) {
	klineService := client.NewKlinesService().Symbol(symbol).Interval(interval).Limit(limit)
	klines, err := klineService.Do(context.Background())

	if err != nil {
		errorMessage := fmt.Sprintf("error: %v", err)
		return nil, errors.New(errorMessage)
	}

	candles := []models.Candle{}

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
		candles = append(candles, candle)
	}

	return candles, nil
}

func SendCandleMessage(candle models.Candle, sender chan<- models.Candle) {
	sender <- candle
	log.Printf("Sent candle: %v", candle)
}

func StartCandleJob(symbol string, interval string, limit int, candleChannel chan models.Candle) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		candles, err := FetchCandles(symbol, interval, limit)

		if err != nil {
			log.Printf("failed to fetch candle with params(symbol: %s interval: %s limit: %d)", symbol, interval, limit)
			continue
		}

		for _, candle := range candles {
			SendCandleMessage(candle, candleChannel)
		}
	}
}
