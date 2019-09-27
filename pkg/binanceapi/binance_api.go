package binanceapi

import (
	"strconv"
	"encoding/json"
	"fmt"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candle"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/transport"
)

const (
	apiURL = "https://api.binance.com"
	makeCandlesCount = 100
)

// BinanceAPI is API for binance.com
type BinanceAPI struct {
	URL       string
	Transport transport.Transport
}

// NewBinanceAPI creates new API instance
func NewBinanceAPI(transport transport.Transport) BinanceAPI {
	api := BinanceAPI{
		URL:       apiURL,
		Transport: transport,
	}
	return api
}

// GetCandles fetches past candle data from Binance
func (api *BinanceAPI) GetCandles(pairs []candle.Pair, intervals []candle.Interval) ([]candle.Candle, error) {
	type ResultsElement struct {
		Pair candle.Pair
		Interval candle.Interval
		Channel chan transport.Message
	}
	results := []ResultsElement{}
	// For each combination fetch candles
	for _, pair := range pairs {
		for _, interval := range intervals {
			results = append(
				results,
				ResultsElement{
					Pair: pair,
					Interval: interval,
					Channel: api.Transport.Receive(
						GetCandlesRequestParams{
							Symbol: string(pair),
							Interval: string(interval),
						},
					),
				},
			)
		}
	}
	candles := make([]candle.Candle, makeCandlesCount)
	// Wait for all candles and aggregate results
	for _, result := range(results) {
		message := <-result.Channel
		candlesJSON := make([]CandleJSON, makeCandlesCount)
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
	}
	return candles, nil
}

// GetCandlesRequestParams are request params for /klines
type GetCandlesRequestParams struct {
	Symbol string
	Interval string
}

// Hash returns hash representation of GetCandlesRequestParams
func (p GetCandlesRequestParams) Hash() string {
	return fmt.Sprintf("%s#%s", p.Symbol, p.Interval)
}

// Map returns map representation of GetCandlesRequestParams
func (p GetCandlesRequestParams) Map() map[string]string {
	return map[string]string{
		"symbol": p.Symbol,
		"interval": p.Interval,
	}
}
