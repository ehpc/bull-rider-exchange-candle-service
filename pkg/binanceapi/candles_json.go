package binanceapi

import (
	"errors"
	"encoding/json"
)

// CandleJSON is a type for unmarshaling binance klines
type CandleJSON struct {
	OpenTime int64
	Open string
	High string
	Low string
	Close string
	Volume string
	CloseTime int64
	QuoteVolume string
	TradesCount int64
	TakerBuyVolume string
	TakerBuyQuoteVolume string
	Ignored1 string
}

// UnmarshalJSON provides unmarshaler for binance kline tuple
func (c *CandleJSON) UnmarshalJSON(data []byte) error {
	tuple := []interface{}{
		&c.OpenTime,
		&c.Open,
		&c.High,
		&c.Low,
		&c.Close,
		&c.Volume,
		&c.CloseTime,
		&c.QuoteVolume,
		&c.TradesCount,
		&c.TakerBuyVolume,
		&c.TakerBuyQuoteVolume,
		&c.Ignored1,
	}
	expectedLen := len(tuple)
	err := json.Unmarshal(data, &tuple)
	if expectedLen != len(tuple) {
		return errors.New("Unexpected tuple length")
	}
	return err
}
