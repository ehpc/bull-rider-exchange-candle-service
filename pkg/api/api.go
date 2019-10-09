package api

import (
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candle"
)

// API is an interface for API implementations
type API interface {
	// GetCandles synchronously gets recent candles for specified intervals
	GetCandles(pairs []candle.Pair, intervals []candle.Interval) ([]candle.Candle, error)
	// WaitForCandles retrieves new candles asynchronously
	WaitForCandles(pairs []candle.Pair, intervals []candle.Interval) (chan candle.Candle, chan error)
}
