package candlemodel

import (
	"github.com/golang/protobuf/proto"

	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candle"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/transport"
	protoCandle "github.com/ehpc/bull-rider/protobuf/go/candle"
)

// CandleModel is a model for storing candles
type CandleModel struct {
	transport transport.Transport
}

// NewCandleModel creates new Candle Model
func NewCandleModel(transport transport.Transport) CandleModel {
	return CandleModel{
		transport: transport,
	}
}

// AddCandle stores one candle
// Returns true on success
func (model *CandleModel) AddCandle(oneCandle candle.Candle) (bool, error) {
	return model.AddCandles([]candle.Candle{oneCandle})
}

// AddCandles stores many candles
// Returns true on success
func (model *CandleModel) AddCandles(candles []candle.Candle) (bool, error) {
	protoCandles := make([]*protoCandle.Candle, len(candles))
	for i, cndl := range candles {
		protoCandles[i] = &protoCandle.Candle{
			Exchange: string(cndl.Exchange),
			Pair: string(cndl.Pair),
			Interval: string(cndl.Interval),
			OpenTime: cndl.OpenTime,
			CloseTime: cndl.CloseTime,
			Open: cndl.Open,
			Close: cndl.Close,
			High: cndl.High,
			Low: cndl.Low,
			Volume: cndl.Volume,
			QuoteVolume: cndl.QuoteVolume,
			TradesCount: cndl.TradesCount,
		}
	}
	protoCandlesMarshaled, err := proto.Marshal(&protoCandle.Candles{
		Candles: protoCandles,
	})
	if err != nil {
		return false, err
	}
	message := transport.Message{
		Body: protoCandlesMarshaled,
	}
	return model.transport.Send(message)
}
