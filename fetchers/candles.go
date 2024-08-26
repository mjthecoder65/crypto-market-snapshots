package fetchers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/mjthecoder65/crypto-market-snapshots/models"
	"gorm.io/gorm"
)

func FetchCandles(symbol string, interval string, limit int, db *gorm.DB) ([]models.Candle, error) {
	latestCandle, err := models.GetLatestCandle(symbol, interval, db)

	startTime := time.Now().Add(-1 * time.Hour)

	if err != nil {
		startTime = latestCandle.OpenTime
	}

	klineService := client.NewKlinesService().Symbol(symbol).Interval(interval).StartTime(
		startTime.UnixNano() / int64(time.Millisecond)).Limit(limit)
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
			OpenTime:  time.Unix(0, kline.OpenTime*int64(time.Millisecond)),
			CloseTime: time.Unix(0, kline.CloseTime*int64(time.Millisecond)),
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

func StartCandleJob(symbol string, interval string, limit int, candleChannel chan models.Candle, db *gorm.DB) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		candles, err := FetchCandles(symbol, interval, limit, db)

		if err != nil {
			log.Println(err)
			continue
		}

		for _, candle := range candles {
			SendCandleMessage(candle, candleChannel)
		}
	}
}
