package main

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/golang/protobuf/proto"
	
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/binanceapi"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candle"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candlemodel"
	myTesting "github.com/ehpc/bull-rider-exchange-candle-service/pkg/testing"
	protoCandle "github.com/ehpc/bull-rider/protobuf/go/candle"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/transport"
)

func TestMainFlow(t *testing.T) {
	const testCandlesCount = 2
	
	// Preparing mock binance.com transport
	apiTransport := myTesting.NewTransportMock()
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
	api := binanceapi.NewAPI(apiTransport)
	candles, err := api.GetCandles(
		[]candle.Pair{candle.PairIOTAUSDT},
		[]candle.Interval{candle.Interval15m},
	)
	assert.NoError(t, err)

	// Creating model
	modelTransport := myTesting.NewTransportMock()
	model := candlemodel.NewCandleModel(modelTransport)

	t.Run("AddCandles", func (t *testing.T) {
		// Pushing data to recipients
		result, err := model.AddCandles(candles)
		assert.NoError(t, err)
		assert.True(t, result)

		// Verifying acceptable outgoing message format
		protoCandles := protoCandle.Candles{
			Candles: []*protoCandle.Candle{
				&myTesting.BinanceIOTAUSDT15mCandleExampleProtobuf,
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
	})

	t.Run("AddCandle", func (t *testing.T) {
		// Pushing data to recipients
		result, err := model.AddCandle(candles[0])
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
	})

}
