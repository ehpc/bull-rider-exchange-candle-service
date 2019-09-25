package binanceapi

import (
	"fmt"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candle"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/transport"
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
	results := [](chan transport.Message){}
	// For each combination fetch candles
	for _, pair := range pairs {
		for _, interval := range intervals {
			results = append(
				results,
				api.Transport.Receive(
					GetCandlesRequestParams{
						Symbol: string(pair),
						Interval: string(interval),
					},
				),
			)
		}
	}
	// Wait for all candles and aggregate results
	for _, ch := range(results) {
		result := <-ch
		fmt.Println(result)
	}
	return []candle.Candle{}
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
