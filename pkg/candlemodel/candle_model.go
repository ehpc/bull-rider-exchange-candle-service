package candlemodel

import (
	"ehpc.io/bull-rider/exchange-candle-service/pkg/candle"
	"ehpc.io/bull-rider/exchange-candle-service/pkg/transport"
)

// CandleModel is a model for storing candles
type CandleModel struct{
	transport transport.Transport
}

//NewCandleModel creates new Candle Model
func NewCandleModel(transport transport.Transport) CandleModel {
	return CandleModel{
		transport: transport,
	}
}

// AddCandle stores a candle
// Returns true on success
func (model *CandleModel) AddCandle(candle candle.Candle) bool {
	return true
}
