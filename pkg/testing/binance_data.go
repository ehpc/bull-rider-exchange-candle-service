package testing

import (
	"github.com/golang/protobuf/proto"
	"github.com/ehpc/bull-rider/protobuf/go/candle"
)

// BinanceCandleExampleJSON is an example of Binance API candle
const BinanceCandleExampleJSON = `[
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
]`

// BinanceCandleExampleProtobuf is an example of Binance candle converted to
// protobuf format
var BinanceCandleExampleProtobuf = &candle.Candle{}

// BinanceCandleExampleProtobufMarshaled is an example of Binance candle 
// converted to protobuf format and marshaled to []bytes
var BinanceCandleExampleProtobufMarshaled, _ = proto.Marshal(BinanceCandleExampleProtobuf)

// GenerateCandlesJSON generates json with candles from Binance API
func GenerateCandlesJSON(candleJSON string, count int) string {
	json := "["
	for i := 0; i < count; i++ {
		json += candleJSON
		if i < count-1 {
			json += ","
		}
	}
	return json + "]"
}
