package binanceapi

import (
	"encoding/json"
	"errors"
)

// CandleRESTJSON is a type for unmarshaling binance REST klines
type CandleRESTJSON struct {
	OpenTime            int64
	Open                string
	High                string
	Low                 string
	Close               string
	Volume              string
	CloseTime           int64
	QuoteVolume         string
	TradesCount         int64
	TakerBuyVolume      string
	TakerBuyQuoteVolume string
	Ignored1            string
}

// UnmarshalJSON provides unmarshaler for binance REST kline tuple
func (c *CandleRESTJSON) UnmarshalJSON(data []byte) error {
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
		return errors.New("unexpected tuple length")
	}
	return err
}

// CandleWebsocketJSON is a type for unmarshaling websocket klines
type CandleWebsocketJSON struct {
	EventType string `json:"e"`
	EventTime int64 `json:"E"`
	Pair string `json:"s"`
	Kline CandleWebsocketJSONKline `json:"k"`
}

// CandleWebsocketJSONKline is an inner kline object of CandleWebsocketJSON
type CandleWebsocketJSONKline struct {
	OpenTime            int64 `json:"t"`
	CloseTime           int64 `json:"T"`
	Pair string `json:"s"`
	Interval string `json:"i"`
	FirstTradeID int64 `json:"f"`
	LastTradeID int64 `json:"L"`
	Open                string `json:"o"`
	Close               string `json:"c"`
	High                string `json:"h"`
	Low                 string `json:"l"`
	Volume              string `json:"v"`
	TradesCount         int64 `json:"n"`
	Closed	bool `json:"x"`
	QuoteVolume         string `json:"q"`
	TakerBuyVolume      string `json:"V"`
	TakerBuyQuoteVolume string `json:"Q"`
	Ignored1            string `json:"B"`
}
