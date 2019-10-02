package binanceapi

import (
	"strconv"
	"encoding/json"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candle"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/transport"
)

const (
	// APIURL is a base URL for API
	APIURL = "https://api.binance.com"
)

// BinanceAPI is API for binance.com
type BinanceAPI struct {
	Transport transport.Transport
}

// NewAPI creates new API instance
func NewAPI(transport transport.Transport) BinanceAPI {
	api := BinanceAPI{
		Transport: transport,
	}
	return api
}

// GetCandles fetches past candle data from Binance
func (api *BinanceAPI) GetCandles(pairs []candle.Pair, intervals []candle.Interval) ([]candle.Candle, error) {
	type ResultsElement struct {
		Pair candle.Pair
		Interval candle.Interval
		MessageChannel chan transport.Message
		ErrorChannel chan error
	}
	results := []ResultsElement{}
	// For each combination fetch candles
	for _, pair := range pairs {
		for _, interval := range intervals {
			messageChannel, errorChannel := api.Transport.Receive(
				transport.RequestParams{
					"HTTPMethod": "GET",
					"HTTPPath": "/api/v1/klines",
					"symbol": string(pair),
					"interval": string(interval),
				},
			)
			results = append(
				results,
				ResultsElement{
					Pair: pair,
					Interval: interval,
					MessageChannel: messageChannel,
					ErrorChannel: errorChannel,
				},
			)
		}
	}
	var candles []candle.Candle
	// Wait for all candles and aggregate results
	for _, result := range(results) {
		select{
		case message := <-result.MessageChannel:
			var candlesJSON []CandleJSON
			if err := json.Unmarshal(message.Body, &candlesJSON); err != nil {
				return nil, err
			}
			for _, candleJSON := range candlesJSON {
				open, err := strconv.ParseFloat(candleJSON.Open, 64)
				if err != nil {
					return nil, err
				}
				close, err := strconv.ParseFloat(candleJSON.Close, 64)
				if err != nil {
					return nil, err
				}
				high, err := strconv.ParseFloat(candleJSON.High, 64)
				if err != nil {
					return nil, err
				}
				low, err := strconv.ParseFloat(candleJSON.Low, 64)
				if err != nil {
					return nil, err
				}
				volume, err := strconv.ParseFloat(candleJSON.Volume, 64)
				if err != nil {
					return nil, err
				}
				quoteVolume, err := strconv.ParseFloat(candleJSON.QuoteVolume, 64)
				if err != nil {
					return nil, err
				}
				candles = append(
					candles,
					candle.NewCandle(
						candle.ExchangeBinance,
						result.Pair,
						result.Interval,
						candleJSON.OpenTime,
						candleJSON.CloseTime,
						open,
						close,
						high,
						low,
						volume,
						quoteVolume,
						candleJSON.TradesCount,
					),
				)
			}
		case err := <-result.ErrorChannel:
			return nil, err
		}
	}
	return candles, nil
}
