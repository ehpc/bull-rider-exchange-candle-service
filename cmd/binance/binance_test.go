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

func TestMainFlow(t *testing.T) {
	const testCandlesCount= 2
	
	// Preparing mock binance.com transport
	apiTransport := myTesting.TransportMock{}
	apiTransport.AddReceivableMessage(
		transport.Message{
			Body: []byte(
				myTesting.GenerateCandlesJSON(myTesting.BinanceCandleExampleJSON, testCandlesCount),
			),
		},
		binanceapi.GetCandlesRequestParams{
			Symbol: candle.PairIOTAUSDT,
			Interval: candle.Interval1h,
		},
	)

	// Fetching data from Binance
	api := binanceapi.NewBinanceAPI(&apiTransport)
	candles, err := api.GetCandles(
		[]candle.Pair{candle.PairIOTAUSDT},
		[]candle.Interval{candle.Interval1h},
	)
	assert.NoError(t, err)

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
