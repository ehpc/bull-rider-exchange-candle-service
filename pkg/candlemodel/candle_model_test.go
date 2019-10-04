package candlemodel

import (
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candle"
	myTesting "github.com/ehpc/bull-rider-exchange-candle-service/pkg/testing"
	protoCandle "github.com/ehpc/bull-rider/protobuf/go/candle"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCandleModel(t *testing.T) {
	testCandle := candle.Candle{
		Exchange: candle.ExchangeBinance,
		Pair: candle.PairIOTAUSDT,
		Interval: candle.Interval15m,
		OpenTime:     1568707200000,
		CloseTime:    1568721599999,
		Open:         0.25000000,
		Close:        0.24850000,
		High:         0.25250000,
		Low:          0.24800000,
		Volume:       722111.71000000,
		QuoteVolume:  180376.63551900,
		TradesCount: 599,
	}
	protoCandles := protoCandle.Candles{
		Candles: []*protoCandle.Candle{
			&protoCandle.Candle{
				Exchange: string(testCandle.Exchange),
				Pair: string(testCandle.Pair),
				Interval: string(testCandle.Interval),
				OpenTime: testCandle.OpenTime,
				CloseTime: testCandle.CloseTime,
				Open: testCandle.Open,
				Close: testCandle.Close,
				High: testCandle.High,
				Low: testCandle.Low,
				Volume: testCandle.Volume,
				QuoteVolume: testCandle.QuoteVolume,
				TradesCount: testCandle.TradesCount,
			},
		},
	}
	t.Run("Successfully add one candle", func(t *testing.T) {
		modelTransport := myTesting.NewTransportMock()
		candleModel := NewCandleModel(modelTransport)
		got, err := candleModel.AddCandle(testCandle)
		assert.NoError(t, err)
		assert.Equal(t, true, got)
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
