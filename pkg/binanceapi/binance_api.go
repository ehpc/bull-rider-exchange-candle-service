package binanceapi

import (
	"strconv"
	"encoding/json"
	"fmt"
	"reflect"
	
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/api"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candle"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/transport"
	myReflect "github.com/ehpc/bull-rider-exchange-candle-service/pkg/reflect"
)

// BinanceAPI is API for binance.com
type BinanceAPI struct {
	restTransport transport.Transport
	streamTransport transport.Transport
}

// NewAPI creates new Binance API instance
func NewAPI(restTransport, streamTransport transport.Transport) api.API {
	return &BinanceAPI{
		restTransport: restTransport,
		streamTransport: streamTransport,
	}
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
			messageChannel, errorChannel := api.restTransport.Receive(
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
			var candlesJSON []CandleRESTJSON
			if err := json.Unmarshal(message.Body, &candlesJSON); err != nil {
				return nil, err
			}
			for _, candleJSON := range candlesJSON {
				cndl, err := CandleRESTJSONtoCandle(
					candle.ExchangeBinance,
					result.Pair,
					result.Interval,
					candleJSON,
				)
				if err != nil {
					return nil, err
				}
				candles = append(candles, cndl)
			}
		case err := <-result.ErrorChannel:
			return nil, err
		}
	}
	return candles, nil
}

// WaitForCandles asynchronously fetches candles from Binance streams
func (api *BinanceAPI) WaitForCandles(pairs []candle.Pair, intervals []candle.Interval) (chan candle.Candle, chan error) {
	candleChannel := make(chan candle.Candle)
	errorChannel := make(chan error)

	go func(){
		messageChannels := []chan transport.Message{}
		errorChannels := []chan error{}
		// For each combination start fetching candles
		for _, pair := range pairs {
			for _, interval := range intervals {
				messageChannel, errorChannel := api.streamTransport.Receive(
					transport.RequestParams{
						"WebsocketPath": fmt.Sprintf("/ws/%s@kline_%s", pair, interval),
						"symbol": string(pair),
						"interval": string(interval),
					},
				)
				messageChannels = append(messageChannels, messageChannel)
				errorChannels = append(errorChannels, errorChannel)
			}
		}
	
		// Generating channel select logic
		selectCases := make([]reflect.SelectCase, len(messageChannels) + len(errorChannels))
		for i, ch := range messageChannels {
			selectCases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
		}
		for i, ch := range errorChannels {
			selectCases[len(messageChannels) + i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
		}
		// Pull data from all open channels
		remaining := len(selectCases)
		for remaining > 0 {
			chosen, value, ok := reflect.Select(selectCases)
			// If channel is closed
			if !ok {
				selectCases[chosen].Chan = reflect.ValueOf(nil)
				remaining--
				continue
			}
			// If error occurred
			if myReflect.IsValueAnError(value) {
				errorChannel <- myReflect.ValueToError(value)
				continue
			}
			// We received a message
			message := value.Interface().(transport.Message)
			var candleJSON CandleWebsocketJSON
			if err := json.Unmarshal(message.Body, &candleJSON); err != nil {
				errorChannel <- err
				continue
			}
			cndl, err := CandleWebsocketJSONtoCandle(
				candle.ExchangeBinance,
				candleJSON,
			)
			if err != nil {
				errorChannel <- err
				continue
			}
			candleChannel <- cndl
		}
		close(candleChannel)
		close(errorChannel)
	}()

	return candleChannel, errorChannel
}

// CandleRESTJSONtoCandle converts Binance REST JSON candle to Candle
func CandleRESTJSONtoCandle(exchange candle.Exchange, pair candle.Pair, interval candle.Interval, candleJSON CandleRESTJSON) (candle.Candle, error) {
	open, err := strconv.ParseFloat(candleJSON.Open, 64)
	if err != nil {
		return candle.Candle{}, err
	}
	close, err := strconv.ParseFloat(candleJSON.Close, 64)
	if err != nil {
		return candle.Candle{}, err
	}
	high, err := strconv.ParseFloat(candleJSON.High, 64)
	if err != nil {
		return candle.Candle{}, err
	}
	low, err := strconv.ParseFloat(candleJSON.Low, 64)
	if err != nil {
		return candle.Candle{}, err
	}
	volume, err := strconv.ParseFloat(candleJSON.Volume, 64)
	if err != nil {
		return candle.Candle{}, err
	}
	quoteVolume, err := strconv.ParseFloat(candleJSON.QuoteVolume, 64)
	if err != nil {
		return candle.Candle{}, err
	}
	candle := candle.NewCandle(
		exchange,
		pair,
		interval,
		candleJSON.OpenTime,
		candleJSON.CloseTime,
		open,
		close,
		high,
		low,
		volume,
		quoteVolume,
		candleJSON.TradesCount,
	)
	return candle, nil
}

// CandleWebsocketJSONtoCandle converts Binance Websocket JSON candle to Candle
func CandleWebsocketJSONtoCandle(exchange candle.Exchange, candleJSON CandleWebsocketJSON) (candle.Candle, error) {
	open, err := strconv.ParseFloat(candleJSON.Kline.Open, 64)
	if err != nil {
		return candle.Candle{}, err
	}
	close, err := strconv.ParseFloat(candleJSON.Kline.Close, 64)
	if err != nil {
		return candle.Candle{}, err
	}
	high, err := strconv.ParseFloat(candleJSON.Kline.High, 64)
	if err != nil {
		return candle.Candle{}, err
	}
	low, err := strconv.ParseFloat(candleJSON.Kline.Low, 64)
	if err != nil {
		return candle.Candle{}, err
	}
	volume, err := strconv.ParseFloat(candleJSON.Kline.Volume, 64)
	if err != nil {
		return candle.Candle{}, err
	}
	quoteVolume, err := strconv.ParseFloat(candleJSON.Kline.QuoteVolume, 64)
	if err != nil {
		return candle.Candle{}, err
	}
	candle := candle.NewCandle(
		exchange,
		candle.Pair(candleJSON.Kline.Pair),
		candle.Interval(candleJSON.Kline.Interval),
		candleJSON.Kline.OpenTime,
		candleJSON.Kline.CloseTime,
		open,
		close,
		high,
		low,
		volume,
		quoteVolume,
		candleJSON.Kline.TradesCount,
	)
	return candle, nil
}
