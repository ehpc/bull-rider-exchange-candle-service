package main

import (
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/binanceapi"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candle"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candlemodel"
	myTesting "github.com/ehpc/bull-rider-exchange-candle-service/pkg/testing"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/transport"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testCandlesNumber = 2

func TestMainFlow(t *testing.T) {
	// Preparing mock binance.com transport
	apiTransport := myTesting.TransportMock{}
	apiTransport.AddReceivableMessage(transport.Message{
		Body: []byte(
			myTesting.GenerateCandlesJSON(myTesting.BinanceCandleExampleJSON, 2),
		),
	})

	// Fetching data from Binance
	api := binanceapi.NewBinanceAPI(&apiTransport)
	candles := api.GetCandles(
		[]candle.Pair{candle.IOTAUSDT},
		[]candle.Interval{candle.Interval1h},
	)

	// Pushing data to recipients
	modelTransport := myTesting.TransportMock{}
	model := candlemodel.NewCandleModel(&modelTransport)
	model.AddCandles(candles)

	// Verifying acceptable outgoing message format
	lastSentMessage, ok := modelTransport.GetLastSentMessage()
	assert.True(t, ok)
	assert.Equal(
		t,
		myTesting.BinanceCandleExampleProtobufMarshaled,
		lastSentMessage.Body,
	)
}
