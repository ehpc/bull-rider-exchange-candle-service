package candlemodel

import (
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candle"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/transport"
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
func (model *CandleModel) AddCandle(candle candle.Candle) bool {
	return true
}

// AddCandles stores candles
// Returns true on success
func (model *CandleModel) AddCandles(candles []candle.Candle) bool {
	message := transport.Message{
		Body: []byte{},
	}
	return model.transport.Send(message)
}
