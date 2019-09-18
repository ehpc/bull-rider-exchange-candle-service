package binanceapi

import (
	"ehpc.io/bull-rider/exchange-candle-service/pkg/candle"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCandles(t *testing.T) {
	
	api := NewBinanceAPI()
	pairs := []string{"IOTAUSDT", "ETHUSDT"}
	intervals := []candle.Interval{candle.Interval15m, candle.Interval1h}
	got := api.GetCandles(pairs, intervals)
	assert.Contains(t, got, "IOTAUSDT")
	assert.Contains(t, got, "ETHUSDT")
	assert.Contains(t, got["IOTAUSDT"], candle.Interval15m)
	assert.Contains(t, got["IOTAUSDT"], candle.Interval1h)
	assert.Contains(t, got["ETHUSDT"], candle.Interval15m)
	assert.Contains(t, got["ETHUSDT"], candle.Interval1h)
	assert.NotEmpty(t, got["IOTAUSDT"][candle.Interval15m])
	assert.NotEmpty(t, got["IOTAUSDT"][candle.Interval1h])
	assert.NotEmpty(t, got["ETHUSDT"][candle.Interval15m])
	assert.NotEmpty(t, got["ETHUSDT"][candle.Interval1h])
}
