package binanceapi

import (
	"fmt"
	"testing"
	"encoding/json"

	"github.com/stretchr/testify/assert"

	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candle"
	myTesting "github.com/ehpc/bull-rider-exchange-candle-service/pkg/testing"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/transport"
)

func TestCandlesJSON(t *testing.T) {
	t.Run("REST candle", func (t *testing.T) {
		jsonData := fmt.Sprintf("[%s]", myTesting.BinanceCandleExampleRESTJSON)
		candles := make([]CandleRESTJSON, 1)
		err := json.Unmarshal([]byte(jsonData), &candles)
		assert.NoError(t, err)
		assert.NotEmpty(t, candles)
		assert.Equal(t, int64(1561622400000), candles[0].OpenTime)
		assert.Equal(t, "0.42680000", candles[0].High)
	})
	t.Run("Websocket candle", func (t *testing.T) {
		jsonData := fmt.Sprintf("%s", myTesting.BinanceCandleExampleWebsocketJSON)
		candle := CandleWebsocketJSON{}
		err := json.Unmarshal([]byte(jsonData), &candle)
		assert.NoError(t, err)
		assert.Equal(t, "kline", candle.EventType)
		assert.Equal(t, int64(123400000), candle.Kline.OpenTime)
		assert.Equal(t, "0.0025", candle.Kline.High)
	})
}

func TestRESTAPI(t *testing.T) {
	const testCandlesCount = 10

	// Populating fake data
	apiTransport := myTesting.NewTransportMock()
	apiTransport.AddReceivableMessage(
		transport.Message{
			Body: []byte(
				myTesting.GenerateCandlesJSON(myTesting.BinanceCandleExampleRESTJSON, testCandlesCount),
			),
		},
		transport.RequestParams{
			"HTTPMethod": "GET",
			"HTTPPath": "/api/v1/klines",
			"symbol": candle.PairIOTAUSDT,
			"interval": candle.Interval15m,
		},
	)
	apiTransport.AddReceivableMessage(
		transport.Message{
			Body: []byte(
				myTesting.GenerateCandlesJSON(myTesting.BinanceCandleExampleRESTJSON, testCandlesCount),
			),
		},
		transport.RequestParams{
			"HTTPMethod": "GET",
			"HTTPPath": "/api/v1/klines",
			"symbol": candle.PairIOTAUSDT,
			"interval": candle.Interval1h,
		},
	)
	apiTransport.AddReceivableMessage(
		transport.Message{
			Body: []byte(
				myTesting.GenerateCandlesJSON(myTesting.BinanceCandleExampleRESTJSON, testCandlesCount),
			),
		},
		transport.RequestParams{
			"HTTPMethod": "GET",
			"HTTPPath": "/api/v1/klines",
			"symbol": candle.PairETHUSDT,
			"interval": candle.Interval15m,
		},
	)
	apiTransport.AddReceivableMessage(
		transport.Message{
			Body: []byte(
				myTesting.GenerateCandlesJSON(myTesting.BinanceCandleExampleRESTJSON, testCandlesCount),
			),
		},
		transport.RequestParams{
			"HTTPMethod": "GET",
			"HTTPPath": "/api/v1/klines",
			"symbol": candle.PairETHUSDT,
			"interval": candle.Interval1h,
		},
	)

	// Getting candles
	api := NewAPI(apiTransport, nil)
	pairs := []candle.Pair{candle.PairIOTAUSDT, candle.PairETHUSDT}
	intervals := []candle.Interval{candle.Interval15m, candle.Interval1h}
	candles, err := api.GetCandles(pairs, intervals)
	assert.NoError(t, err)

	// Checking that all pairs are present
	iotaCandles := []candle.Candle{}
	ethCandles := []candle.Candle{}
	for _, x := range candles {
		switch x.Pair {
		case candle.PairIOTAUSDT:
			iotaCandles = append(iotaCandles, x)
		case candle.PairETHUSDT:
			ethCandles = append(ethCandles, x)
		}
	}
	assert.NotEmpty(t, iotaCandles)
	assert.NotEmpty(t, ethCandles)
	
	// Checking that all intervals are present
	interval15mCandles := []candle.Candle{}
	interval1hCandles := []candle.Candle{}
	for _, x := range iotaCandles {
		switch x.Interval {
		case candle.Interval15m:
			interval15mCandles = append(interval15mCandles, x)
		case candle.Interval1h:
			interval1hCandles = append(interval1hCandles, x)
		}
	}
	assert.NotEmpty(t, interval15mCandles)
	assert.NotEmpty(t, interval1hCandles)
	interval15mCandles = nil
	interval1hCandles = nil
	for _, x := range ethCandles {
		switch x.Interval {
		case candle.Interval15m:
			interval15mCandles = append(interval15mCandles, x)
		case candle.Interval1h:
			interval1hCandles = append(interval1hCandles, x)
		}
	}
	assert.NotEmpty(t, interval15mCandles)
	assert.NotEmpty(t, interval1hCandles)

	// Checking inner values
	assert.Equal(t, int64(1561622400000), candles[0].OpenTime)
	assert.Equal(t, 0.42680000, candles[0].High)
}
