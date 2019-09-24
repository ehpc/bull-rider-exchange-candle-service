package binanceapi

import (
	"ehpc.io/bull-rider/exchange-candle-service/pkg/candle"
	myTesting "ehpc.io/bull-rider/exchange-candle-service/pkg/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCandles(t *testing.T) {
	transport := myTesting.TransportMock{}
	api := NewBinanceAPI(&transport)
	pairs := []candle.Pair{candle.IOTAUSDT, candle.ETHUSDT}
	intervals := []candle.Interval{candle.Interval15m, candle.Interval1h}
	candles := api.GetCandles(pairs, intervals)
	// Checking that all pairs are present
	iotaCandles := []candle.Candle{}
	ethCandles := []candle.Candle{}
	for _, x := range candles {
		switch x.Pair{
		case candle.IOTAUSDT:
			iotaCandles = append(iotaCandles, x)
		case candle.ETHUSDT:
			ethCandles = append(ethCandles, x)
		}
	}
	assert.NotEmpty(t, iotaCandles)
	assert.NotEmpty(t, ethCandles)
	// Checking that all intervals are present
	interval15mCandles := []candle.Candle{}
	interval1hCandles := []candle.Candle{}
	for _, x := range iotaCandles {
		switch x.Interval{
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
		switch x.Interval{
		case candle.Interval15m:
			interval15mCandles = append(interval15mCandles, x)
		case candle.Interval1h:
			interval1hCandles = append(interval1hCandles, x)
		}
	}
	assert.NotEmpty(t, interval15mCandles)
	assert.NotEmpty(t, interval1hCandles)
}
