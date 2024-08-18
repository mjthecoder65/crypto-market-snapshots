package fetchers

import (
	"github.com/adshao/go-binance/v2"
	"github.com/mjthecoder65/crypto-market-snapshots/config"
)

var client *binance.Client = binance.NewClient(config.Settings.BinanceAPIKey, config.Settings.BinanceSecretKey)
