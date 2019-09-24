package binanceapi

import (
	"ehpc.io/bull-rider/exchange-candle-service/pkg/candle"
	"ehpc.io/bull-rider/exchange-candle-service/pkg/transport"
)

const apiURL = "https://api.binance.com"

//BinanceAPI is API for binance.com
type BinanceAPI struct {
	URL       string
	Transport transport.Transport
}

//NewBinanceAPI creates new API instance
func NewBinanceAPI(transport transport.Transport) BinanceAPI {
	api := BinanceAPI{
		URL:       apiURL,
		Transport: transport,
	}
	return api
}

//GetCandles fetches past candle data from Binance
func (api *BinanceAPI) GetCandles(pairs []candle.Pair, intervals []candle.Interval) []candle.Candle {
	return []candle.Candle{}
}
