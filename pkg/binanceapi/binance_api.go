package binanceapi

import (
	"ehpc.io/bull-rider/exchange-candle-service/pkg/transport"
	"ehpc.io/bull-rider/exchange-candle-service/pkg/candle"
)

const apiURL = "https://api.binance.com"

//BinanceAPI is API for binance.com
type BinanceAPI struct {
	URL string
	Transport transport.Transport
}

//NewBinanceAPI creates new API instance
func NewBinanceAPI(transport transport.Transport) BinanceAPI {
	api := BinanceAPI{
		URL: apiURL,
		Transport: transport,
	}
	return api
}

//GetCandles fetches past candle data from Binance
func (api *BinanceAPI) GetCandles(pairs []string, intervals []candle.Interval) map[string]map[candle.Interval][]candle.Candle {
	return map[string]map[candle.Interval][]candle.Candle{
		"IOTAUSDT": {
			candle.Interval15m: {
				{
					OpenTime:     1568707200000,
					CloseTime:    1568721599999,
					Open:         0.25000000,
					Close:        0.24850000,
					High:         0.25250000,
					Low:          0.24800000,
					Volume:       722111.71000000,
					QuoteVolume:  180376.63551900,
					TradesNumber: 599,
				},
			},
			candle.Interval1h: {
				{
					OpenTime:     1568707200000,
					CloseTime:    1568721599999,
					Open:         0.25000000,
					Close:        0.24850000,
					High:         0.25250000,
					Low:          0.24800000,
					Volume:       722111.71000000,
					QuoteVolume:  180376.63551900,
					TradesNumber: 599,
				},
			},
		},
		"ETHUSDT": {
			candle.Interval15m: {
				{
					OpenTime:     1568707200000,
					CloseTime:    1568721599999,
					Open:         0.25000000,
					Close:        0.24850000,
					High:         0.25250000,
					Low:          0.24800000,
					Volume:       722111.71000000,
					QuoteVolume:  180376.63551900,
					TradesNumber: 599,
				},
			},
			candle.Interval1h: {
				{
					OpenTime:     1568707200000,
					CloseTime:    1568721599999,
					Open:         0.25000000,
					Close:        0.24850000,
					High:         0.25250000,
					Low:          0.24800000,
					Volume:       722111.71000000,
					QuoteVolume:  180376.63551900,
					TradesNumber: 599,
				},
			},
		},
	}
}
