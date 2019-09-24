package main

import (
	"ehpc.io/bull-rider/exchange-candle-service/pkg/binanceapi"
	"ehpc.io/bull-rider/exchange-candle-service/pkg/candle"
	"ehpc.io/bull-rider/exchange-candle-service/pkg/candlemodel"
	myTesting "ehpc.io/bull-rider/exchange-candle-service/pkg/testing"
	"ehpc.io/bull-rider/exchange-candle-service/pkg/transport"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testCandlesNumber = 2

func TestMainFlow(t *testing.T) {
	message := transport.Message{
		Body: []byte(
			myTesting.GenerateCandlesJSON(2),
		),
	}

	apiTransport := myTesting.TransportMock{}
	apiTransport.AddReceivableMessage(message)
	// Fetching data from Binance
	api := binanceapi.NewBinanceAPI(&apiTransport)
	candles := api.GetCandles([]candle.Pair{candle.IOTAUSDT}, []candle.Interval{candle.Interval1h})

	modelTransport := myTesting.TransportMock{}
	// Pushing data to recipients
	model := candlemodel.NewCandleModel(&modelTransport)
	model.AddCandles(candles)

	got, ok := modelTransport.GetLastSentMessageAsString()
	assert.True(t, ok)
	assert.Contains(t, got, "1561622400000")
	assert.Contains(t, got, "5370")
}
