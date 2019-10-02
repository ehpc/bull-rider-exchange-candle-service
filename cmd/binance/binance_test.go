package main

import (
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/binanceapi"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candle"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candlemodel"
	myTesting "github.com/ehpc/bull-rider-exchange-candle-service/pkg/testing"
	protoCandle "github.com/ehpc/bull-rider/protobuf/go/candle"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/transport"
	"github.com/stretchr/testify/assert"
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestMainFlow(t *testing.T) {
	const testCandlesCount= 1
	
	// Preparing mock binance.com transport
	apiTransport := myTesting.TransportMock{}
	apiTransport.AddReceivableMessage(
		transport.Message{
			Body: []byte(
				myTesting.GenerateCandlesJSON(myTesting.BinanceCandleExampleJSON, testCandlesCount),
			),
		},
		transport.RequestParams{
			"HTTPMethod": "GET",
			"HTTPPath": "/api/v1/klines",
			"symbol": candle.PairIOTAUSDT,
			"interval": candle.Interval15m,
		},
	)

	// Fetching data from Binance
	api := binanceapi.NewAPI(&apiTransport)
	candles, err := api.GetCandles(
		[]candle.Pair{candle.PairIOTAUSDT},
		[]candle.Interval{candle.Interval15m},
	)
	assert.NoError(t, err)

	// Pushing data to recipients
	modelTransport := myTesting.TransportMock{}
	model := candlemodel.NewCandleModel(&modelTransport)
	result, err := model.AddCandles(candles)
	assert.NoError(t, err)
	assert.True(t, result)

	// Verifying acceptable outgoing message format
	protoCandles := protoCandle.Candles{
		Candles: []*protoCandle.Candle{
			&myTesting.BinanceIOTAUSDT15mCandleExampleProtobuf,
		},
	}
	protoCandlesMarshaled, err := proto.Marshal(&protoCandles)
	assert.NoError(t, err)
	lastSentMessage, ok := modelTransport.GetLastSentMessage()
	assert.True(t, ok)
	assert.Equal(
		t,
		protoCandlesMarshaled,
		lastSentMessage.Body,
	)
}

func TestRealFlow(t *testing.T) {
	// // Creating API transport
	// apiTransport := transport.NewHTTPTransport(binanceapi.APIURL)
	// // Fetching data from Binance
	// api := binanceapi.NewAPI(&apiTransport)
	// candles, err := api.GetCandles(
	// 	[]candle.Pair{candle.PairIOTAUSDT},
	// 	[]candle.Interval{candle.Interval15m},
	// )
	// assert.NoError(t, err)
	// // Creating model transport
	// modelTransport := transport.NewRabbitMQTransport()
	// // Pushing data to recipients
	// model := candlemodel.NewCandleModel(&modelTransport)
	// result, err := model.AddCandles(candles)
	// assert.NoError(t, err)
	// assert.True(t, result)
}