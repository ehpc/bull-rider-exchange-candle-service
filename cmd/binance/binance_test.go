package main

import (
	"fmt"
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

func TestRESTFlow(t *testing.T) {
	const testCandlesCount = 2
	
	// Preparing mock binance.com transport
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

	// Fetching data from Binance
	api := binanceapi.NewAPI(apiTransport, nil)
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

func TestWebsocketFlow(t *testing.T) {
	const testCandlesCount = 2

	// Preparing mock binance.com transport
	apiTransport := myTesting.NewTransportMock()
	apiTransport.AddReceivableMessage(
		transport.Message{
			Body: []byte(
				myTesting.BinanceCandleExampleWebsocketJSON,
			),
		},
		transport.RequestParams{
			"WebsocketPath": fmt.Sprintf("/ws/%s@kline_%s", candle.PairIOTAUSDT, candle.Interval15m),
		},
	)

	// Fetching data from Binance
	api := binanceapi.NewAPI(nil, apiTransport)
	candleChannel, errorChannel := api.WaitForCandles(
		[]candle.Pair{candle.PairIOTAUSDT},
		[]candle.Interval{candle.Interval15m},
	)
	select{
	case cndl := <-candleChannel:
		assert.NotNil(t, cndl)
		assert.Greater(t, cndl.OpenTime, int64(0))
		assert.Greater(t, cndl.CloseTime, int64(0))
	case err := <-errorChannel:
		assert.NoError(t, err)
	}
}
