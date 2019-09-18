package main

import (
	"ehpc.io/bull-rider/exchange-candle-service/pkg/binanceapi"
	"ehpc.io/bull-rider/exchange-candle-service/pkg/transport"
	"ehpc.io/bull-rider/exchange-candle-service/pkg/candle"
	"ehpc.io/bull-rider/exchange-candle-service/pkg/candlemodel"
	mytesting "ehpc.io/bull-rider/exchange-candle-service/pkg/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMainFlow(t *testing.T) {
	message := transport.Message{
		Body: []byte(
			`[
				[
					1561622400000,
					"0.42590000",
					"0.42680000",
					"0.40050000",
					"0.42210000",
					"2589747.41000000",
					1561636799999,
					"1073262.29118300",
					5370,
					"1185070.39000000",
					"492584.11707600",
					"0"
				],
				[
					1561636800000,
					"0.42230000",
					"0.42500000",
					"0.40570000",
					"0.41890000",
					"1931449.37000000",
					1561651199999,
					"800614.36450200",
					3613,
					"1062134.08000000",
					"440770.60913900",
					"0"
				]
			]`,
		),
	}

	apiTransport := mytesting.TransportMock{}
	apiTransport.AddReceivableMessage(message)
	api := binanceapi.NewBinanceAPI(&apiTransport)
	candles := api.GetCandles([]string{"IOTAUSDT"}, []candle.Interval{candle.Interval1h})

	modelTransport := mytesting.TransportMock{}
	model := candlemodel.NewCandleModel(&modelTransport)
	model.AddCandles(candles)

	got := transport.GetLastSentMessageAsString()
	want := string(message.Body)

	assert.Equal(t, got, want)
}
