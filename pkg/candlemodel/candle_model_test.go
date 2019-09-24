package candlemodel

import (
	. "github.com/ehpc/bull-rider-exchange-candle-service/pkg/candle"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCandleModel(t *testing.T) {
	candle := Candle{
		OpenTime:     1568707200000,
		CloseTime:    1568721599999,
		Open:         0.25000000,
		Close:        0.24850000,
		High:         0.25250000,
		Low:          0.24800000,
		Volume:       722111.71000000,
		QuoteVolume:  180376.63551900,
		TradesNumber: 599,
	}
	t.Run("Add a candle", func(t *testing.T) {
		candleModel := CandleModel{}
		got := candleModel.AddCandle(candle)
		want := true
		assert.Equal(t, got, want)
	})
}
